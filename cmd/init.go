package cmd

import (
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize depx configuration",
	Long:  "Initialize depx configuration file in the current project",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
