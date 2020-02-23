package cmd

import (
    "github.com/spf13/cobra"
    log "github.com/sirupsen/logrus")

// newCmd represents the list command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "create a new chroot environment",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("list called")
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
