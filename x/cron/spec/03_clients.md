<!--
order: 3
-->

# Clients

## Command Line Interface (CLI)

The CLI has been updated with new queries and transactions for the `x/cron` module. View the entire list below.

### Queries

| Command                   | Subcommand              | Arguments | Description                                       |
| :------------------------ | :---------------------- | :-------- | :------------------------------------------------ |
| `wasmrollappd query cron` | `params`                |           | Get Cron params                                   |
| `wasmrollappd query cron` | `whitelisted-contracts` |           | Get the list of Whitelisted Contracts for cronjob |

### Transactions

| Command                | Subcommand             | Arguments                                  | Description                            |
| :--------------------- | :--------------------- | :----------------------------------------- | :------------------------------------- |
| `wasmrollappd tx cron` | `register-contract`    | [game-name] [contract-address] [game-type] | Register the contract for cron job     |
| `wasmrollappd tx cron` | `de-register-contract` | [game-id]                                  | De-Register the contract from cron job |
