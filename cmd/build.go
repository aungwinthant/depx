package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build Docker image",
	Long:  "Build a Docker image for deployment using the depx config file",
	RunE: func(cmd *cobra.Command, args []string) error {
		configPath, err := cmd.Flags().GetString("config")
		if err != nil {
			return fmt.Errorf("reading config flag: %w", err)
		}
		cfg, err := loadConfig(configPath)
		if err != nil {
			return fmt.Errorf("loading config: %w", err)
		}
		if err := runCmd("docker", "build", "-t", imageRef(cfg), "-f", cfg.App.Dockerfile, cfg.App.Context); err != nil {
			return fmt.Errorf("build failed: %w", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
