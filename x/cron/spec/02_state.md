<!--
order: 2
-->

# State

## State Objects

The `x/cron` module keeps the following object in the state: WhitelistedContract.

This object is used to store the state of a

- `WhitelistedContract` - to store the contract which is whitelisted according to game type.

```go
// this defines the details of the contract registered for cronjob
message WhitelistedContract {
    // id of the game(counter)
    uint64 game_id = 1;
    // address which was used to register the contract
    string security_address = 2;
    // admin of the contract
    string contract_admin = 3;
    // name of the game contract
    string game_name = 4;
    // CosmWasm contract address
    string contract_address = 5;
    // single player -> 1, multi player -> 2, both single and multiplayer -> 3
    uint64 game_type = 6;
}
```

## Genesis & Params

The `x/cron` module's `GenesisState` defines the state necessary for initializing the chain from a previously exported height. It contains the module Parameters and Whitelisted Contract. The params are used to control the Contract GasLimit and Security Address which is responsible to whitelist contract. This value can be modified with a governance proposal.

```go
// GenesisState defines the cron module's genesis state.
message GenesisState {
  Params params = 1 [
    (gogoproto.moretags) = "yaml:\"params\"",
    (gogoproto.nullable) = false
  ];
  repeated WhitelistedContract whitelisted_contracts  = 2  [
    (gogoproto.moretags) = "yaml:\"whitelisted_contracts\"",
    (gogoproto.nullable) = false
  ];
}
```

```go
// Params defines the parameters for the module.
message Params {
  // Security address that can whitelist/delist contract
  repeated string security_address = 1 [
    (gogoproto.jsontag) = "security_address,omitempty",
    (gogoproto.moretags) = "yaml:\"security_address\""
  ];

  uint64 contract_gas_limit = 2 [
    (gogoproto.jsontag) = "contract_gas_limit,omitempty",
    (gogoproto.moretags) = "yaml:\"contract_gas_limit\""
  ];
}
```

## State Transitions

The following state transitions are possible:

- Register the contract for cronjob
- DeRegister the contract from cronjob
