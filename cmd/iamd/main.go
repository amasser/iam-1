package main

import (
	"log"
	"os"

	"github.com/pkg/errors"
	cli "gopkg.in/urfave/cli.v1"
	db "upper.io/db.v3/mongo"
)

var mongoDB string

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
	}
	app.Action = startDaemon

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func startDaemon(c *cli.Context) error {
	url, err := db.ParseURL(mongoDB)
	if err != nil {
		return errors.Wrapf(err, "Error occurred while parsing URL %s", mongoDB)
	}
	session, err := db.Open(url)
	if err != nil {
		return errors.Wrapf(err, "Error occurred while opening connection to %s", mongoDB)
	}
	defer session.Close()

	//	tenantRepo := &mongo.TenantRepository{Database: session}
	//	userRepo := &mongo.UserRepository{Database: session}
	//	groupRepo := &mongo.GroupRepository{Database: session}
	//	roleRepo := &mongo.RoleRepository{Database: session}

	return nil
}
