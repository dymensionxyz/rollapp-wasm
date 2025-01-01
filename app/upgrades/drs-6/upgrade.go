package drs6

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	rollappparamskeeper "github.com/dymensionxyz/dymension-rdk/x/rollappparams/keeper"

	drs4 "github.com/dymensionxyz/rollapp-wasm/app/upgrades/drs-4"
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
	if rpKeeper.Version(ctx) < 5 {
		// first run drs-5 migration
		if err := drs4.HandleUpgrade(ctx, rpKeeper); err != nil {
			return err
		}
	}
	// upgrade drs to 6
	if err := rpKeeper.SetVersion(ctx, uint32(6)); err != nil {
		return err
	}
	return nil
}
