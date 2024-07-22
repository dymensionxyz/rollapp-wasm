package types

import (
	wasmdtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	cwerrortypes "github.com/dymensionxyz/rollapp-wasm/x/cwerrors/types"
)

type WasmKeeperExpected interface {
	HasContractInfo(ctx sdk.Context, contractAddress sdk.AccAddress) bool
	GetContractInfo(ctx sdk.Context, contractAddress sdk.AccAddress) *wasmdtypes.ContractInfo
	Sudo(ctx sdk.Context, contractAddress sdk.AccAddress, msg []byte) ([]byte, error)
}

type BankKeeperExpected interface {
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error
	BlockedAddr(addr sdk.AccAddress) bool
}

type ErrorsKeeperExpected interface {
	SetError(ctx sdk.Context, sudoErr cwerrortypes.SudoError) error
}
