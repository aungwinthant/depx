package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push Docker image to registry",
	Long:  "Push the built Docker image to a registry using depx.yaml configuration",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := loadConfig("depx.yaml")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error loading depx.yaml:", err)
			return
		}
		if err := runCmd("docker", "push", imageRef(cfg)); err != nil {
			fmt.Fprintln(os.Stderr, "Push failed:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}
