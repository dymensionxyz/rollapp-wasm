#!/bin/bash

rollappd tx wasm store ./scripts/bytecode/cw20_ics20.wasm --from rol-user --gas 5000000 --yes

sleep 5

rollappd tx wasm instantiate 2 '{"default_timeout":1000,"gov_contract":"rol1h9htcc6hntfh02x5jrtkya6f3vzcycu27zm3um","allowlist":[{"contract":"rol14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9syy8zyg"}]}' --label ics20 --no-admin --from rol-user --gas 50000000 --yes
contract=$(rollappd q wasm list-contract-by-code 2 --output json | jq -r '.contracts[0]' )

echo "ICS20 contract deployed at: $contract"

# Query rol-user balances
# rollappd q wasm contract-state smart $contract '{"balance":{"address":"rol1h9htcc6hntfh02x5jrtkya6f3vzcycu27zm3um"}}'