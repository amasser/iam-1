package main

import (
	"google.golang.org/grpc"
	"upper.io/db.v3"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

// Injected variables
var (
	Version string
	Commit  string
	Branch  string
)

var (
	mongoDB     string
	grpcPort    int
	environment string
)

var (
	database   db.Database
	grpcServer *grpc.Server
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

}
