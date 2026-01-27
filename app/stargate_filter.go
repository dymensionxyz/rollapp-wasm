package app

import (
	"strings"

	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
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

// NewIBCFilteredStargateEncoder returns a stargate encoder that blocks IBC messages.
// It wraps the default EncodeStargateMsg and rejects any message whose TypeURL
// matches a blocked IBC prefix, preventing contracts from bypassing ante handler
// restrictions (e.g. relayer whitelisting, connection open restrictions).
func NewIBCFilteredStargateEncoder(unpacker codectypes.AnyUnpacker) wasmkeeper.StargateEncoder {
	defaultEncoder := wasmkeeper.EncodeStargateMsg(unpacker)

	return func(sender sdk.AccAddress, msg *wasmvmtypes.StargateMsg) ([]sdk.Msg, error) {
		if isBlockedStargateMsg(msg.TypeURL) {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
				"ibc message %s is not allowed from cosmwasm contracts", msg.TypeURL)
		}
		return defaultEncoder(sender, msg)
	}
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
