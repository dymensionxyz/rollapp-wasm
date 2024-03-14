#!/bin/bash

ROLLAPP_KEY_NAME_GENESIS="rol-user"
ROLLAPP_GENESIS_ADDR="$(rollappd keys show $ROLLAPP_KEY_NAME_GENESIS | grep "address:" | cut -d' ' -f3)"

# Store code for cw20_base contract
rollappd tx wasm store ./scripts/bytecode/cw20_base.wasm --from rol-user --gas 5000000 --yes

sleep 5

CW20_CODE_ID="$(rollappd q wasm  list-code | grep "code_id:" | tail -n 1 | cut -d' ' -f3 | tr -d '"')"

# Instantiate contract
INIT_CW20=$(cat <<EOF
{
  "name": "My first token",
  "symbol": "test",
  "decimals": 6,
  "initial_balances": [{
    "address": "$ROLLAPP_GENESIS_ADDR",
    "amount": "100000000000"
  }]
}
EOF
)
rollappd tx wasm instantiate $CW20_CODE_ID "$INIT_CW20" --label test --no-admin --from $ROLLAPP_KEY_NAME_GENESIS --yes
CW20_ADDR=$(rollappd q wasm list-contract-by-code $CW20_CODE_ID --output json | jq -r '.contracts[0]' )
echo "Token contract deployed at: $CW20_ADDR"

# Query rol-user balances
QUERY_MSG=$(cat <<EOF
{"balance":{"address":"$ROLLAPP_GENESIS_ADDR"}}
EOF
)
balance=$(rollappd q wasm contract-state smart $CW20_ADDR "$QUERY_MSG" | grep "balance:" | cut -d' ' -f4 | tr -d '"')
echo "User $ROLLAPP_GENESIS_ADDR has balance $balance for contract $CW20_ADDR"


# Store code for ics20 contract
rollappd tx wasm store ./scripts/bytecode/cw20_ics20.wasm --from $ROLLAPP_KEY_NAME_GENESIS --gas 5000000 --yes
sleep 5
ICS20_CODE_ID="$(rollappd q wasm  list-code | grep "code_id:" | tail -n 1 | cut -d' ' -f3 | tr -d '"')"

INIT_ICS20=$(cat <<EOF
{
    "default_timeout":1000,
    "gov_contract":"$ROLLAPP_GENESIS_ADDR",
    "allowlist":[{
        "contract":"$CW20_ADDR"
    }]
} 
EOF
)
rollappd tx wasm instantiate $ICS20_CODE_ID "$INIT_ICS20" --label ics20 --no-admin --from rol-user --gas 50000000 --yes
ICS20_ADDR=$(rollappd q wasm list-contract-by-code $ICS20_CODE_ID --output json | jq -r '.contracts[0]' )

echo "ICS20 contract deployed at: $ICS20_ADDR"

# Query rol-user balances
# rollappd q wasm contract-state smart $contract '{"balance":{"address":"rol1h9htcc6hntfh02x5jrtkya6f3vzcycu27zm3um"}}'