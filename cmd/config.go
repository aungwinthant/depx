package cmd

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App    AppConfig    `yaml:"app"`
	Image  ImageConfig  `yaml:"image"`
	Deploy DeployConfig `yaml:"deploy"`
}

type AppConfig struct {
	Name       string `yaml:"name"`
	Dockerfile string `yaml:"dockerfile"`
	Context    string `yaml:"context"`
}

type ImageConfig struct {
	Name     string `yaml:"name"`
	Tag      string `yaml:"tag"`
	Registry string `yaml:"registry"`
}

type DeployConfig struct {
	Host      string          `yaml:"host"`
	Port      int             `yaml:"port"`
	User      string          `yaml:"user"`
	Key       string          `yaml:"key"`
	Container ContainerConfig `yaml:"container"`
}

type ContainerConfig struct {
	Name    string `yaml:"name"`
	Port    string `yaml:"port"`
	EnvFile string `yaml:"env_file"`
}

func loadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	data = []byte(os.ExpandEnv(string(data)))
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func expandPath(path string) string {
	if strings.HasPrefix(path, "~/") {
		home, _ := os.UserHomeDir()
		return filepath.Join(home, path[2:])
	}
	return path
}

func imageRef(cfg *Config) string {
	return cfg.Image.Registry + "/" + cfg.Image.Name + ":" + cfg.Image.Tag
}

func runCmd(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
