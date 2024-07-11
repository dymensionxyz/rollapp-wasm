# Messages

Section describes the processing of the module messages

## MsgSubscribeToError

A contract can be subscribed to errors by using the [MsgSubscribeToError](../../../proto/archway/cwerrors/v1/tx.proto) message.

```protobuf
message MsgSubscribeToError {
    option (cosmos.msg.v1.signer) = "sender";
    // sender is the address of who is registering the contracts for callback on error
    string sender = 1;
    // contract_address is the contarct subscribing to the error
    string contract_address = 2;
    // fee is the subscription fee for the feature
    cosmos.base.v1beta1.Coin fee = 3 [ (gogoproto.nullable) = false ];
}
```

On success
* A subscription is created valid for the duration as specified in the module params.
* The subscription fees are sent to the fee collector
* In case a subscription already exists, it is extended.

This message is expected to fail if:
* The sender address and contract address are not valid addresses
* There is no contract with given address
* The sender is not authorized to subscribe - the sender is not the contract owner/admin or the contract itself
* The user does not send enough funds or doesnt have enough funds