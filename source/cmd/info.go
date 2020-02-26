package cmd

import (
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    log "github.com/sirupsen/logrus")

// infoCmd represents the list command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show detailed information about chroot",
	Run: func(cmd *cobra.Command, args []string) {
        getInfo()
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}

func getInfo(){
    log.Info("list called")
    viper.Get("name")
}
