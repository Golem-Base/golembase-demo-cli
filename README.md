# Golem Base CLI

The Golem Base CLI is a command-line interface written in Go for interacting with the Golem Base storage system (golembase-op-geth). It provides tools for account management, entity creation, and querying the storage system.

## Table of Contents

-   [Requirements](#requirements)
-   [Installation](#installation)
-   [Configuration](#configuration)
-   [Funding Your Account](#funding-your-account)
-   [Available Commands](#available-commands)
    -   [Account Management](#account-management)
    -   [Entity Management](#entity-management)
    -   [Query Operations](#query-operations)
    -   [Entity Content Display](#entity-content-display)
-   [Usage Examples](#usage-examples)
    -   [Detailed Use Cases](./use-cases/)
-   [Development](#development)

## Requirements

-   Go 1.23.5 or later
-   A running Golem Base node (default: https://api.golembase.demo.golem-base.io/)

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

-   On macOS: `~/Library/Application Support/golembase/`
-   On Linux: `~/.config/golembase/`
-   On Windows: `%LOCALAPPDATA%\golembase\`

The configuration directory stores:

-   Private keys for account management
-   Node connection settings
-   Other persistent settings

## Funding Your Account

Golem Base operations (like creating or updating data) require a small amount of ETH to pay for transaction gas fees on the network.

-   **Faucet:** Obtain test ETH for the Golem Base Demo network by visiting the official faucet at `https://faucet.golembase.demo.golem-base.io/`.
-   **Process:** You will need your Golem Base account address (which you can get from `account create` or `account import`). Paste this address into the faucet interface to receive test ETH.
-   **Verify:** After using the faucet, check your balance using `golembase-demo-cli account balance`.

## Available Commands

### Account Management

-   **Create a new account.**

    ```bash
    golembase-demo-cli account create
    ```

    -   Generates a new private key if one doesn't exist.
    -   Saves it to the XDG config directory (e.g., `~/Library/Application Support/golembase/private.key` on macOS).
    -   Displays the generated Ethereum address.

-   **Check the ETH balance** of the configured account.

    ```bash
    golembase-demo-cli account balance
    ```

    -   Displays account address and current ETH balance.
    -   Accepts optional `--node-url <url>` flag.

-   **Import an account** using a hex private key.
    ```bash
    golembase-demo-cli account import
    ```
    -   Requires the `--privatekey <hex-key>` flag (alias `--key`).
    -   Overwrites any existing private key.

### Entity Management

In Golem Base, an "entity" is simply a piece of data (payload) stored on the network, identified by a unique key. You can manage these entities using the following commands:

-   **Create a new entity** in Golem Base.

    ```bash
    golembase-demo-cli entity create
    ```

    -   Optional flags:
        -   `--data <payload>`: Custom payload data (Default: `"this is a test"`).
        -   `--ttl <blocks>`: Custom time-to-live value in blocks (Default: `100`).
    -   Returns the entity key upon success.

-   **Update an existing entity's** payload and/or TTL.

    ```bash
    golembase-demo-cli entity update
    ```

    -   Requires the `--key <entity-key>` flag.
    -   Optional flags:
        -   `--data <payload>`: New payload data.
        -   `--ttl <blocks>`: New time-to-live value.

-   **Delete an existing entity** by its key.
    ```bash
    golembase-demo-cli entity delete
    ```
    -   Requires the `--key <entity-key>` flag.
    -   Note: May sometimes fail with `tx failed`.

### Query Operations

-   **Query the storage system** for entities based on annotations.
    ```bash
    golembase-demo-cli query '<query-string>'
    ```
    -   Takes a `<query-string>` argument (e.g., `'foo = "bar"'`).
    -   String comparisons use a single `=`.
    -   For detailed query syntax, see the [Query Language Support section](https://github.com/Golem-Base/golembase-op-geth/blob/main/golem-base/README.md#api-functionality).

### Entity Content Display

-   **Display the raw payload content** of a specified entity.
    ```bash
    golembase-demo-cli cat <entity-key>
    ```
    -   Takes an `<entity-key>` argument.

## Usage Examples

Here's a quick 3-step guide to get started using the `golembase-demo-cli`:

1.  **Create a new account:**

    ```bash
    golembase-demo-cli account create
    ```

    _(Remember to fund this account using the faucet mentioned in the [Funding Your Account](#funding-your-account) section)._

2.  **Create a new entity:**

    ```bash
    golembase-demo-cli entity create --data "custom data" --ttl 200
    ```

    _(This command will output an `entity-key` upon success)._

3.  **Display entity payload:**
    ```bash
    golembase-demo-cli cat <entity-key>
    ```
    _(Replace `<entity-key>` with the key obtained from the previous step)._

---

The `golembase-demo-cli` simplifies the underlying JSON-RPC interactions for easy experimentation. For more detailed examples illustrating how real-world applications might integrate with Golem Base directly via its RPC interface or through software libraries, please refer to our dedicated use case documents:

-   [See Detailed Use Cases](./use-cases/)

## Development

This project is written in Go and uses the following main dependencies:

-   `github.com/urfave/cli/v2` for command-line interface
-   `github.com/ethereum/go-ethereum` for Ethereum interaction
-   `github.com/adrg/xdg` for XDG base directory support
-   `github.com/dustin/go-humanize` for human-readable output

To build the project:

```bash
go build
```
