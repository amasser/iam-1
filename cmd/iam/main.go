package main

import (
	"log"
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

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
	app.Commands = []cli.Command{
		cli.Command{
			Name:        "tenant",
			Usage:       "Run a command on specific tenant",
			Aliases:     []string{"t"},
			Subcommands: []cli.Command{},
		},
		cli.Command{
			Name:        "user",
			Usage:       "Run a command on specific user",
			Aliases:     []string{"u"},
			Subcommands: []cli.Command{},
		},
		cli.Command{
			Name:        "group",
			Usage:       "Run a command on specific group",
			Aliases:     []string{"g"},
			Subcommands: []cli.Command{},
		},
		cli.Command{
			Name:        "role",
			Usage:       "Run a command on specific role",
			Aliases:     []string{"r"},
			Subcommands: []cli.Command{},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
