# Dymension WASM Rollapp

## rollappd - A template WASM RollApp chain

This repository hosts `rollappd`, a template implementation of a dymension rollapp with `WASM` execution layer.

`rollappd` is an example of a working RollApp using `dymension-RDK` and `dymint`.

It uses Cosmos-SDK's [simapp](https://github.com/cosmos/cosmos-sdk/tree/main/simapp) as a reference, but with the following changes:

- minimal app setup
- wired IBC for [ICS 20 Fungible Token Transfers](https://github.com/cosmos/ibc/tree/main/spec/app/ics-020-fungible-token-transfer)
- Uses `dymint` for block sequencing and replacing `tendermint`
- Uses modules from `dymension-RDK` to sync with `dymint` and provide RollApp custom logic

## Overview

**Note**: Requires [Go 1.20](https://go.dev/)

## Quick guide

Get started with [building RollApps](https://docs.dymension.xyz/develop/get-started/setup)

## Installing / Getting started

Build and install the ```rollapp-wasm``` binary:

```shell
make install
```

### Initial configuration

export the following variables:

```shell
export EXECUTABLE="rollapp-wasm"
export ROLLAPP_CHAIN_ID="rollappwasm_1234-1"
export KEY_NAME_ROLLAPP="rol-user"
export BASE_DENOM="awsm"
export DENOM=$(echo "$BASE_DENOM" | sed 's/^.//')
export MONIKER="$ROLLAPP_CHAIN_ID-sequencer"

export ROLLAPP_HOME_DIR="$HOME/.rollapp-wasm"
export ROLLAPP_SETTLEMENT_INIT_DIR_PATH="${ROLLAPP_HOME_DIR}/init"
```

And initialize the rollapp:

```shell
sh scripts/init.sh
```

You can find out in <https://github.com/CosmWasm/wasmd#compile-time-parameters> that:

There are a few variables was allow blockchains to customize at compile time. If you build your own chain and import x/wasm, you can adjust a few items via module parameters, but a few others did not fit in that, as they need to be used by stateless ValidateBasic(). Thus, we made them as flags and set them in start.go so that they can be overridden on your custom chain.

```shell
rollapp-wasm start --max-label-size 64 --max-wasm-size 2048000 --max-wasm-proposal-size 2048000
```

Those flags are optional, the default value was set as:

```go
wasmtypes.MaxLabelSize          = 128
wasmtypes.MaxWasmSize           = 819200
wasmtypes.MaxProposalWasmSize   = 3145728
```

### Download cw20-ics20 smartcontract

Download cw20-ics20 smartcontract with a specific version:

```shell
bash scripts/download_release.sh v1.0.0
```

### Run rollapp

```shell
rollapp-wasm start
```

You should have a running local rollapp!

## Run a rollapp with a settlement node

### Run local dymension hub node

Follow the instructions on [Dymension Hub docs](https://docs.dymension.xyz/develop/get-started/run-base-layers) to run local dymension hub node

all scripts are adjusted to use local hub node that's hosted on the default port `localhost:36657`.

configuration with a remote hub node is also supported, the following variables must be set:

```shell
export HUB_RPC_ENDPOINT="http://localhost"
export HUB_RPC_PORT="36657" # default: 36657
export HUB_RPC_URL="http://${HUB_RPC_ENDPOINT}:${HUB_RPC_PORT}"
export HUB_CHAIN_ID="dymension_100-1"

dymd config chain-id ${HUB_CHAIN_ID}
dymd config node ${HUB_RPC_URL}

export HUB_KEY_WITH_FUNDS="hub-user" # This key should exist on the keyring-backend test
```

### Create sequencer keys

create sequencer key using `dymd`

```shell
dymd keys add sequencer --keyring-dir ~/.rollapp-wasm/sequencer_keys --keyring-backend test
SEQUENCER_ADDR=`dymd keys show sequencer --address --keyring-backend test --keyring-dir ~/.rollapp-wasm/sequencer_keys`
```

fund the sequencer account (if you're using a remote hub node, you must fund the sequencer account or you must have an account with enough funds in your keyring)

```shell
# this will retrieve the min bond amount from the hub
# if you're using an new address for registering a sequencer,
# you have to account for gas fees so it should the final value should be increased
BOND_AMOUNT="$(dymd q sequencer params -o json --node ${HUB_RPC_URL} | jq -r '.params.min_bond.amount')$(dymd q sequencer params -o jsono | jq -r '.params.min_bond.denom')"
echo $BOND_AMOUNT

# Extract the numeric part
NUMERIC_PART=$(echo $BOND_AMOUNT | sed 's/adym//')

# Add 100000000000000000000 for fees
NEW_NUMERIC_PART=$(echo "$NUMERIC_PART + 100000000000000000000" | bc)

# Append 'adym' back
TRANSFER_AMOUNT="${NEW_NUMERIC_PART}adym"

dymd tx bank send $HUB_KEY_WITH_FUNDS $SEQUENCER_ADDR ${TRANSFER_AMOUNT} --keyring-backend test --broadcast-mode block --fees 1dym -y --node ${HUB_RPC_URL} --chain-id ${HUB_CHAIN_ID}
```

### Generate denommetadata

```shell

sh scripts/settlement/generate_denom_metadata.sh
```

### Add genesis accounts

```shell
sh scripts/settlement/add_genesis_accounts.sh
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

Modify `dymint.toml` in the chain directory (`${ROLLAPP_HOME_DIR}/config`)

set:

```shell
sed -i 's/settlement_layer.*/settlement_layer = "dymension"/' ${ROLLAPP_HOME_DIR}/config/dymint.toml
sed -i '/node_address =/c\node_address = '\"$HUB_RPC_URL\" "${ROLLAPP_HOME_DIR}/config/dymint.toml"
sed -i '/rollapp_id =/c\rollapp_id = '\"$ROLLAPP_CHAIN_ID\" "${ROLLAPP_HOME_DIR}/config/dymint.toml"
```

### Update the Genesis file to include the denommetadata, genesis accounts, module account and elevated accounts

```shell
sh scripts/update_genesis_file.sh
```

### Update the Genesis file to include the denommetadata, genesis accounts, module account and elevated accounts 

```shell
# this script automatically adds 2 vesting accounts, adjust the timestampts to your liking or skip this step
sh scripts/add_vesting_accounts_to_genesis_file.sh
```

### Run rollapp locally

```shell
rollappd start
```

or as a systemd service:

```shell
sudo tee /etc/systemd/system/rollapp-wasm.service > /dev/null <<EOF
[Unit] 
Description=rollapp-wasm
After=network.target 
[Service] 
Type=simple
User=$USER
ExecStart=$(which rollapp-wasm) start
Restart=on-failure
RestartSec=10
LimitNOFILE=65535
[Install]
WantedBy=multi-user.target
EOF
sudo systemctl daemon-reload
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

or as a systemd service:

```shell
sudo tee /etc/systemd/system/relayer.service > /dev/null <<EOF
[Unit]
Description=rollapp
After=network.target
[Service]
Type=simple
User=$USER
ExecStart=$(which rly) start hub-rollapp
Restart=on-failure
RestartSec=10
LimitNOFILE=65535
[Install]
WantedBy=multi-user.target
EOF
```


## Developers guide

TODO
