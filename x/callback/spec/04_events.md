# Events

Section describes the module events

The module emits the following proto-events

| Source type | Source name          | Protobuf  reference                                                                  |
| ----------- | -------------------- |--------------------------------------------------------------------------------------|
| Message     | `MsgRequestCallback` | [CallbackRegisteredEvent](../../../proto/rollapp/callback/v1/events.proto#L11)       |
| Message     | `MsgCancelCallback`  | [CallbackCancelledEvent](../../../proto/rollapp/callback/v1/events.proto#L25)        |
| Module      | `EndBlocker`         | [CallbackExecutedSuccessEvent](../../../proto/rollapp/callback/v1/events.proto#L39)  |
| Module      | `EndBlocker`         | [CallbackExecutedFailedEvent](../../../proto/rollapp/callback/v1/events.proto#L53)   |
