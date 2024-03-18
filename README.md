<!-- markdownlint-disable MD033 -->
<h1 align="center">Dymension Rollapp</h1>
<!-- markdownlint-enable MD033 -->

# Rollappd - A template RollApp chain

This repository hosts `rollappd`, a template implementation of a dymension rollapp.

`rollappd` is an example of a working RollApp using `dymension-RDK` and `dymint`.

It uses Cosmos-SDK's [simapp](https://github.com/cosmos/cosmos-sdk/tree/main/simapp) as a reference, but with the following changes:

- minimal app setup
- wired IBC for [ICS 20 Fungible Token Transfers](https://github.com/cosmos/ibc/tree/main/spec/app/ics-020-fungible-token-transfer)
- Uses `dymint` for block sequencing and replacing `tendermint`
- Uses modules from `dymension-RDK` to sync with `dymint` and provide RollApp custom logic

## Overview

**Note**: Requires [Go 1.21](https://go.dev/)

## Quick guide

Get started with [building RollApps](https://docs.dymension.xyz/develop/get-started/setup)

## Installing / Getting started

Build and install the ```rollappd``` binary:

```shell
make install
```

### Initial configuration

export the following variables:

```shell
export ROLLAPP_CHAIN_ID="rollappwasm_1234-1"
export KEY_NAME_ROLLAPP="rol-user"
export DENOM="urax"
export MONIKER="$ROLLAPP_CHAIN_ID-sequencer"
```

if you want to change the max wasm size:

```shell
export MAX_WASM_SIZE=YOUR_MAX_WASM_SIZE
```

And initialize the rollapp:

```shell
sh scripts/init.sh
```

### Download cw20-ics20 smartcontract

Download cw20-ics20 smartcontract with a specific version:

```shell
sh scripts/download_release.sh v1.0.0
```

### Run rollapp

```shell
rollappd start
```

You should have a running local rollapp!

## Run a rollapp with local settlement node

### Run local dymension hub node

Follow the instructions on [Dymension Hub docs](https://docs.dymension.xyz/develop/get-started/run-base-layers) to run local dymension hub node

### Create sequencer keys

create sequencer key using `dymd`

```shell
dymd keys add sequencer --keyring-dir ~/.rollapp/sequencer_keys --keyring-backend test
SEQUENCER_ADDR=`dymd keys show sequencer --address --keyring-backend test --keyring-dir ~/.rollapp/sequencer_keys`
```

fund the sequencer account

```shell
dymd tx bank send local-user $SEQUENCER_ADDR 10000000000000000000000adym --keyring-backend test --broadcast-mode block --fees 20000000000000adym -y
```

### Register rollapp on settlement

```shell
sh scripts/settlement/register_rollapp_to_hub.sh
```

### Register sequencer for rollapp on settlement

```shell
sh scripts/settlement/register_sequencer_to_hub.sh
```

### Configure the rollapp

Modify `dymint.toml` in the chain directory (`~/.rollapp/config`)
set:

```shell
settlement_layer = "dymension"
gas_prices = "0.025adym"
```

### Run rollapp locally

```shell
rollappd start
```

## Setup IBC between rollapp and local dymension hub node

### Install dymension relayer

```shell
git clone https://github.com/dymensionxyz/go-relayer.git --branch v0.2.0-v2.3.1-relayer
cd go-relayer && make install
```

### Establish IBC channel

while the rollapp and the local dymension hub node running, run:

```shell
sh scripts/ibc/setup_ibc.sh
```

After successful run, the new established channels will be shown

### run the relayer

```shell
rly start hub-rollapp
```

### Deploy the installed contract

```shell
sh scripts/wasm/deploy_contract.sh
```

### Make the ibc transfer

```shell
sh scripts/wasm/ibc_transfer.sh
```

## Developers guide

TODO
