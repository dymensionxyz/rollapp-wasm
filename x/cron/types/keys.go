package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "cron"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_cron"
)

var (
	SetContractKeyPrefix = []byte{0x11}
	GameIDKey            = []byte{0x12}
)

func ContractKey(gameID uint64) []byte {
	return append(SetContractKeyPrefix, sdk.Uint64ToBigEndian(gameID)...)
}
