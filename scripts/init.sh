#!/bin/bash
tmp=$(mktemp)

set_denom() {
  denom=$1
  jq --arg denom $denom '.app_state.mint.params.mint_denom = $denom' "$GENESIS_FILE" >"$tmp" && mv "$tmp" "$GENESIS_FILE"
  jq --arg denom $denom '.app_state.staking.params.bond_denom = $denom' "$GENESIS_FILE" >"$tmp" && mv "$tmp" "$GENESIS_FILE"
  jq --arg denom $denom '.app_state.gov.deposit_params.min_deposit[0].denom = $denom' "$GENESIS_FILE" >"$tmp" && mv "$tmp" "$GENESIS_FILE"
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
