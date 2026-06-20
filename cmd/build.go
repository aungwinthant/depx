package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build Docker image",
	Long:  "Build a Docker image for deployment using depx.yaml configuration",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := loadConfig("depx.yaml")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error loading depx.yaml:", err)
			return
		}
		if err := runCmd("docker", "build", "-t", imageRef(cfg), "-f", cfg.App.Dockerfile, cfg.App.Context); err != nil {
			fmt.Fprintln(os.Stderr, "Build failed:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
