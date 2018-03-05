package main

import (
	"fmt"
	"net"
	"os"

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
)

// Injected variables
var (
	Version string
	Commit  string
	Branch  string
)

var g inject.Graph
var mongoDB string
var grpcPort int
var environment string

var database db.Database
var tenantServer *grpc_adapter.TenantServer

func main() {
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

	log.WithField("database", database).Debug("Trying to parse database URL")
	url, err := mongo.ParseURL(mongoDB)
	if err != nil {
		log.WithError(err).WithField("database", mongoDB).Error("An error occurred parsing database URL")
		return err
	}
	log.WithField("database", database).Info("Connecting to database")
	database, err = mongo.Open(url)
	if err != nil {
		log.WithError(err).WithField("database", database).Error("An error occurred connecting to database")
		return err
	}
	log.WithField("database", database).Debug("Database connection successful")

	tenantServer = &grpc_adapter.TenantServer{}

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
	grpcServer := grpc.NewServer()
	grpc_adapter.RegisterTenantServiceServer(grpcServer, tenantServer)

	log.WithField("port", grpcPort).Info("Starting GRPC server")

	err = grpcServer.Serve(lis)
	if err != nil {
		log.WithError(err).Error("An error occurred while starting GRPC server")
		return err
	}
	log.Debug("Shutting down GRPC server")
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
