package app

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/CosmWasm/wasmd/x/wasm"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sequencerstypes "github.com/dymensionxyz/dymension-rdk/x/sequencers/types"
	cronTypes "github.com/dymensionxyz/rollapp-wasm/x/cron/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"
)

// DefaultConsensusParams defines the default Tendermint consensus params used in
// App testing.
var DefaultConsensusParams = &abci.ConsensusParams{
	Block: &abci.BlockParams{
		MaxBytes: 200000,
		MaxGas:   -1,
	},
	Evidence: &tmproto.EvidenceParams{
		MaxAgeNumBlocks: 302400,
		MaxAgeDuration:  504 * time.Hour, // 3 weeks is the max duration
		MaxBytes:        10000,
	},
	Validator: &tmproto.ValidatorParams{
		PubKeyTypes: []string{
			tmtypes.ABCIPubKeyTypeEd25519,
		},
	},
}

var (
	ProposerPK       = simapp.CreateTestPubKeys(1)[0]
	ProposerConsAddr = sdk.ConsAddress(ProposerPK.Address())

	OperatorPK = secp256k1.GenPrivKey().PubKey()
)

var TestChainID = "rollappwasm_1234-1"

func setup(withGenesis bool, invCheckPeriod uint) (*App, map[string]json.RawMessage) {
	db := dbm.NewMemDB()

	encCdc := MakeEncodingConfig()
	var emptyWasmOpts []wasm.Option
	testApp := NewRollapp(
		log.NewNopLogger(), db, nil, true, map[int64]bool{}, "/tmp", invCheckPeriod, encCdc, GetEnabledProposals(), EmptyAppOptions{}, emptyWasmOpts,
	)
	if withGenesis {
		return testApp, NewDefaultGenesisState(encCdc.Codec)
	}
	return testApp, map[string]json.RawMessage{}
}

// Setup initializes a new Rollapp. A Nop logger is set in Rollapp.
func Setup(t *testing.T, isCheckTx bool) *App {
	t.Helper()

	pk, err := cryptocodec.ToTmProtoPublicKey(ProposerPK)
	require.NoError(t, err)

	app, genesisState := setup(true, 5)

	// setup for sequencer
	seqGenesis := sequencerstypes.GenesisState{
		Params:                 sequencerstypes.DefaultParams(),
		GenesisOperatorAddress: sdk.ValAddress(OperatorPK.Address()).String(),
	}
	genesisState[sequencerstypes.ModuleName] = app.AppCodec().MustMarshalJSON(&seqGenesis)

	// setup for cron
	DefaultSecurityAddress := []string{"cosmos1xkxed7rdzvmyvgdshpe445ddqwn47fru24fnlp"}
	params := cronTypes.Params{SecurityAddress: DefaultSecurityAddress}
	cronGenesis := cronTypes.GenesisState{
		Params:               params,
		WhitelistedContracts: []cronTypes.WhitelistedContract{},
	}
	genesisState[cronTypes.ModuleName] = app.AppCodec().MustMarshalJSON(&cronGenesis)

	// for now bank genesis won't be set here, funding accounts should be called with fund utils.FundModuleAccount

	stateBytes, err := json.MarshalIndent(genesisState, "", " ")
	require.NoError(t, err)
	// init chain will set the validator set and initialize the genesis accounts
	app.InitChain(
		abci.RequestInitChain{
			Time:            time.Time{},
			ChainId:         TestChainID,
			ConsensusParams: DefaultConsensusParams,
			Validators: []abci.ValidatorUpdate{
				{PubKey: pk, Power: 1},
			},
			AppStateBytes: stateBytes,
			InitialHeight: 0,
		},
	)

	return app
}

// EmptyAppOptions is a stub implementing AppOptions
type EmptyAppOptions struct{}

// Get implements AppOptions
func (ao EmptyAppOptions) Get(o string) interface{} {
	return nil
}
