<!--
order: 1
-->

# Concepts

## Cron

The cron module provides functionality for scheduling and executing tasks, including executing sudo contract calls during specific phases, such as begin blockers. By integrating scheduled contract executions, `x/cron` enhances the functionality of smart contracts, ensuring critical operations are performed automatically and reliably.
Developers can register their contracts if the address is whitelisted in the module parameters. Contracts can be deregistered if the cron job is no longer needed

### Registering a contract

```console
foo@bar:~$ wasmrollappd tx cron register-contract [game name] [contract address] [game type]
```

e.g

```console
foo@bar:~$ wasmrollappd tx cron register-contract coin-flip rol14hj2tavq8f.... 1 100000000awasm --from cooluser --chain-id test-1
```

In the above tx -

- `game name` - name of the game contract
- `contract address` - CosmWasm contract address.
- `game type` - 1 for single player, 2 for multi player, 3 for both single and multiplayer

> Note : only the security address authorized can register the contract

### De-Registering a contract

```console
foo@bar:~$ wasmrollappd tx cron de-register-contract [game id]
```

e.g

```console
foo@bar:~$ wasmrollappd tx cron de-register-contract 2 100000000awasm --from cooluser --chain-id test-1
```

In the above tx -

- `game id` - id of the game contract to de-register

> Note : only the security address and contract admin are authorized can de-register the contract
