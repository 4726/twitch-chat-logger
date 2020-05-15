package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "twitch-chat-logger",
	Short: "Logs twitch chat and allows searching of messages.",
}
var cfgFile string

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config_dev.yml", "path to config file")
	rootCmd.AddCommand(serverCmd())
}

func Execute() error {
	return rootCmd.Execute()
}
