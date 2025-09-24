package drs10

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	rollappparamskeeper "github.com/dymensionxyz/dymension-rdk/x/rollappparams/keeper"

	"github.com/dymensionxyz/rollapp-wasm/app/upgrades"
	drs9 "github.com/dymensionxyz/rollapp-wasm/app/upgrades/drs-9"
)

func CreateUpgradeHandler(
	kk upgrades.UpgradeKeepers,
	mm *module.Manager,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		if err := HandleUpgrade(ctx, kk.RpKeeper); err != nil {
			return nil, err
		}
		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}

func HandleUpgrade(ctx sdk.Context, rpKeeper rollappparamskeeper.Keeper) error {
	if rpKeeper.Version(ctx) < 9 {
		// first run drs-9 migration
		if err := drs9.HandleUpgrade(ctx, rpKeeper); err != nil {
			return err
		}
	}
	// upgrade drs to 10
	if err := rpKeeper.SetVersion(ctx, uint32(10)); err != nil {
		return err
	}
	return nil
}
