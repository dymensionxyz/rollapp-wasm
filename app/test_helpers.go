package app

import (
	"encoding/json"
	"testing"
	"time"

	"cosmossdk.io/math"
	"github.com/CosmWasm/wasmd/x/wasm"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/testutil/mock"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	sequencerstypes "github.com/dymensionxyz/dymension-rdk/x/sequencers/types"
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

var TestChainID = "rollappwasm_1234-1"
var (
	AccountPubKeyPrefix        = AccountAddressPrefix + sdk.PrefixPublic
	ValidatorAddressPrefix     = AccountAddressPrefix + sdk.PrefixValidator + sdk.PrefixOperator
	ValidatorPubKeyPrefix      = ValidatorAddressPrefix + sdk.PrefixPublic
	ConsensusNodeAddressPrefix = AccountAddressPrefix + sdk.PrefixValidator + sdk.PrefixConsensus
	ConsensusNodePubKeyPrefix  = ConsensusNodeAddressPrefix + sdk.PrefixPublic
)

// SetAccountAddressPrefixes sets the global prefix to be used when serializing addresses to bech32 strings.
func SetAccountAddressPrefixes() {
	config := sdk.GetConfig()

	config.SetBech32PrefixForAccount(AccountAddressPrefix, AccountPubKeyPrefix)
	config.SetBech32PrefixForValidator(ValidatorAddressPrefix, ValidatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(ConsensusNodeAddressPrefix, ConsensusNodePubKeyPrefix)
}

func setup(_ *testing.T, withGenesis bool) (*App, GenesisState) {
	db := dbm.NewMemDB()
	var emptyWasmOpts []wasm.Option
	app := NewRollapp(log.NewNopLogger(), db, nil, true, map[int64]bool{}, DefaultNodeHome, 5, MakeEncodingConfig(), GetEnabledProposals(), EmptyAppOptions{}, emptyWasmOpts)
	if withGenesis {
		return app, NewDefaultGenesisState(app.AppCodec())
	}
	return app, GenesisState{}
}

// Setup initializes a new App. A Nop logger is set in App.
func Setup(t *testing.T, isCheckTx bool) *App {
	t.Helper()

	SetAccountAddressPrefixes()

	privVal := mock.NewPV()
	pubKey, err := privVal.GetPubKey()
	require.NoError(t, err)

	// create validator set with single validator
	validator := tmtypes.NewValidator(pubKey, 1)
	valSet := tmtypes.NewValidatorSet([]*tmtypes.Validator{validator})

	// generate genesis account
	senderPrivKey := secp256k1.GenPrivKey()
	acc := authtypes.NewBaseAccount(senderPrivKey.PubKey().Address().Bytes(), senderPrivKey.PubKey(), 0, 0)
	balance := banktypes.Balance{
		Address: acc.GetAddress().String(),
		Coins:   sdk.NewCoins(sdk.NewCoin("aaib", sdk.NewInt(100000000000000))),
	}

	// default genesis operator address for dymint sequencers
	genesisOperatorAddr, _ := sdk.ValAddressFromHex(validator.Address.String())

	app := SetupWithGenesisValSet(t, valSet, []authtypes.GenesisAccount{acc}, genesisOperatorAddr, balance)

	return app
}

func SetupWithGenesisValSet(t *testing.T, valSet *tmtypes.ValidatorSet, genAccs []authtypes.GenesisAccount, genesisOperatorAddr sdk.ValAddress, balances ...banktypes.Balance) *App {
	t.Helper()

	app, genesisState := setup(t, true)
	genesisState = genesisStateWithValSet(t, app, genesisState, valSet, genAccs, genesisOperatorAddr, balances...)

	stateBytes, err := json.MarshalIndent(genesisState, "", " ")
	require.NoError(t, err)

	// init chain will set the validator set and initialize the genesis accounts
	app.InitChain(
		abci.RequestInitChain{
			ChainId: TestChainID,
			Validators: []abci.ValidatorUpdate{
				abci.Ed25519ValidatorUpdate(valSet.Validators[0].PubKey.Bytes(), 1000),
			},
			ConsensusParams: DefaultConsensusParams,
			AppStateBytes:   stateBytes,
		},
	)

	return app
}

func genesisStateWithValSet(t *testing.T,
	app *App, genesisState GenesisState,
	valSet *tmtypes.ValidatorSet, genAccs []authtypes.GenesisAccount,
	genesisOperatorAddr sdk.ValAddress,
	balances ...banktypes.Balance,
) GenesisState {
	codec := app.AppCodec()

	// set genesis accounts
	authGenesis := authtypes.NewGenesisState(authtypes.DefaultParams(), genAccs)
	genesisState[authtypes.ModuleName] = codec.MustMarshalJSON(authGenesis)

	validators := make([]stakingtypes.Validator, 0, len(valSet.Validators))
	delegations := make([]stakingtypes.Delegation, 0, len(valSet.Validators))

	bondAmt := sdk.DefaultPowerReduction

	for _, val := range valSet.Validators {
		pk, err := cryptocodec.FromTmPubKeyInterface(val.PubKey)
		require.NoError(t, err)
		pkAny, err := codectypes.NewAnyWithValue(pk)
		require.NoError(t, err)
		validator := stakingtypes.Validator{
			OperatorAddress:   sdk.ValAddress(val.Address).String(),
			ConsensusPubkey:   pkAny,
			Jailed:            false,
			Status:            stakingtypes.Bonded,
			Tokens:            bondAmt,
			DelegatorShares:   sdk.OneDec(),
			Description:       stakingtypes.Description{},
			UnbondingHeight:   int64(0),
			UnbondingTime:     time.Unix(0, 0).UTC(),
			Commission:        stakingtypes.NewCommission(sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec()),
			MinSelfDelegation: math.ZeroInt(),
		}
		validators = append(validators, validator)
		delegations = append(delegations, stakingtypes.NewDelegation(genAccs[0].GetAddress(), val.Address.Bytes(), sdk.OneDec()))

	}

	defaultStParams := stakingtypes.DefaultParams()
	stParams := stakingtypes.NewParams(
		defaultStParams.UnbondingTime,
		defaultStParams.MaxValidators,
		defaultStParams.MaxEntries,
		defaultStParams.HistoricalEntries,
		"aaib",
		defaultStParams.MinCommissionRate, // 5%
	)

	// set validators and delegations
	stakingGenesis := stakingtypes.NewGenesisState(stParams, validators, delegations)
	genesisState[stakingtypes.ModuleName] = codec.MustMarshalJSON(stakingGenesis)

	// add bonded amount to bonded pool module account
	balances = append(balances, banktypes.Balance{
		Address: authtypes.NewModuleAddress(stakingtypes.BondedPoolName).String(),
		Coins:   sdk.Coins{sdk.NewCoin("aaib", bondAmt.MulRaw(int64(len(valSet.Validators))))},
	})

	totalSupply := sdk.NewCoins()
	for _, b := range balances {
		// add genesis acc tokens to total supply
		totalSupply = totalSupply.Add(b.Coins...)
	}

	// update genesisOperatorAddr in sequencers module
	sequencersGenesis := sequencerstypes.DefaultGenesis()
	sequencersGenesis.GenesisOperatorAddress = genesisOperatorAddr.String()
	genesisState[sequencerstypes.ModuleName] = codec.MustMarshalJSON(sequencersGenesis)

	// update total supply
	bankGenesis := banktypes.NewGenesisState(banktypes.DefaultGenesisState().Params, balances, totalSupply, []banktypes.Metadata{})
	genesisState[banktypes.ModuleName] = codec.MustMarshalJSON(bankGenesis)
	// println("genesisStateWithValSet bankState:", string(genesisState[banktypes.ModuleName]))

	return genesisState
}

// EmptyAppOptions is a stub implementing AppOptions
type EmptyAppOptions struct{}

// Get implements AppOptions
func (ao EmptyAppOptions) Get(o string) interface{} {
	return nil
}
