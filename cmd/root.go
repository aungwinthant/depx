package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "depx",
	Short: "A minimal Docker deployment tool",
	Long:  "depx is a CLI tool that deploys Docker images to remote servers over SSH",
}

func Execute() {
	rootCmd.SilenceUsage = true
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("config", "c", "depx.yaml", "path to depx config file")
	rootCmd.PersistentFlags().Bool("insecure-skip-host-key", false, "skip SSH host key verification (insecure)")
}
