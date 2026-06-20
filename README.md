# depx

A minimal Docker deployment tool that deploys Docker images to remote servers over SSH. Replaces GitLab CI deployment pipelines.

## Features

- Single binary, no external dependencies beyond Go runtime
- Simple, explicit CLI using Cobra
- SSH-based deployment via `golang.org/x/crypto/ssh`
- No background workers, queues, or complex pipelines

## Tech Stack

- Go
- [Cobra](https://github.com/spf13/cobra) for CLI
- [golang.org/x/crypto/ssh](https://pkg.go.dev/golang.org/x/crypto/ssh) for SSH connections

## Installation

### From Source
```bash
git clone <repo-url> depx
cd depx
go build -o depx .
```

## Usage

```bash
depx [command] [flags]
```

### Commands

| Command  | Description                          |
|----------|--------------------------------------|
| init     | Create a depx.yaml configuration file|
| build    | Build Docker image for deployment    |
| push     | Push Docker image to registry        |
| release  | Deploy image to remote server via SSH|

### Workflow

```bash
# 1. Initialize configuration
depx init

# 2. Build the Docker image
depx build

# 3. Push to registry
depx push

# 4. Deploy to remote server
depx release
```

### Configuration

depx reads a `depx.yaml` file in the current directory:

```yaml
app:
  name: my-app          # application name
  dockerfile: Dockerfile
  context: .

image:
  name: my-app          # image repository name
  tag: latest
  registry: docker.io   # registry host

deploy:
  host: your-server.com
  port: 22
  user: deploy
  key: ~/.ssh/id_rsa
  container:
    name: my-app        # container name on remote
    port: "3000:3000"   # host:container port mapping
    env_file: .env
```

Values support `${VAR}` syntax for environment variable substitution
(e.g. `registry: ${DOCKER_REGISTRY}`).

## Project Structure

```
.
├── main.go          # Entry point
├── cmd/
│   ├── root.go      # Root Cobra command
│   ├── config.go    # Shared config struct and helpers
│   ├── init.go      # init command
│   ├── build.go     # build command
│   ├── push.go      # push command
│   └── release.go   # release command
├── go.mod
└── go.sum
```

## Development

```bash
# Add dependencies
go mod tidy

# Build
go build -o depx .

# Run
./depx --help
```
