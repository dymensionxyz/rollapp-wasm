ROLLAPP_KEY_NAME_GENESIS="rol-user"
ROLLAPP_GENESIS_ADDR="$($EXECUTABLE keys show $ROLLAPP_KEY_NAME_GENESIS | grep "address:" | cut -d' ' -f3)"
ROLLAPP_CHAIN_ID="rollappwasm_1234-1"

SETTLEMENT_KEY_NAME_GENESIS="local-user"
SETTLEMENT_GENESIS_ADDR="$(dymd keys show $SETTLEMENT_KEY_NAME_GENESIS | grep "address:" | cut -d' ' -f3)"
SETTLEMENT_CHAIN_ID="dymension_100-1"

ICS20_CODE_ID="$($EXECUTABLE q wasm  list-code | grep "code_id:" | tail -n 1 | cut -d' ' -f3 | tr -d '"')"
CW20_CODE_ID=$((ICS20_CODE_ID - 1))
CW20_ADDR=$($EXECUTABLE q wasm list-contract-by-code $CW20_CODE_ID --output json | jq -r '.contracts[0]' )
ICS20_ADDR=$($EXECUTABLE q wasm list-contract-by-code $ICS20_CODE_ID --output json | jq -r '.contracts[0]' )

ICS20_PATH="ics20-hub"
version=ics20-1
transfer_port=transfer

# get ics20 wasm port
wasm_port=$($EXECUTABLE q wasm contract-state smart $ICS20_ADDR '{"port":{}}' | grep "port_id:" | cut -d' ' -f4)
echo "contract $ICS20_ADDR has wasm port: $wasm_port"

# Create channel 
rly paths new "$ROLLAPP_CHAIN_ID" "$SETTLEMENT_CHAIN_ID" "$ICS20_PATH" --src-port "$wasm_port" --dst-port "$transfer_port" --version "$version"
rly tx link "$ICS20_PATH" --src-port "$wasm_port" --dst-port "$transfer_port" --version "$version"

# make sure that channel has registered 
channel_id=$($EXECUTABLE q wasm contract-state smart $ICS20_ADDR '{"list_channels":{}}' | grep "channel_id:" | tail -n 1 | awk '{print $2}')

# encode transfer msg
TRANSFER_MSG=$(cat <<EOF
{
    "channel": "$channel_id",
    "remote_address": "$SETTLEMENT_GENESIS_ADDR"
}
EOF
)
encoded=$(echo "$TRANSFER_MSG" | jq -c . | base64)
SEND_MSG=$(cat <<EOF
{
    "send": {
        "contract": "$ICS20_ADDR",
        "amount": "100000",
        "msg": "$encoded"
    }
}
EOF
)

$EXECUTABLE tx wasm execute $CW20_ADDR "$SEND_MSG" --from rol-user --gas 50000000 --yes
sleep 5

#query balance
QUERY_MSG=$(cat <<EOF
{"balance":{"address":"$ROLLAPP_GENESIS_ADDR"}}
EOF
)
$EXECUTABLE q wasm contract-state smart $contract "$QUERY_MSG"

# balance of the hub will not change yet, we need to start relayer to make it work
dymd q bank balances $SETTLEMENT_GENESIS_ADDR

# start relayer
rly start
sleep 10

# check the balance in hub again
dymd q bank balances $SETTLEMENT_GENESIS_ADDR
