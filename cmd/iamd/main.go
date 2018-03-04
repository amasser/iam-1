package main

import (
	"fmt"
	"log"
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

	"github.com/pkg/errors"
	cli "gopkg.in/urfave/cli.v1"
)

var g inject.Graph
var mongoDB string
var grpcPort int

func main() {

	app := cli.NewApp()
	app.Name = "iamd"
	app.Usage = "Identity and Access Manager Daemon"
	app.Version = "1.0.0"
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
	url, err := mongo.ParseURL(mongoDB)
	if err != nil {
		return errors.Wrapf(err, "Error occurred while parsing URL %s", mongoDB)
	}
	session, err := mongo.Open(url)
	if err != nil {
		return errors.Wrapf(err, "Error occurred while opening connection to %s", mongoDB)
	}

	err = g.Provide(
		&inject.Object{Value: session},

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

		&inject.Object{Value: grpc_adapter.TenantServer},
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
	grpc_adapter.RegisterTenantServiceServer(grpcServer, grpc_adapter.TenantServer)
	return grpcServer.Serve(lis)
}

func shutdownContext(c *cli.Context) error {
	for _, obj := range g.Objects() {
		if session, ok := obj.Value.(db.Database); ok {
			return session.Close()
		}
	}
	return nil
}
