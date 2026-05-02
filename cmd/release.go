package cmd

import (
	"github.com/spf13/cobra"
)

var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Release Docker image to remote server",
	Long:  "Deploy the Docker image to a remote server via SSH",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(releaseCmd)
}
