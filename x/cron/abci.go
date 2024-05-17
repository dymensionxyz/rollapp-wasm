package cron

import (
	"github.com/dymensionxyz/rollapp-wasm/x/cron/keeper"
	cronTypes "github.com/dymensionxyz/rollapp-wasm/x/cron/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	allContracts := k.GetAllContract(ctx)

	for _, data := range allContracts {
		if data.GameType == 1 {
			k.SinglePlayer(ctx, data.ContractAddress, cronTypes.ResolveSinglePlayer, data.GameName)
		} else if data.GameType == 2 {
			k.MultiPlayer(ctx, data.ContractAddress, cronTypes.SetupMultiPlayer, cronTypes.ResolveMultiPlayer, data.GameName)
		} else {
			k.SinglePlayer(ctx, data.ContractAddress, cronTypes.ResolveSinglePlayer, data.GameName)
			k.MultiPlayer(ctx, data.ContractAddress, cronTypes.SetupMultiPlayer, cronTypes.ResolveMultiPlayer, data.GameName)
		}
	}
}
