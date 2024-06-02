package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	botCmd = &cobra.Command{
		Use:              "gitlab-orchestrator [OPTIONS] [COMMANDS]",
		Short:            "gitlab-orchestrator – command-line tool to interact with gitlab platform",
		Version:          "0.0.1",
		SilenceErrors:    true,
		SilenceUsage:     true,
		TraverseChildren: true,
	}
)

func Execute() error {
	return botCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	botCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $PWD/.gitops-bot.yaml)")
	viper.BindPFlag("config", botCmd.PersistentFlags().Lookup("config"))

	log.SetPrefix("gitlab: ")
}

func initConfig() {
	viper.SetEnvPrefix("gitlab")
	viper.AutomaticEnv()
}
