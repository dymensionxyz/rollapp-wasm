#!/bin/bash
tmp=$(mktemp)

set_denom() {
  base_denom=$1
  denom=$(echo "$base_denom" | sed 's/^.//')
  jq --arg base_denom $base_denom '.app_state.mint.params.mint_denom = $base_denom' "$GENESIS_FILE" >"$tmp" && mv "$tmp" "$GENESIS_FILE"
  jq --arg base_denom $base_denom '.app_state.staking.params.bond_denom = $base_denom' "$GENESIS_FILE" >"$tmp" && mv "$tmp" "$GENESIS_FILE"
  jq --arg base_denom $base_denom '.app_state.gov.deposit_params.min_deposit[0].denom = $base_denom' "$GENESIS_FILE" >"$tmp" && mv "$tmp" "$GENESIS_FILE"

  jq --arg base_denom $base_denom --arg denom $denom '.app_state.bank.denom_metadata = [
        {
            "base": $base_denom,
            "denom_units": [
                {
                    "aliases": [],
                    "denom": $base_denom,
                    "exponent": 0
                },
                {
                    "aliases": [],
                    "denom": $denom,
                    "exponent": 18
                }
            ],
            "description": "Denom metadata for Rollapp Wasm",
            "display": $denom,
            "name": $denom,
            "symbol": "WASM"
        }
    ]' "$GENESIS_FILE" > "$tmp" && mv "$tmp" "$GENESIS_FILE"

}

set_consensus_params() {
  BLOCK_SIZE="500000"
  jq --arg block_size "$BLOCK_SIZE" '.consensus_params["block"]["max_bytes"] = $block_size' "$GENESIS_FILE" >"$tmp" && mv "$tmp" "$GENESIS_FILE"
  jq --arg block_size "$BLOCK_SIZE" '.consensus_params["evidence"]["max_bytes"] = $block_size' "$GENESIS_FILE" >"$tmp" && mv "$tmp" "$GENESIS_FILE"
}

# ---------------------------- initial parameters ---------------------------- #
# Assuming 1,000,000 tokens
#half is staked
TOKEN_AMOUNT="1000000000000000000000000$BASE_DENOM"
STAKING_AMOUNT="500000000000000000000000$BASE_DENOM"

CONFIG_DIRECTORY="$ROLLAPP_HOME_DIR/config"
GENESIS_FILE="$CONFIG_DIRECTORY/genesis.json"
TENDERMINT_CONFIG_FILE="$CONFIG_DIRECTORY/config.toml"
APP_CONFIG_FILE="$CONFIG_DIRECTORY/app.toml"

# --------------------------------- run init --------------------------------- #
if ! command -v $EXECUTABLE >/dev/null; then
  echo "$EXECUTABLE does not exist"
  echo "please run make install"
  exit 1
fi

if [ -z "$ROLLAPP_CHAIN_ID" ]; then
  echo "ROLLAPP_CHAIN_ID is not set"
  exit 1
fi

# Verify that a genesis file doesn't exists for the dymension chain
if [ -f "$GENESIS_FILE" ]; then
  printf "\n======================================================================================================\n"
  echo "A genesis file already exists [$GENESIS_FILE]. building the chain will delete all previous chain data. continue? (y/n)"
  printf "\n======================================================================================================\n"
  read -r answer
  if [ "$answer" != "${answer#[Yy]}" ]; then
    rm -rf "$ROLLAPP_HOME_DIR"
  else
    exit 1
  fi
fi

# ------------------------------- init rollapp ------------------------------- #
$EXECUTABLE init "$MONIKER" --chain-id "$ROLLAPP_CHAIN_ID"

# ------------------------------- client config ------------------------------ #
$EXECUTABLE config keyring-backend test
$EXECUTABLE config chain-id "$ROLLAPP_CHAIN_ID"

# -------------------------------- app config -------------------------------- #
sed -i'' -e "s/^minimum-gas-prices *= .*/minimum-gas-prices = \"0$BASE_DENOM\"/" "$APP_CONFIG_FILE"
set_denom "$BASE_DENOM"
set_consensus_params
# --------------------- adding keys and genesis accounts --------------------- #
#local genesis account
$EXECUTABLE keys add "$KEY_NAME_ROLLAPP" --keyring-backend test
$EXECUTABLE add-genesis-account "$KEY_NAME_ROLLAPP" "$TOKEN_AMOUNT" --keyring-backend test

# set sequencer's operator address
operator_address=$($EXECUTABLE keys show "$KEY_NAME_ROLLAPP" -a --keyring-backend test --bech val)
jq --arg addr $operator_address '.app_state["sequencers"]["genesis_operator_address"] = $addr' "$GENESIS_FILE" > "$tmp" && mv "$tmp" "$GENESIS_FILE"

# set sequencer's operator address
operator_address=$($EXECUTABLE keys show "$KEY_NAME_ROLLAPP" -a --keyring-backend test --bech val)
jq --arg addr $operator_address '.app_state["sequencers"]["genesis_operator_address"] = $addr' "$GENESIS_FILE" >"$tmp" && mv "$tmp" "$GENESIS_FILE"

echo "Do you want to include staker on genesis? (Y/n) "
read -r answer
if [ ! "$answer" != "${answer#[Nn]}" ]; then
  set -x
  $EXECUTABLE gentx "$KEY_NAME_ROLLAPP" "$STAKING_AMOUNT" --chain-id "$ROLLAPP_CHAIN_ID" --keyring-backend test --home "$ROLLAPP_HOME_DIR"
  $EXECUTABLE collect-gentxs --home "$ROLLAPP_HOME_DIR"
  set +x
fi

$EXECUTABLE validate-genesis
