package app

import (
	"strings"

	errorsmod "cosmossdk.io/errors"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// blockedStargateMsgPrefixes defines IBC message type URL prefixes that contracts
// must not be able to dispatch via stargate. This prevents CosmWasm contracts from
// bypassing the ante handler's IBC message filtering.
var blockedStargateMsgPrefixes = []string{
	"/ibc.core.client.",
	"/ibc.core.connection.",
	"/ibc.core.channel.",
	"/ibc.applications.transfer.",
}

// IBCFilteringMessageHandlerDecorator returns a decorator that filters IBC messages
// from CosmWasm contracts before they reach the underlying message handler.
func IBCFilteringMessageHandlerDecorator() func(old wasmkeeper.Messenger) wasmkeeper.Messenger {
	return func(old wasmkeeper.Messenger) wasmkeeper.Messenger {
		return &ibcFilteringMessenger{wrapped: old}
	}
}

type ibcFilteringMessenger struct {
	wrapped wasmkeeper.Messenger
}

// DispatchMsg implements the Messenger interface
func (m *ibcFilteringMessenger) DispatchMsg(
	ctx sdk.Context,
	contractAddr sdk.AccAddress,
	contractIBCPortID string,
	msg wasmvmtypes.CosmosMsg,
) (events []sdk.Event, data [][]byte, err error) {
	// Check if this is a stargate message with a blocked IBC type URL
	if msg.Stargate != nil && isBlockedStargateMsg(msg.Stargate.TypeURL) {
		return nil, nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized,
			"ibc message %s is not allowed from cosmwasm contracts", msg.Stargate.TypeURL)
	}
	return m.wrapped.DispatchMsg(ctx, contractAddr, contractIBCPortID, msg)
}

// isBlockedStargateMsg checks if a message type URL matches any blocked IBC prefix.
func isBlockedStargateMsg(typeURL string) bool {
	for _, prefix := range blockedStargateMsgPrefixes {
		if strings.HasPrefix(typeURL, prefix) {
			return true
		}
	}
	return false
}
