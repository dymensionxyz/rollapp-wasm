#!/bin/bash

rollappd tx wasm store ./scripts/bytecode/cw20_base.wasm --from rol-user --gas 5000000 --yes

sleep 5

rollappd tx wasm instantiate 1 '{"name":"test","symbol":"test","decimals":6,"initial_balances":[{"address":"rol1h9htcc6hntfh02x5jrtkya6f3vzcycu27zm3um","amount":"100000000"}]}' --label test --no-admin --from rol-user --yes

contract=$(rollappd q wasm list-contract-by-code 1 --output json | jq -r '.contracts[0]' )

echo "Token contract deployed at: $contract"

# Query rol-user balances
# rollappd q wasm contract-state smart $contract '{"balance":{"address":"rol1h9htcc6hntfh02x5jrtkya6f3vzcycu27zm3um"}}'