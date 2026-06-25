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
  insecure_skip_host_key: false   # set true only to opt out of known_hosts verification
  container:
    name: my-app
    port: "3000:3000"             # host:container port mapping
    env_file: .env                # passed to docker run as --env-file
    network: ""                   # join an existing docker network, e.g. myapp_default
    aliases: []                   # network aliases, e.g. [backend]
`

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize depx configuration",
	Long:  "Create a depx configuration file in the current project",
	RunE: func(cmd *cobra.Command, args []string) error {
		configPath, err := cmd.Flags().GetString("config")
		if err != nil {
			return fmt.Errorf("reading config flag: %w", err)
		}
		if _, err := os.Stat(configPath); err == nil {
			return fmt.Errorf("%s already exists", configPath)
		}
		if err := os.WriteFile(configPath, []byte(configTemplate), 0644); err != nil {
			return fmt.Errorf("creating %s: %w", configPath, err)
		}
		fmt.Printf("Created %s\n", configPath)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
