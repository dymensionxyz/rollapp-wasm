package drs7

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	rollappparamskeeper "github.com/dymensionxyz/dymension-rdk/x/rollappparams/keeper"

	drs6 "github.com/dymensionxyz/rollapp-wasm/app/upgrades/drs-6"
)

func CreateUpgradeHandler(
	rpKeeper rollappparamskeeper.Keeper,
	mm *module.Manager,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		if err := HandleUpgrade(ctx, rpKeeper); err != nil {
			return nil, err
		}
		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}

func HandleUpgrade(ctx sdk.Context, rpKeeper rollappparamskeeper.Keeper) error {
	if rpKeeper.Version(ctx) < 6 {
		// first run drs-6 migration
		if err := drs6.HandleUpgrade(ctx, rpKeeper); err != nil {
			return err
		}
	}
	// upgrade drs to 7
	if err := rpKeeper.SetVersion(ctx, uint32(7)); err != nil {
		return err
	}
	return nil
}
