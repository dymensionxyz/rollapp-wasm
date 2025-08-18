# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

This is `rollapp-wasm`, a template implementation of a Dymension RollApp featuring:
- CosmWasm smart contract execution
- IBC (Inter-Blockchain Communication) for token transfers
- Dymint for block sequencing (replacing Tendermint)
- Integration with Dymension Hub as settlement layer
- Support for multiple Data Availability layers (Celestia, Mock, Avail, WeaveVM)

## Build Commands

### Primary Build & Install
```bash
# Set BECH32_PREFIX before building (required)
export BECH32_PREFIX=rol

# Build and install the rollapp-wasm binary
make install BECH32_PREFIX=$BECH32_PREFIX

# Build only (output to ./build/rollapp-wasm)
make build BECH32_PREFIX=$BECH32_PREFIX
```

### Testing
```bash
# Run all tests with race condition detection
go test -race ./...

# Run tests with coverage
go test -race -coverprofile=coverage.txt ./...

# Run specific package tests
go test -v ./x/callback/...
go test -v ./x/cwerrors/...
```

### Linting & Formatting
```bash
# Run golangci-lint if installed
golangci-lint run

# Format code with gofumpt
gofumpt -w .
```

### Protocol Buffers
```bash
# Generate protobuf files (uses Docker)
make proto-gen

# Clean proto generation container
make proto-clean
```

### Genesis File Generation
```bash
# Generate genesis file for mainnet
make generate-genesis env=mainnet

# Generate genesis file for testnet  
make generate-genesis env=testnet

# With specific DRS version
make generate-genesis env=mainnet DRS_VERSION=10
```

## Architecture

### Core Modules (`x/`)
- **callback**: Handles CosmWasm contract callbacks and fee management
  - Manages callback execution, gas fees, and sudo messages
  - Key types: `Callback`, `CallbackFees`, `SudoMsg`
  
- **cwerrors**: Custom error handling for CosmWasm contracts
  - Manages error subscriptions and sudo error messages
  - Provides standardized error reporting to contracts

- **wasm**: Custom authorization types for CosmWasm operations

### Application Structure (`app/`)
- **app.go**: Main application initialization and module wiring
- **ante.go**: Custom ante handler with fee decorators
- **wasm.go**: CosmWasm integration and configuration
- **upgrades/**: Protocol upgrade handlers (DRS-2 through DRS-10)
  - Each upgrade in separate directory with constants and upgrade logic

### Key Dependencies
- CosmWasm/wasmd v0.33.0 for smart contract execution
- Dymension-RDK v1.10.0 for RollApp modules
- Dymint v1.5.0 for consensus
- Cosmos SDK v0.46.16
- IBC-Go v6.3.0

### Configuration Files
- `~/.rollapp-wasm/config/dymint.toml`: Dymint consensus configuration
- `~/.rollapp-wasm/config/app.toml`: Application configuration
- `~/.rollapp-wasm/config/genesis.json`: Chain genesis state

## Common Development Workflows

### Local Development Setup
```bash
# Set environment variables
export EXECUTABLE="rollapp-wasm"
export CELESTIA_NETWORK="mock"
export BECH32_PREFIX="rol"
export ROLLAPP_CHAIN_ID="rollappwasm_1234-1"
export BASE_DENOM="awsm"
export ROLLAPP_HOME_DIR="$HOME/.rollapp-wasm"

# Initialize the rollapp
sh scripts/init.sh

# Start the rollapp
rollapp-wasm start
```

### Working with Smart Contracts
```bash
# Download cw20-ics20 contract
sh scripts/download_release.sh v1.0.0

# Deploy contract (requires running chain)
sh scripts/wasm/deploy_contract.sh

# Execute IBC transfer
sh scripts/wasm/ibc_transfer.sh
```

### Settlement Layer Integration
```bash
# Generate denomination metadata
sh scripts/settlement/generate_denom_metadata.sh

# Add genesis accounts
sh scripts/settlement/add_genesis_accounts.sh

# Register rollapp on hub
sh scripts/settlement/register_rollapp_to_hub.sh

# Register sequencer
sh scripts/settlement/register_sequencer_to_hub.sh
```

### IBC Setup
```bash
# Setup IBC channel between rollapp and hub
sh scripts/ibc/setup_ibc.sh

# Start relayer
rly start hub-rollapp
```

## Important Notes

- Go 1.21+ required
- Requires `dasel` and `jq` for configuration scripts
- Always set `BECH32_PREFIX` environment variable before building
- DRS (Dymension Rollapp Standard) version is currently 10
- Smart contract size limits can be adjusted via start flags:
  - `--max-label-size` (default: 128)
  - `--max-wasm-size` (default: 819200)
  - `--max-proposal-wasm-size` (default: 3145728)

## Module-Specific Testing

### Callback Module
```bash
go test -v ./x/callback/...
go test -v ./x/callback/keeper/...
```

### CWErrors Module
```bash
go test -v ./x/cwerrors/...
go test -v ./x/cwerrors/keeper/...
```

### E2E Testing
```bash
go test -v ./e2e/testing/...
```

## Upgrade Handlers

The rollapp supports protocol upgrades from DRS-2 through DRS-10. Each upgrade handler is in `app/upgrades/drs-X/`:
- Constants define upgrade name and version
- Upgrade logic handles state migrations
- Test files verify upgrade behavior

Current version: DRS-10 (as defined in Makefile)