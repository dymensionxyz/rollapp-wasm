package cron_test

import (
	"testing"

	"github.com/dymensionxyz/rollapp-wasm/app"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/dymensionxyz/rollapp-wasm/x/cron"
	"github.com/dymensionxyz/rollapp-wasm/x/cron/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	appd := app.Setup(t, false)
	ctx := appd.BaseApp.NewContext(false, tmproto.Header{})

	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	cron.InitGenesis(ctx, appd.CronKeeper, genesisState)
	got := cron.ExportGenesis(ctx, appd.CronKeeper)
	require.NotNil(t, got)
}
