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

### Global Flags

| Flag                          | Default     | Description                                                       |
|-------------------------------|-------------|-------------------------------------------------------------------|
| `-c, --config <path>`         | `depx.yaml` | Path to the config file (e.g. `depx.prod.yaml`). Applies to all commands. |
| `--insecure-skip-host-key`    | `false`     | Skip SSH host key verification. Use only as a last resort.        |

### Commands

| Command  | Description                          |
|----------|--------------------------------------|
| init     | Create a depx configuration file     |
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

Use per-environment configs by passing `--config`/`-c`:

```bash
depx -c depx.staging.yaml build
depx -c depx.staging.yaml push
depx -c depx.staging.yaml release

depx -c depx.prod.yaml release
```

### Exit Codes

`build`, `push`, and `release` return a non-zero exit code on any failure
(config load, build/push error, SSH failure, remote command failure). This makes
`set -e`, CI pipelines, and wrapper scripts reliable.

### Configuration

depx reads a config file (default `depx.yaml`, overridable with `--config`) in
the current directory:

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
  insecure_skip_host_key: false   # set true to opt out of known_hosts (insecure)
  container:
    name: my-app                  # container name on remote
    port: "3000:3000"             # host:container port mapping
    env_file: .env                # passed to docker run as --env-file
    network: ""                   # join an existing docker network, e.g. myapp_default
    aliases: []                   # network aliases, e.g. [backend]
```

Values support `${VAR}` syntax for environment variable substitution
(e.g. `registry: ${DOCKER_REGISTRY}`).

### SSH Host Key Verification

By default, `depx release` verifies the remote host's SSH key against
`~/.ssh/known_hosts`. If the host key is unknown, release fails rather than
connecting silently. Seed the entry first:

```bash
ssh-keyscan -H your-server.com >> ~/.ssh/known_hosts
```

To opt out for an emergency or throwaway host, either pass the global flag:

```bash
depx release --insecure-skip-host-key
```

or set `deploy.insecure_skip_host_key: true` in the config file. The flag takes
precedence over the config value; both default to secure.

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
