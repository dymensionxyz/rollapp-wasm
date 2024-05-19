<!--
order: 3
-->

# Clients

## Command Line Interface (CLI)

The CLI has been updated with new queries and transactions for the `x/cron` module. View the entire list below.

### Queries

| Command           | Subcommand              | Arguments | Description                                       |
| :---------------- | :---------------------- | :-------- | :------------------------------------------------ |
| `aibd query cron` | `params`                |           | Get Cron params                                   |
| `aibd query cron` | `whitelisted-contracts` |           | Get the list of Whitelisted Contracts for cronjob |

### Transactions

| Command        | Subcommand             | Arguments                                  | Description                            |
| :------------- | :--------------------- | :----------------------------------------- | :------------------------------------- |
| `aibd tx cron` | `register-contract`    | [game-name] [contract-address] [game-type] | Register the contract for cron job     |
| `aibd tx cron` | `de-register-contract` | [game-id]                                  | De-Register the contract from cron job |
