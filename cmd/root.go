package cmd

import (
	"calendarproto/internal/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "calendarproto",
	Short: "Calendar micro-service",
	Long:  `Calendar micro-service`,
	Run:   run,
}

func RootCommand() *cobra.Command {
	rootCmd.Flags().IntP("port", "p", 8080, "HTTP port (default 8080)")
	rootCmd.Flags().StringP("config", "c", "",
		"config file (default is calendarproto.[json|toml|yaml|yml|properties|props|prop|hcl])")
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
	logger.Info("Starting..")
}
