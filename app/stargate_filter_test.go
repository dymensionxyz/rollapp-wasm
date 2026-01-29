package app

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsBlockedStargateMsg(t *testing.T) {
	tests := []struct {
		name    string
		typeURL string
		blocked bool
	}{
		// IBC client messages - should be blocked
		{
			name:    "MsgCreateClient",
			typeURL: "/ibc.core.client.v1.MsgCreateClient",
			blocked: true,
		},
		{
			name:    "MsgUpdateClient",
			typeURL: "/ibc.core.client.v1.MsgUpdateClient",
			blocked: true,
		},
		// IBC connection messages - should be blocked
		{
			name:    "MsgConnectionOpenInit",
			typeURL: "/ibc.core.connection.v1.MsgConnectionOpenInit",
			blocked: true,
		},
		{
			name:    "MsgConnectionOpenTry",
			typeURL: "/ibc.core.connection.v1.MsgConnectionOpenTry",
			blocked: true,
		},
		{
			name:    "MsgConnectionOpenAck",
			typeURL: "/ibc.core.connection.v1.MsgConnectionOpenAck",
			blocked: true,
		},
		{
			name:    "MsgConnectionOpenConfirm",
			typeURL: "/ibc.core.connection.v1.MsgConnectionOpenConfirm",
			blocked: true,
		},
		// IBC channel messages - should be blocked
		{
			name:    "MsgChannelOpenInit",
			typeURL: "/ibc.core.channel.v1.MsgChannelOpenInit",
			blocked: true,
		},
		{
			name:    "MsgChannelOpenTry",
			typeURL: "/ibc.core.channel.v1.MsgChannelOpenTry",
			blocked: true,
		},
		{
			name:    "MsgRecvPacket",
			typeURL: "/ibc.core.channel.v1.MsgRecvPacket",
			blocked: true,
		},
		{
			name:    "MsgAcknowledgement",
			typeURL: "/ibc.core.channel.v1.MsgAcknowledgement",
			blocked: true,
		},
		{
			name:    "MsgTimeout",
			typeURL: "/ibc.core.channel.v1.MsgTimeout",
			blocked: true,
		},
		// IBC transfer messages - should be blocked
		{
			name:    "MsgTransfer",
			typeURL: "/ibc.applications.transfer.v1.MsgTransfer",
			blocked: true,
		},
		// Non-IBC messages - should be allowed
		{
			name:    "BankMsgSend",
			typeURL: "/cosmos.bank.v1beta1.MsgSend",
			blocked: false,
		},
		{
			name:    "StakingDelegate",
			typeURL: "/cosmos.staking.v1beta1.MsgDelegate",
			blocked: false,
		},
		{
			name:    "WasmExecute",
			typeURL: "/cosmwasm.wasm.v1.MsgExecuteContract",
			blocked: false,
		},
		{
			name:    "GovSubmitProposal",
			typeURL: "/cosmos.gov.v1beta1.MsgSubmitProposal",
			blocked: false,
		},
		{
			name:    "empty string",
			typeURL: "",
			blocked: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := isBlockedStargateMsg(tc.typeURL)
			require.Equal(t, tc.blocked, result,
				"typeURL %s: expected blocked=%v, got blocked=%v", tc.typeURL, tc.blocked, result)
		})
	}
}
