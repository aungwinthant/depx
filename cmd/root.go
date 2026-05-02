package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "depx",
	Short: "A minimal Docker deployment tool",
	Long:  "depx is a CLI tool that deploys Docker images to remote servers over SSH",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
