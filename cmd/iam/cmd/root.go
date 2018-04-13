package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile            string
	enableTLS          bool
	caFile             string
	serverAddress      string
	serverHostOverride string
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default to $HOME/.iam.yaml)")
	rootCmd.PersistentFlags().BoolVar(&enableTLS, "enableTLS", false, "enable tls (default false)")
	rootCmd.PersistentFlags().StringVar(&caFile, "caFile", "", "certification authority file")
	rootCmd.PersistentFlags().StringVar(&serverAddress, "serverAddress", "localhost:3000", "address of iamd server (default to localhost:3000)")
	rootCmd.PersistentFlags().StringVar(&serverHostOverride, "serverHostOverride", "", "override for server host")
	rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.BindPFlag("enableTLS", rootCmd.PersistentFlags().Lookup("enableTLS"))
	viper.BindPFlag("caFile", rootCmd.PersistentFlags().Lookup("caFile"))
	viper.BindPFlag("serverAddress", rootCmd.PersistentFlags().Lookup("serverAddress"))
	viper.BindPFlag("serverHostOverride", rootCmd.PersistentFlags().Lookup("serverHostOverride"))
	viper.SetDefault("TLS", false)
	viper.SetDefault("CaFile", "")
	viper.SetDefault("ServerAddress", "localhost:3000")
	viper.SetDefault("ServerHostOverride", "")
}

var rootCmd = &cobra.Command{
	Use:   "iam",
	Short: "iam is the command line client tool to interact with iamd",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Execute the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Println("Can't read config:", err)
			os.Exit(1)
		}
	}
}

/*

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
}*/
