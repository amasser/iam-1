package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"upper.io/db.v3"
	"upper.io/db.v3/mongo"

	"github.com/maurofran/iam/internal/app/application"

	"github.com/maurofran/iam/internal/app/domain/model"

	"github.com/facebookgo/inject"
	grpc_adapter "github.com/maurofran/iam/internal/app/ports/adapter/grpc"
	mongo_adapter "github.com/maurofran/iam/internal/app/ports/adapter/mongo"

	log "github.com/sirupsen/logrus"
	cli "gopkg.in/urfave/cli.v1"

	"github.com/spf13/viper"
)

// Injected variables
var (
	Version string
	Commit  string
	Branch  string
)

var g inject.Graph

var (
	mongoDB     string
	grpcPort    int
	environment string
)

var (
	database     db.Database
	grpcServer   *grpc.Server
	tenantServer *grpc_adapter.TenantServer
	groupServer  *grpc_adapter.GroupServer
	roleServer   *grpc_adapter.RoleServer
	userServer   *grpc_adapter.UserServer
)

func main() {
	viper.SetDefault("DatabaseUrl", "mongodb://localhost:27017/iam_db")
	viper.SetDefault("GrpcServerPort", 3000)
	viper.SetDefault("Environment", "dev")

	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/iamd")
	viper.AddConfigPath("$HOME/.iamd")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatal(err)
		}
	}

	viper.SetEnvPrefix("iamd")
	viper.AutomaticEnv()

	app := cli.NewApp()
	app.Name = "iamd"
	app.Usage = "Identity and Access Manager Daemon"
	app.Version = Version
	app.Metadata = map[string]interface{}{
		"Commit": Commit,
		"Branch": Branch,
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "dburl",
			Value:       "mongodb://localhost:27017/iam_db",
			Usage:       "Connect to MongoDB at `URL`",
			Destination: &mongoDB,
			EnvVar:      "DB_URL",
		},
		cli.IntFlag{
			Name:        "grpcPort",
			Value:       3000,
			Usage:       "Exposes GRPC server at `grpcPort`",
			Destination: &grpcPort,
			EnvVar:      "GRPC_PORT",
		},
		cli.StringFlag{
			Name:        "environment",
			Usage:       "Provide the environment: dev, test, qas, prod",
			Destination: &environment,
			EnvVar:      "ENVIRONMENT",
		},
	}
	app.Before = setupContext
	app.Action = runDaemon
	app.After = shutdownContext

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func setupContext(c *cli.Context) error {
	setupLogging()

	log.WithField("database", mongoDB).Debug("Trying to parse database URL")
	url, err := mongo.ParseURL(mongoDB)
	if err != nil {
		log.WithError(err).WithField("database", mongoDB).Error("An error occurred parsing database URL")
		return err
	}
	log.WithField("database", mongoDB).Info("Connecting to database")
	database, err = mongo.Open(url)
	if err != nil {
		log.WithError(err).WithField("database", database).Error("An error occurred connecting to database")
		return err
	}
	log.WithField("database", mongoDB).Debug("Database connection successful")

	tenantServer = &grpc_adapter.TenantServer{}
	roleServer = &grpc_adapter.RoleServer{}

	err = g.Provide(
		&inject.Object{Value: database},

		&inject.Object{Value: new(mongo_adapter.TenantRepository)},
		&inject.Object{Value: new(mongo_adapter.UserRepository)},
		&inject.Object{Value: new(mongo_adapter.GroupRepository)},
		&inject.Object{Value: new(mongo_adapter.RoleRepository)},

		&inject.Object{Value: new(model.GroupMemberService)},
		&inject.Object{Value: new(model.TenantProvisioningService)},
		&inject.Object{Value: new(model.AuthenticationService)},
		&inject.Object{Value: new(model.AuthorizationService)},

		&inject.Object{Value: new(application.TenantService)},
		&inject.Object{Value: new(application.UserService)},
		&inject.Object{Value: new(application.GroupService)},
		&inject.Object{Value: new(application.RoleService)},

		&inject.Object{Value: tenantServer},
		&inject.Object{Value: groupServer},
		&inject.Object{Value: roleServer},
		&inject.Object{Value: userServer},
	)
	if err != nil {
		return err
	}
	return g.Populate()
}

func runDaemon(c *cli.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		return err
	}
	grpcServer = grpc.NewServer()
	grpc_adapter.RegisterTenantServiceServer(grpcServer, tenantServer)
	grpc_adapter.RegisterGroupServiceServer(grpcServer, groupServer)
	grpc_adapter.RegisterRoleServiceServer(grpcServer, roleServer)
	grpc_adapter.RegisterUserServiceServer(grpcServer, userServer)

	log.WithField("port", grpcPort).Info("Starting GRPC server")

	go signalHandler()

	err = grpcServer.Serve(lis)
	if err != nil {
		log.WithError(err).Error("An error occurred while starting GRPC server")

		return err
	}

	return nil
}

func shutdownContext(c *cli.Context) error {
	if database == nil {
		return nil
	}
	log.WithField("database", mongoDB).Debug("Shutting down connection")
	if err := database.Close(); err != nil {
		log.WithError(err).Error("An error occurred while closing connection")
		return err
	}
	log.Debug("Connection correctly shut down")
	return nil
}

func setupLogging() {
	switch environment {
	case "dev", "test":
		log.SetFormatter(&log.TextFormatter{})
		log.SetLevel(log.DebugLevel)
	case "qas":
		log.SetFormatter(&log.TextFormatter{})
		log.SetLevel(log.InfoLevel)
	case "prod":
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.ErrorLevel)
	}
	log.SetOutput(os.Stdout)
}

func signalHandler() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		syscall.SIGTERM, syscall.SIGINT, syscall.SIGUSR1,
		syscall.SIGTTIN, syscall.SIGTTOU,
		syscall.SIGUSR2)

	for {
		select {
		case sig := <-signalChan:
			switch sig {
			case syscall.SIGINT:
				log.Warn("SIGINT received, starting graceful shutdown")
				grpcServer.GracefulStop()
			case syscall.SIGTERM:
				log.Warn("SIGTERM received, shutting down immediately")
				grpcServer.Stop()
			default:
				log.WithField("signal", sig).Info("ignoring unknown signal")
			}
		default:
			time.Sleep(time.Second)
		}
	}
}
