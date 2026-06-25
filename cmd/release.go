package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Release Docker image to remote server",
	Long:  "Deploy the Docker image to a remote server via SSH using the depx config file",
	RunE: func(cmd *cobra.Command, args []string) error {
		configPath, err := cmd.Flags().GetString("config")
		if err != nil {
			return fmt.Errorf("reading config flag: %w", err)
		}
		cfg, err := loadConfig(configPath)
		if err != nil {
			return fmt.Errorf("loading config: %w", err)
		}

		keyData, err := os.ReadFile(expandPath(cfg.Deploy.Key))
		if err != nil {
			return fmt.Errorf("reading SSH key: %w", err)
		}
		signer, err := ssh.ParsePrivateKey(keyData)
		if err != nil {
			return fmt.Errorf("parsing SSH key: %w", err)
		}

		insecureFlag, err := cmd.Flags().GetBool("insecure-skip-host-key")
		if err != nil {
			return fmt.Errorf("reading insecure-skip-host-key flag: %w", err)
		}
		hostKeyCallback, err := hostKeyCallbackFor(cfg.Deploy.Host, insecureFlag || cfg.Deploy.InsecureSkipHostKey)
		if err != nil {
			return err
		}

		client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", cfg.Deploy.Host, cfg.Deploy.Port), &ssh.ClientConfig{
			User:            cfg.Deploy.User,
			Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
			HostKeyCallback: hostKeyCallback,
		})
		if err != nil {
			return fmt.Errorf("SSH connection failed: %w", err)
		}
		defer client.Close()

		ref := imageRef(cfg)
		pullCmd := fmt.Sprintf("docker pull %s", ref)
		removeCmd := fmt.Sprintf("docker rm -f %s || true", cfg.Deploy.Container.Name)
		runCmd := buildDockerRunCmd(cfg, ref)

		commands := []string{pullCmd, removeCmd, runCmd}
		for _, command := range commands {
			session, err := client.NewSession()
			if err != nil {
				return fmt.Errorf("SSH session error: %w", err)
			}
			session.Stdout = os.Stdout
			session.Stderr = os.Stderr
			if err := session.Run(command); err != nil {
				session.Close()
				return fmt.Errorf("remote command failed (%q): %w", command, err)
			}
			session.Close()
		}
		return nil
	},
}

func hostKeyCallbackFor(host string, insecure bool) (ssh.HostKeyCallback, error) {
	if insecure {
		return ssh.InsecureIgnoreHostKey(), nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("locating home dir for known_hosts: %w", err)
	}
	knownHostsPath := filepath.Join(home, ".ssh", "known_hosts")
	cb, err := knownhosts.New(knownHostsPath)
	if err != nil {
		return nil, fmt.Errorf("loading known_hosts (%s): %w", knownHostsPath, err)
	}
	return cb, nil
}

func buildDockerRunCmd(cfg *Config, ref string) string {
	parts := []string{"docker", "run", "-d", "--name", cfg.Deploy.Container.Name}
	if cfg.Deploy.Container.Port != "" {
		parts = append(parts, "-p", cfg.Deploy.Container.Port)
	}
	if cfg.Deploy.Container.EnvFile != "" {
		parts = append(parts, "--env-file", cfg.Deploy.Container.EnvFile)
	}
	if cfg.Deploy.Container.Network != "" {
		parts = append(parts, "--network", cfg.Deploy.Container.Network)
	}
	for _, alias := range cfg.Deploy.Container.Aliases {
		parts = append(parts, "--network-alias", alias)
	}
	parts = append(parts, ref)
	return strings.Join(parts, " ")
}

func init() {
	rootCmd.AddCommand(releaseCmd)
}
