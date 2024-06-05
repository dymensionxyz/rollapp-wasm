package keeper_test

import (
	"testing"

	"github.com/dymensionxyz/rollapp-wasm/app"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	cronTypes "github.com/dymensionxyz/rollapp-wasm/x/cron/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	appd := app.Setup(t, false)
	ctx := appd.BaseApp.NewContext(false, tmproto.Header{})

	defaultSecurityAddress := []string{"cosmos1xkxed7rdzvmyvgdshpe445ddqwn47fru24fnlp"}
	params := cronTypes.Params{SecurityAddress: defaultSecurityAddress, EnableCron: true}

	appd.CronKeeper.SetParams(ctx, params)

	require.EqualValues(t, params, appd.CronKeeper.GetParams(ctx))
}
