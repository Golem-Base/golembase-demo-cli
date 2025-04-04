# Golem Base CLI

The Golem Base CLI is a command-line interface written in Go for interacting with the Golem Base storage system (golembase-op-geth). It provides tools for account management, entity creation, and querying the storage system.

## Requirements

- Go 1.23.5 or later
- A running Golem Base node (default: https://api.golembase.demo.golem-base.io/)

## Installation

### Using Homebrew (macOS)

```bash
brew install Golem-Base/demo/golembase-demo-cli
```

### Using Linux Binaries

Download the appropriate binary for your system from the [GitHub releases page](https://github.com/Golem-Base/golembase-demo-cli/releases/tag/v0.0.2).

### Using Go Install

```bash
go install github.com/Golem-Base/golembase-demo-cli@latest
```

## Configuration

The CLI follows the XDG Base Directory Specification for storing configuration files:
- On macOS: `~/Library/Application Support/golembase/`
- On Linux: `~/.config/golembase/`
- On Windows: `%LOCALAPPDATA%\golembase\`

The configuration directory stores:
- Private keys for account management
- Node connection settings
- Other persistent settings

## Available Commands

### Account Management

- `account create`: Creates a new account
  - Generates a new private key
  - Saves it to the XDG config directory (e.g., `~/Library/Application Support/golembase/private.key` on macOS)
  - Displays the generated Ethereum address

- `account balance`: Checks account balance
  - Displays account address and current ETH balance
  - Connects to the configured node to fetch the latest balance

### Entity Management

- `entity create`: Creates a new entity in Golem Base
  - Creates entity with default data and TTL (100 blocks)
  - Signs and submits transaction to the node
  - Optional flags:
    - `--node-url`: Specify different node URL
    - `--data`: Custom payload data
    - `--ttl`: Custom time-to-live value in blocks

### Query Operations

- `query`: Commands for querying the storage system
  - Execute custom queries using the Golem Base query language
  - Search entities by annotations
  - Retrieve entity metadata
  - For detailed query syntax and examples, see the [Query Language Support section](https://github.com/Golem-Base/golembase-op-geth/blob/main/golem-base/README.md#api-functionality)

### Entity Content Display

- `cat`: Display entity payload content
  - Similar to Unix `cat` command
  - Dumps the raw payload data of a specified entity
  - Useful for viewing the contents of stored entities

## Usage Examples

1. Create a new account:
```bash
golembase-demo-cli account create
```

2. Create a new entity:
```bash
golembase-demo-cli entity create --data "custom data" --ttl 200
```

3. Display entity payload:
```bash
golembase-demo-cli cat <entity-key>
```

## Development

This project is written in Go and uses the following main dependencies:
- `github.com/urfave/cli/v2` for command-line interface
- `github.com/ethereum/go-ethereum` for Ethereum interaction
- `github.com/adrg/xdg` for XDG base directory support
- `github.com/dustin/go-humanize` for human-readable output

To build the project:
```bash
go build
```
