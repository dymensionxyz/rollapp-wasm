package drs2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	rollappparamskeeper "github.com/dymensionxyz/dymension-rdk/x/rollappparams/keeper"
	rollappparamstypes "github.com/dymensionxyz/dymension-rdk/x/rollappparams/types"
)

func CreateUpgradeHandler(
	rpKeeper rollappparamskeeper.Keeper,
	mm *module.Manager,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		//migrate rollapp params with missing min-gas-prices and updating drs to 2
		err := rpKeeper.SetVersion(ctx, uint32(2))
		if err != nil {
			return nil, err
		}
		err = rpKeeper.SetMinGasPrices(ctx, rollappparamstypes.DefaultParams().MinGasPrices)
		if err != nil {
			return nil, err
		}

		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}
