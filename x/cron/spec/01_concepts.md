<!--
order: 1
-->

# Concepts

## Cron

The cron module provides functionality for scheduling and executing tasks, including executing sudo contract calls during specific phases, such as begin blockers. By integrating scheduled contract executions, `x/cron` enhances the functionality of smart contracts, ensuring critical operations are performed automatically and reliably.
Developers can register their contracts as a cron job if the address is whitelisted in the module parameters. Cron job can be deleted/updated if is no longer needed

### Registering a Cron

```console
foo@bar:~$ wasmrollappd tx cron register-cron [name] [description] [contract_address] [json_msg]
```

e.g

```console
foo@bar:~$ wasmrollappd tx cron register-cron cronjob1 "the is the 1st cron job" rol14hj2tavq8f.... {"msg_cron":{}} 100000000awasm --from cooluser --chain-id test-1
```

In the above tx -

- `name` - name of the cron job
- `description` - description of the cron job
- `contract address` - CosmWasm contract address.
- `json_msg` - sudo msg of the contract in json format

> Note : only the security address authorized can register the contract

### Delete cron job

```console
foo@bar:~$ wasmrollappd tx cron update-cron-job [id] [contract_address] [json_msg]
```

e.g

```console
foo@bar:~$ wasmrollappd tx cron update-cron-job 1 rol14hj2tavq8f.... {"msg_new_cron":{}} 100000000awasm --from cooluser --chain-id test-1
```

In the above tx -

- `id` - id of the cron job
- `contract address` - CosmWasm contract address.
- `json_msg` - sudo msg of the contract in json format

> Note : only the security address and contract admin are authorized can update the cron job

### Update cron job

```console
foo@bar:~$ wasmrollappd tx cron delete-cron-job [id] [contract_address]
```

e.g

```console
foo@bar:~$ wasmrollappd tx cron delete-cron-job 1 rol14hj2tavq8f.... 100000000awasm --from cooluser --chain-id test-1
```

In the above tx -

- `id` - id of the cron job
- `contract address` - CosmWasm contract address.

> Note : only the security address and contract admin are authorized can delete the cron job
