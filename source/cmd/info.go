package cmd

import (
    "github.com/spf13/cobra"
    log "github.com/sirupsen/logrus")

// infoCmd represents the list command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show detailed information about chroot",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("list called")
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
