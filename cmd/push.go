package cmd

import (
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push Docker image to registry",
	Long:  "Push the built Docker image to a registry",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}
