package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
)

var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Release Docker image to remote server",
	Long:  "Deploy the Docker image to a remote server via SSH using depx.yaml configuration",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := loadConfig("depx.yaml")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error loading depx.yaml:", err)
			return
		}

		keyData, err := os.ReadFile(expandPath(cfg.Deploy.Key))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading SSH key:", err)
			return
		}
		signer, err := ssh.ParsePrivateKey(keyData)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing SSH key:", err)
			return
		}

		client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", cfg.Deploy.Host, cfg.Deploy.Port), &ssh.ClientConfig{
			User:            cfg.Deploy.User,
			Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		})
		if err != nil {
			fmt.Fprintln(os.Stderr, "SSH connection failed:", err)
			return
		}
		defer client.Close()

		ref := imageRef(cfg)
		commands := []string{
			fmt.Sprintf("docker pull %s", ref),
			fmt.Sprintf("docker rm -f %s || true", cfg.Deploy.Container.Name),
			fmt.Sprintf("docker run -d --name %s -p %s %s", cfg.Deploy.Container.Name, cfg.Deploy.Container.Port, ref),
		}
		for _, command := range commands {
			session, err := client.NewSession()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Session error:", err)
				return
			}
			session.Stdout = os.Stdout
			session.Stderr = os.Stderr
			if err := session.Run(command); err != nil {
				fmt.Fprintln(os.Stderr, "Command failed:", command, err)
				session.Close()
				return
			}
			session.Close()
		}
	},
}

func init() {
	rootCmd.AddCommand(releaseCmd)
}
