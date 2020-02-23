package cmd

import (
    "github.com/spf13/cobra"
    log "github.com/sirupsen/logrus")

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all chroot enviroiment",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("list called")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
