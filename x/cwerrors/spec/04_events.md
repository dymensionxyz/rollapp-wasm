# Events

Section describes the module events

The module emits the following proto-events

| Source type | Source name           | Protobuf  reference                                                                  |
| ----------- | --------------------- |--------------------------------------------------------------------------------------|
| Message     | `MsgUpdateParams`     | [ParamsUpdatedEvent](../../../proto/rollapp/cwerrors/v1/events.proto#L12)            |
| Message     | `MsgSubscribeToError` | [SubscribedToErrorsEvent](../../../proto/rollapp/cwerrors/v1/events.proto#L20)       |
| Keeper      | `SetErrorInState`     | [StoringErrorEvent](../../../proto/rollapp/cwerrors/v1/events.proto#L32)             |
| Module      | `EndBlocker`          | [SudoErrorCallbackFailedEvent](../../../proto/rollapp/cwerrors/v1/events.proto#L40)  |
