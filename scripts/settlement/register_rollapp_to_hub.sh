#!/bin/bash

KEYRING_PATH="$HOME/.rollapp/sequencer_keys"
MAX_SEQUENCERS=5

# this account must be whitelisted on the hub for permissioned deployment setup
DEPLOYER="local-user"

#Register rollapp
set -x
dymd tx rollapp create-rollapp "$ROLLAPP_CHAIN_ID" "$MAX_SEQUENCERS" '{"Addresses":[]}' \
  "$DENOM_METADATA_PATH" \
  --genesis-accounts-path "$GENESIS_ACCOUNTS_PATH" \
  --from "$DEPLOYER" \
  --keyring-backend test \
  --keyring-dir "$KEYRING_PATH" \
  --broadcast-mode block \
  --fees 1dym \
  -y
set +x
