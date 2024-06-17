<!--
order: 3
-->

# Clients

## Command Line Interface (CLI)

The CLI has been updated with new queries and transactions for the `x/cron` module. View the entire list below.

### Queries

| Command                   | Subcommand | Arguments | Description                    |
| :------------------------ | :--------- | :-------- | :----------------------------- |
| `wasmrollappd query cron` | `params`   |           | Get Cron params                |
| `wasmrollappd query cron` | `crons`    |           | Get the list of the cronJobs   |
| `wasmrollappd query cron` | `cron`     | [id]      | Get the details of the cronJob |

### Transactions

| Command                | Subcommand        | Arguments                                          | Description                               |
| :--------------------- | :---------------- | :------------------------------------------------- | :---------------------------------------- |
| `wasmrollappd tx cron` | `register-cron`   | [name] [description] [contract_address] [json_msg] | Register the cron job                     |
| `wasmrollappd tx cron` | `update-cron-job` | [id] [contract_address] [json_msg]                 | update the cron job                       |
| `wasmrollappd tx cron` | `delete-cron-job` | [id] [contract_address]                            | delete the cron job for the contract      |
| `wasmrollappd tx cron` | `toggle-cron-job` | [id]                                               | Toggle the cron job for the given cron ID |
