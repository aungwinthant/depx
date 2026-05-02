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
| init     | Initialize depx configuration        |
| build    | Build Docker image for deployment    |
| push     | Push Docker image to registry        |
| release  | Deploy image to remote server via SSH|

Run `depx [command] --help` for more details on each command.

## Project Structure

```
.
├── main.go          # Entry point
├── cmd/
│   ├── root.go      # Root Cobra command
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

> Note: Business logic for commands is not yet implemented. This is the initial CLI scaffold.
