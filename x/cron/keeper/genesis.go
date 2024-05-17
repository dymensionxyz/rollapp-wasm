package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/dymensionxyz/rollapp-wasm/x/cron/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	var (
		GameID uint64
	)
	// this line is used by starport scaffolding # genesis/module/init
	for _, item := range genState.WhitelistedContracts {
		if item.GameId > GameID {
			GameID = item.GameId
		}
		k.SetContract(ctx, item)
	}
	k.SetGameID(ctx, GameID)
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {

	return &types.GenesisState{
		Params:               k.GetParams(ctx),
		WhitelistedContracts: k.GetAllContract(ctx),
	}
}
