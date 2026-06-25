package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push Docker image to registry",
	Long:  "Push the built Docker image to a registry using the depx config file",
	RunE: func(cmd *cobra.Command, args []string) error {
		configPath, err := cmd.Flags().GetString("config")
		if err != nil {
			return fmt.Errorf("reading config flag: %w", err)
		}
		cfg, err := loadConfig(configPath)
		if err != nil {
			return fmt.Errorf("loading config: %w", err)
		}
		if err := runCmd("docker", "push", imageRef(cfg)); err != nil {
			return fmt.Errorf("push failed: %w", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}
