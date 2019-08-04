package cmd

import (
	"github.com/Brialius/servicesceleton/internal/config"
	"github.com/Brialius/servicesceleton/internal/server"
	"github.com/Brialius/servicesceleton/internal/welcome"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
	build   = "local"
)

var rootCmd = &cobra.Command{
	Use:   "servicesceleton",
	Short: "Template for micro-services",
	Long:  `Template for micro-services`,
	Run:   run,
}

func RootCommand() *cobra.Command {
	rootCmd.Flags().IntP("port", "p", 8080, "HTTP port (default 8080)")
	rootCmd.Flags().StringP("welcomeString", "w", "Hello world!", "HTTPWelcome string")
	rootCmd.Flags().StringP("config", "c", "",
		"config file (default is servicesceleton.[json|toml|yaml|yml|properties|props|prop|hcl])")
	rootCmd.Flags().StringP("verbosity", "v", "warning",
		"Log level (trace, debug, info, warn, error, fatal, panic)")
	return rootCmd
}

func run(cmd *cobra.Command, args []string) {
	conf := config.LoadConfig(cmd)
	logger, err := config.ConfigureLogger(&conf.LogConfig)
	if err != nil {
		logrus.Fatal(err)
	}
	logger.Debugf("Root command args: %v", args)
	logger.Debug(conf)
	logger.Infof("Version %s-%s", version, build)
	logger.Info("Starting..")
	ws := welcome.NewService(logger, conf.WelcomeString)
	hs := server.NewServer(conf.Port, ws.HTTPWelcome)
	logger.Fatal(hs.Start())
}
