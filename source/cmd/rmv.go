package cmd

import (
    "github.com/spf13/cobra"
    log "github.com/sirupsen/logrus")

// remCmd represents the list command
var remCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove a choosen chroot environment",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("list called")
	},
}

func init() {
	rootCmd.AddCommand(remCmd)
    remCmd.Aliases = []string{"rm","remove"}
}
