package drs3

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	rollappparamskeeper "github.com/dymensionxyz/dymension-rdk/x/rollappparams/keeper"

	drs2 "github.com/dymensionxyz/rollapp-wasm/app/upgrades/drs-2"
)

func CreateUpgradeHandler(
	rpKeeper rollappparamskeeper.Keeper,
	mm *module.Manager,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		if rpKeeper.Version(ctx) == 1 {
			// first run drs-2 migration
			if err := drs2.HandleUpgrade(ctx, rpKeeper); err != nil {
				return nil, err
			}
		}
		// upgrade drs to 3
		if err := rpKeeper.SetVersion(ctx, uint32(3)); err != nil {
			return nil, err
		}
		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}
