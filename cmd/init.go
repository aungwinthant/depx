package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const configTemplate = `app:
  name: my-app
  dockerfile: Dockerfile
  context: .

image:
  name: my-app
  tag: latest
  registry: docker.io

deploy:
  host: your-server.com
  port: 22
  user: deploy
  key: ~/.ssh/id_rsa
  container:
    name: my-app
    port: "3000:3000"
    env_file: .env
`

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize depx configuration",
	Long:  "Create a depx.yaml configuration file in the current project",
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat("depx.yaml"); err == nil {
			fmt.Println("depx.yaml already exists")
			return
		}
		if err := os.WriteFile("depx.yaml", []byte(configTemplate), 0644); err != nil {
			fmt.Fprintln(os.Stderr, "Error creating depx.yaml:", err)
			return
		}
		fmt.Println("Created depx.yaml")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
