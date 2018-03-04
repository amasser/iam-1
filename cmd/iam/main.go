package main

import (
	"log"
	"os"

	"github.com/maurofran/iam/internal/app/application"

	"github.com/maurofran/iam/internal/app/domain/model"

	"github.com/facebookgo/inject"
	"github.com/maurofran/iam/internal/app/ports/adapter/mongo"
	"github.com/pkg/errors"
	cli "gopkg.in/urfave/cli.v1"
	db "upper.io/db.v3"
	mongodb "upper.io/db.v3/mongo"
)

var g inject.Graph
var mongoDB string

func main() {

	app := cli.NewApp()
	app.Name = "iam"
	app.Usage = "Identity and Access Manager CLI"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "dburl",
			Value:       "mongodb://localhost:27017/iam_db",
			Usage:       "Connect to MongoDB at `URL`",
			Destination: &mongoDB,
			EnvVar:      "DB_URL",
		},
	}
	app.After = shutdownContext
	app.Commands = []cli.Command{
		cli.Command{
			Name:        "tenant",
			Usage:       "Run a command on specific tenant",
			Aliases:     []string{"t"},
			Before:      setupContext,
			Subcommands: []cli.Command{},
		},
		cli.Command{
			Name:        "user",
			Usage:       "Run a command on specific user",
			Aliases:     []string{"u"},
			Before:      setupContext,
			Subcommands: []cli.Command{},
		},
		cli.Command{
			Name:        "group",
			Usage:       "Run a command on specific group",
			Aliases:     []string{"g"},
			Before:      setupContext,
			Subcommands: []cli.Command{},
		},
		cli.Command{
			Name:        "role",
			Usage:       "Run a command on specific role",
			Aliases:     []string{"r"},
			Before:      setupContext,
			Subcommands: []cli.Command{},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func setupContext(c *cli.Context) error {
	url, err := mongodb.ParseURL(mongoDB)
	if err != nil {
		return errors.Wrapf(err, "Error occurred while parsing URL %s", mongoDB)
	}
	session, err := mongodb.Open(url)
	if err != nil {
		return errors.Wrapf(err, "Error occurred while opening connection to %s", mongoDB)
	}

	err = g.Provide(
		&inject.Object{Value: session},

		&inject.Object{Value: new(mongo.TenantRepository)},
		&inject.Object{Value: new(mongo.UserRepository)},
		&inject.Object{Value: new(mongo.GroupRepository)},
		&inject.Object{Value: new(mongo.RoleRepository)},

		&inject.Object{Value: new(model.GroupMemberService)},
		&inject.Object{Value: new(model.TenantProvisioningService)},
		&inject.Object{Value: new(model.AuthenticationService)},
		&inject.Object{Value: new(model.AuthorizationService)},

		&inject.Object{Value: new(application.TenantService)},
		&inject.Object{Value: new(application.UserService)},
		&inject.Object{Value: new(application.GroupService)},
		&inject.Object{Value: new(application.RoleService)},
	)
	if err != nil {
		return err
	}
	return g.Populate()
}

func shutdownContext(c *cli.Context) error {
	for _, obj := range g.Objects() {
		if session, ok := obj.Value.(db.Database); ok {
			return session.Close()
		}
	}
	return nil
}
