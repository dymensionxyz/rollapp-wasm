package app

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/cosmos/cosmos-sdk/x/authz"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	gaslessmodule "github.com/dymensionxyz/dymension-rdk/x/gasless"
	gaslesskeeper "github.com/dymensionxyz/dymension-rdk/x/gasless/keeper"
	gaslesstypes "github.com/dymensionxyz/dymension-rdk/x/gasless/types"
	"github.com/gorilla/mux"
	"github.com/rakyll/statik/fs"
	abci "github.com/tendermint/tendermint/abci/types"
	tmjson "github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	nodeservice "github.com/cosmos/cosmos-sdk/client/grpc/node"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store/streaming"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth/posthandler"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"

	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"

	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	"github.com/dymensionxyz/dymension-rdk/x/mint"
	mintkeeper "github.com/dymensionxyz/dymension-rdk/x/mint/keeper"
	minttypes "github.com/dymensionxyz/dymension-rdk/x/mint/types"

	"github.com/dymensionxyz/dymension-rdk/x/epochs"
	epochskeeper "github.com/dymensionxyz/dymension-rdk/x/epochs/keeper"
	epochstypes "github.com/dymensionxyz/dymension-rdk/x/epochs/types"

	hubgenesis "github.com/dymensionxyz/dymension-rdk/x/hub-genesis"
	hubgenkeeper "github.com/dymensionxyz/dymension-rdk/x/hub-genesis/keeper"
	hubgentypes "github.com/dymensionxyz/dymension-rdk/x/hub-genesis/types"

	ibctransfer "github.com/cosmos/ibc-go/v6/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v6/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v6/modules/core"
	ibcclient "github.com/cosmos/ibc-go/v6/modules/core/02-client"
	ibcclientclient "github.com/cosmos/ibc-go/v6/modules/core/02-client/client"
	ibcclienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	ibcporttypes "github.com/cosmos/ibc-go/v6/modules/core/05-port/types"
	ibchost "github.com/cosmos/ibc-go/v6/modules/core/24-host"
	ibckeeper "github.com/cosmos/ibc-go/v6/modules/core/keeper"
	ibctestingtypes "github.com/cosmos/ibc-go/v6/testing/types"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmclient "github.com/CosmWasm/wasmd/x/wasm/client"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"

	rollappwasmparams "github.com/dymensionxyz/rollapp-wasm/app/params"

	// unnamed import of statik for swagger UI support
	_ "github.com/cosmos/cosmos-sdk/client/docs/statik"

	"github.com/dymensionxyz/dymension-rdk/x/staking"
	stakingkeeper "github.com/dymensionxyz/dymension-rdk/x/staking/keeper"

	"github.com/dymensionxyz/dymension-rdk/x/sequencers"
	seqkeeper "github.com/dymensionxyz/dymension-rdk/x/sequencers/keeper"
	seqtypes "github.com/dymensionxyz/dymension-rdk/x/sequencers/types"

	distr "github.com/dymensionxyz/dymension-rdk/x/dist"
	distrkeeper "github.com/dymensionxyz/dymension-rdk/x/dist/keeper"

	"github.com/dymensionxyz/dymension-rdk/x/denommetadata"
	denommetadatamoduletypes "github.com/dymensionxyz/dymension-rdk/x/denommetadata/types"
	"github.com/dymensionxyz/dymension-rdk/x/hub"
	hubkeeper "github.com/dymensionxyz/dymension-rdk/x/hub/keeper"
	hubtypes "github.com/dymensionxyz/dymension-rdk/x/hub/types"

	"github.com/dymensionxyz/rollapp-wasm/x/callback"
	callbackKeeper "github.com/dymensionxyz/rollapp-wasm/x/callback/keeper"
	callbackTypes "github.com/dymensionxyz/rollapp-wasm/x/callback/types"

	"github.com/dymensionxyz/rollapp-wasm/x/cwerrors"
	cwerrorsKeeper "github.com/dymensionxyz/rollapp-wasm/x/cwerrors/keeper"
	cwerrorsTypes "github.com/dymensionxyz/rollapp-wasm/x/cwerrors/types"

	rollappparams "github.com/dymensionxyz/dymension-rdk/x/rollappparams"
	rollappparamskeeper "github.com/dymensionxyz/dymension-rdk/x/rollappparams/keeper"
	rollappparamstypes "github.com/dymensionxyz/dymension-rdk/x/rollappparams/types"
)

const (
	Name = "rollapp-wasm"
)

var (
	AccountAddressPrefix string
	kvstorekeys          = []string{
		authtypes.StoreKey, authzkeeper.StoreKey,
		feegrant.StoreKey, banktypes.StoreKey,
		stakingtypes.StoreKey, seqtypes.StoreKey,
		minttypes.StoreKey, distrtypes.StoreKey,
		govtypes.StoreKey, paramstypes.StoreKey,
		ibchost.StoreKey, upgradetypes.StoreKey,
		epochstypes.StoreKey, hubgentypes.StoreKey,
		ibctransfertypes.StoreKey, capabilitytypes.StoreKey,
		wasmtypes.StoreKey, gaslesstypes.StoreKey,
		hubtypes.StoreKey,
		callbackTypes.StoreKey,
		cwerrorsTypes.StoreKey,
		rollappparamstypes.StoreKey,
	}
)

func getGovProposalHandlers() []govclient.ProposalHandler {
	govProposalHandlers := append(
		wasmclient.ProposalHandlers,
		paramsclient.ProposalHandler,
		distrclient.ProposalHandler,
		upgradeclient.LegacyProposalHandler,
		upgradeclient.LegacyCancelProposalHandler,
		ibcclientclient.UpdateClientProposalHandler,
		ibcclientclient.UpgradeProposalHandler,
	)

	return govProposalHandlers
}

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		authzmodule.AppModuleBasic{},
		genutil.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		sequencers.AppModuleBasic{},
		mint.AppModuleBasic{},
		epochs.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(getGovProposalHandlers()),
		params.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		gaslessmodule.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		ibc.AppModuleBasic{},
		ibctransfer.AppModuleBasic{},
		vesting.AppModuleBasic{},
		hubgenesis.AppModuleBasic{},
		wasm.AppModuleBasic{},
		hub.AppModuleBasic{},
		callback.AppModuleBasic{},
		cwerrors.AppModuleBasic{},
		rollappparams.AppModuleBasic{},
	)

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:     nil,
		authz.ModuleName:               nil,
		distrtypes.ModuleName:          nil,
		minttypes.ModuleName:           {authtypes.Minter},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:            {authtypes.Burner},
		ibctransfertypes.ModuleName:    {authtypes.Minter, authtypes.Burner},
		wasmtypes.ModuleName:           {authtypes.Burner},
		hubgentypes.ModuleName:         {authtypes.Minter},
		gaslesstypes.ModuleName:        nil,
		callbackTypes.ModuleName:       nil,
		rollappparamstypes.ModuleName:  nil,
	}

	// module accounts that are allowed to receive tokens
	maccCanReceiveTokens = []string{
		distrtypes.ModuleName,
		hubgentypes.ModuleName,
	}
)

var (
	_ servertypes.Application = (*App)(nil)
	_ simapp.App              = (*App)(nil)
)

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, "."+Name)
}

// App extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type App struct {
	*baseapp.BaseApp

	cdc               *codec.LegacyAmino
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry

	invCheckPeriod uint

	// keys to access the substores
	keys    map[string]*storetypes.KVStoreKey
	tkeys   map[string]*storetypes.TransientStoreKey
	memKeys map[string]*storetypes.MemoryStoreKey

	// keepers
	AccountKeeper    authkeeper.AccountKeeper
	AuthzKeeper      authzkeeper.Keeper
	BankKeeper       bankkeeper.Keeper
	CapabilityKeeper *capabilitykeeper.Keeper
	StakingKeeper    stakingkeeper.Keeper
	SequencersKeeper seqkeeper.Keeper
	MintKeeper       mintkeeper.Keeper
	EpochsKeeper     epochskeeper.Keeper
	DistrKeeper      distrkeeper.Keeper
	GovKeeper        govkeeper.Keeper
	HubGenesisKeeper hubgenkeeper.Keeper
	UpgradeKeeper    upgradekeeper.Keeper
	ParamsKeeper     paramskeeper.Keeper
	IBCKeeper        *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	TransferKeeper   ibctransferkeeper.Keeper
	WasmKeeper       wasmkeeper.Keeper
	FeeGrantKeeper   feegrantkeeper.Keeper
	GaslessKeeper    gaslesskeeper.Keeper
	CallbackKeeper   callbackKeeper.Keeper
	CWErrorsKeeper   cwerrorsKeeper.Keeper

	// make scoped keepers public for test purposes
	ScopedIBCKeeper      capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper capabilitykeeper.ScopedKeeper
	ScopedWasmKeeper     capabilitykeeper.ScopedKeeper

	HubKeeper hubkeeper.Keeper

	RollappConsensusParamsKeeper rollappparamskeeper.Keeper

	// mm is the module manager
	mm *module.Manager

	// sm is the simulation manager
	sm *module.SimulationManager

	// module configurator
	configurator module.Configurator
}

// NewRollapp returns a reference to an initialized blockchain app
func NewRollapp(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	skipUpgradeHeights map[int64]bool,
	homePath string,
	invCheckPeriod uint,
	encodingConfig rollappwasmparams.EncodingConfig,
	enabledProposals []wasm.ProposalType,
	appOpts servertypes.AppOptions,
	wasmOpts []wasmkeeper.Option,
	baseAppOptions ...func(*baseapp.BaseApp),
) *App {
	appCodec := encodingConfig.Codec
	cdc := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	bApp := baseapp.NewBaseApp(
		Name,
		logger,
		db,
		encodingConfig.TxConfig.TxDecoder(),
		baseAppOptions...,
	)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)

	keys := sdk.NewKVStoreKeys(
		kvstorekeys...,
	)
	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey, cwerrorsTypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	// load state streaming if enabled
	if _, _, err := streaming.LoadStreamingServices(bApp, appOpts, appCodec, keys); err != nil {
		panic("failed to load state streaming services: " + err.Error())
	}

	app := &App{
		BaseApp:           bApp,
		cdc:               cdc,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		invCheckPeriod:    invCheckPeriod,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
	}

	app.ParamsKeeper = initParamsKeeper(
		appCodec,
		cdc,
		keys[paramstypes.StoreKey],
		tkeys[paramstypes.TStoreKey],
	)

	// set the BaseApp's parameter store
	bApp.SetParamStore(app.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramstypes.ConsensusParamsKeyTable()))

	app.CapabilityKeeper = capabilitykeeper.NewKeeper(appCodec, keys[capabilitytypes.StoreKey], memKeys[capabilitytypes.MemStoreKey])
	// Applications that wish to enforce statically created ScopedKeepers should call `Seal` after creating
	// their scoped modules in `NewApp` with `ScopeToModule`
	scopedIBCKeeper := app.CapabilityKeeper.ScopeToModule(ibchost.ModuleName)
	scopedTransferKeeper := app.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	scopedWasmKeeper := app.CapabilityKeeper.ScopeToModule(wasmtypes.ModuleName)
	app.CapabilityKeeper.Seal()

	// add keepers
	app.EpochsKeeper = *epochskeeper.NewKeeper(appCodec, keys[epochstypes.StoreKey])

	app.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec,
		keys[authtypes.StoreKey],
		app.GetSubspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount,
		maccPerms,
		sdk.GetConfig().GetBech32AccountAddrPrefix(),
	)

	app.AuthzKeeper = authzkeeper.NewKeeper(
		keys[authzkeeper.StoreKey],
		appCodec,
		app.MsgServiceRouter(),
		app.AccountKeeper,
	)

	app.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec,
		keys[banktypes.StoreKey],
		app.AccountKeeper,
		app.GetSubspace(banktypes.ModuleName),
		app.BlockedAddrs(),
	)

	stakingKeeper := stakingkeeper.NewKeeper(
		appCodec,
		keys[stakingtypes.StoreKey],
		app.AccountKeeper,
		app.BankKeeper,
		app.GetSubspace(stakingtypes.ModuleName),
	)

	app.MintKeeper = mintkeeper.NewKeeper(
		appCodec,
		keys[minttypes.StoreKey],
		app.GetSubspace(minttypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		app.EpochsKeeper,
		authtypes.FeeCollectorName,
	)
	app.MintKeeper.SetHooks(
		minttypes.NewMultiMintHooks(
		// insert mint hooks receivers here
		),
	)

	app.DistrKeeper = distrkeeper.NewKeeper(
		appCodec,
		keys[distrtypes.StoreKey],
		app.GetSubspace(distrtypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		&stakingKeeper,
		&app.SequencersKeeper,
		authtypes.FeeCollectorName,
		// TODO: upgrade to https://github.com/dymensionxyz/dymension-rdk/pull/476
	)

	app.FeeGrantKeeper = feegrantkeeper.NewKeeper(appCodec, keys[feegrant.StoreKey], app.AccountKeeper)
	app.UpgradeKeeper = upgradekeeper.NewKeeper(skipUpgradeHeights, keys[upgradetypes.StoreKey], appCodec, homePath, app.BaseApp, authtypes.NewModuleAddress(govtypes.ModuleName).String())

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	app.StakingKeeper = *stakingKeeper.SetHooks(app.DistrKeeper.Hooks())

	app.EpochsKeeper.SetHooks(
		epochstypes.NewMultiEpochHooks(
			// insert epoch hooks receivers here
			app.MintKeeper.Hooks(),
		),
	)

	app.SequencersKeeper = *seqkeeper.NewKeeper(
		appCodec, keys[seqtypes.StoreKey], app.GetSubspace(seqtypes.ModuleName),
	)

	app.IBCKeeper = ibckeeper.NewKeeper(
		appCodec, keys[ibchost.StoreKey], app.GetSubspace(ibchost.ModuleName), &app.SequencersKeeper, app.UpgradeKeeper, scopedIBCKeeper,
	)

	// Register the proposal types
	// Deprecated: Avoid adding new handlers, instead use the new proposal flow
	// by granting the governance module the right to execute the message.
	// See: https://github.com/cosmos/cosmos-sdk/blob/release/v0.46.x/x/gov/spec/01_concepts.md#proposal-messages
	govRouter := govv1beta1.NewRouter()
	govRouter.
		AddRoute(govtypes.RouterKey, govv1beta1.ProposalHandler).
		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(app.ParamsKeeper)).
		AddRoute(distrtypes.RouterKey, distr.NewCommunityPoolSpendProposalHandler(app.DistrKeeper)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.UpgradeKeeper)).
		AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(app.IBCKeeper.ClientKeeper))

	govConfig := govtypes.DefaultConfig()
	/*
		Example of setting gov params:
		govConfig.MaxMetadataLen = 10000
	*/
	govKeeper := govkeeper.NewKeeper(
		appCodec, keys[govtypes.StoreKey], app.GetSubspace(govtypes.ModuleName), app.AccountKeeper, app.BankKeeper,
		&stakingKeeper, govRouter, app.MsgServiceRouter(), govConfig,
	)

	app.GovKeeper = *govKeeper.SetHooks(
		govtypes.NewMultiGovHooks(
		// register the governance hooks
		),
	)

	app.HubGenesisKeeper = hubgenkeeper.NewKeeper(
		appCodec,
		keys[hubgentypes.StoreKey],
		app.GetSubspace(hubgentypes.ModuleName),
		app.AccountKeeper,
	)

	app.HubKeeper = hubkeeper.NewKeeper(
		appCodec,
		keys[hubtypes.StoreKey],
	)

	denomMetadataMiddleware := denommetadata.NewICS4Wrapper(
		app.IBCKeeper.ChannelKeeper,
		app.HubKeeper,
		app.BankKeeper,
		app.HubGenesisKeeper.GetState,
	)

	genesisTransfersBlocker := hubgenkeeper.NewICS4Wrapper(denomMetadataMiddleware, app.HubGenesisKeeper) // ICS4 Wrapper: claims IBC middleware

	// Create Transfer Keepers
	app.TransferKeeper = ibctransferkeeper.NewKeeper(
		appCodec,
		keys[ibctransfertypes.StoreKey],
		app.GetSubspace(ibctransfertypes.ModuleName),
		genesisTransfersBlocker,
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		scopedTransferKeeper,
	)

	// create IBC module from top to bottom of stack
	var transferStack ibcporttypes.IBCModule

	transferStack = ibctransfer.NewIBCModule(app.TransferKeeper)
	transferStack = denommetadata.NewIBCModule(
		transferStack,
		app.BankKeeper,
		app.TransferKeeper,
		app.HubKeeper,
		denommetadatamoduletypes.NewMultiDenommetadataHooks(),
	)

	transferStack = hubgenkeeper.NewIBCModule(
		transferStack,
		app.TransferKeeper,
		app.HubGenesisKeeper,
		app.BankKeeper,
	)

	app.CallbackKeeper = callbackKeeper.NewKeeper(
		appCodec,
		keys[callbackTypes.StoreKey],
		&app.WasmKeeper,
		app.BankKeeper,
	)

	app.CWErrorsKeeper = cwerrorsKeeper.NewKeeper(
		appCodec,
		keys[cwerrorsTypes.StoreKey],
		tkeys[cwerrorsTypes.TStoreKey],
		&app.WasmKeeper,
		app.BankKeeper,
	)

	wasmDir := filepath.Join(homePath, "wasm")
	wasmConfig, err := wasm.ReadWasmConfig(appOpts)
	if err != nil {
		panic(fmt.Sprintf("error while reading wasm config: %s", err))
	}

	// The gov proposal types can be individually enabled
	if len(enabledProposals) != 0 {
		govRouter.AddRoute(wasm.RouterKey, wasm.NewWasmProposalHandler(app.WasmKeeper, enabledProposals))
	}

	// Include the x/cwerrors query to stargate queries
	wasmOpts = append(wasmOpts, wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Stargate: wasmkeeper.AcceptListStargateQuerier(getAcceptedStargateQueries(), app.GRPCQueryRouter(), appCodec),
	}))

	// The last arguments can contain custom message handlers, and custom query handlers,
	// if we want to allow any custom callbacks
	availableCapabilities := strings.Join(AllCapabilities(), ",")
	app.WasmKeeper = wasmkeeper.NewKeeper(
		appCodec,
		keys[wasmtypes.StoreKey],
		app.GetSubspace(wasm.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		app.StakingKeeper,
		app.DistrKeeper,
		app.IBCKeeper.ChannelKeeper, // may be replaced with middleware such as ics29 feerefunder
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		scopedWasmKeeper,
		app.TransferKeeper,
		app.MsgServiceRouter(),
		app.GRPCQueryRouter(),
		wasmDir,
		wasmConfig,
		availableCapabilities,
		wasmOpts...,
	)

	app.GaslessKeeper = gaslesskeeper.NewKeeper(
		appCodec,
		keys[gaslesstypes.StoreKey],
		app.GetSubspace(gaslesstypes.ModuleName),
		app.interfaceRegistry,
		app.BankKeeper,
		&app.WasmKeeper,
	)

	app.RollappConsensusParamsKeeper = rollappparamskeeper.NewKeeper(
		app.GetSubspace(rollappparamstypes.ModuleName),
	)

	wasmStack := wasm.NewIBCHandler(app.WasmKeeper, app.IBCKeeper.ChannelKeeper, app.IBCKeeper.ChannelKeeper)

	// Create static IBC router, add transfer route, then set and seal it
	ibcRouter := ibcporttypes.NewRouter()
	ibcRouter.AddRoute(ibctransfertypes.ModuleName, transferStack).AddRoute(wasmtypes.ModuleName, wasmStack)
	app.IBCKeeper.SetRouter(ibcRouter)

	/**** Module Options ****/

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.

	modules := []module.AppModule{
		genutil.NewAppModule(
			app.AccountKeeper, app.StakingKeeper, app.BaseApp.DeliverTx,
			encodingConfig.TxConfig,
		),
		auth.NewAppModule(appCodec, app.AccountKeeper, nil),
		authzmodule.NewAppModule(appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		vesting.NewAppModule(app.AccountKeeper, app.BankKeeper),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		gaslessmodule.NewAppModule(appCodec, app.GaslessKeeper),
		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper, app.BankKeeper),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		sequencers.NewAppModule(appCodec, app.SequencersKeeper),
		epochs.NewAppModule(appCodec, app.EpochsKeeper),
		params.NewAppModule(app.ParamsKeeper),
		wasm.NewAppModule(appCodec, &app.WasmKeeper, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		ibc.NewAppModule(app.IBCKeeper),
		ibctransfer.NewAppModule(app.TransferKeeper),
		upgrade.NewAppModule(app.UpgradeKeeper),
		hubgenesis.NewAppModule(appCodec, app.HubGenesisKeeper),
		hub.NewAppModule(appCodec, app.HubKeeper),
		callback.NewAppModule(app.appCodec, app.CallbackKeeper, app.WasmKeeper, app.CWErrorsKeeper),
		cwerrors.NewAppModule(app.appCodec, app.CWErrorsKeeper, app.WasmKeeper),
		rollappparams.NewAppModule(appCodec, app.RollappConsensusParamsKeeper),
	}

	app.mm = module.NewManager(modules...)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	// NOTE: capability module's beginblocker must come before any modules using capabilities (e.g. IBC)
	beginBlockersList := []string{
		upgradetypes.ModuleName,
		capabilitytypes.ModuleName,
		minttypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		seqtypes.ModuleName,
		vestingtypes.ModuleName,
		ibchost.ModuleName,
		ibctransfertypes.ModuleName,
		authtypes.ModuleName,
		authz.ModuleName,
		banktypes.ModuleName,
		govtypes.ModuleName,
		genutiltypes.ModuleName,
		feegrant.ModuleName,
		gaslesstypes.ModuleName,
		epochstypes.ModuleName,
		paramstypes.ModuleName,
		hubgentypes.ModuleName,
		hubtypes.ModuleName,
		wasm.ModuleName,
		callbackTypes.ModuleName,
		cwerrorsTypes.ModuleName, // does not have begin blocker
		rollappparamstypes.ModuleName,
	}
	app.mm.SetOrderBeginBlockers(beginBlockersList...)

	endBlockersList := []string{
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		seqtypes.ModuleName,
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		authz.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		vestingtypes.ModuleName,
		minttypes.ModuleName,
		genutiltypes.ModuleName,
		feegrant.ModuleName,
		gaslesstypes.ModuleName,
		epochstypes.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		ibchost.ModuleName,
		ibctransfertypes.ModuleName,
		hubgentypes.ModuleName,
		hubtypes.ModuleName,
		wasm.ModuleName,
		callbackTypes.ModuleName,
		cwerrorsTypes.ModuleName,
		rollappparamstypes.ModuleName,
	}
	app.mm.SetOrderEndBlockers(endBlockersList...)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: The genutils module must also occur after auth so that it can access the params from auth.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	initGenesisList := []string{
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		authz.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		seqtypes.ModuleName,
		vestingtypes.ModuleName,
		epochstypes.ModuleName,
		govtypes.ModuleName,
		minttypes.ModuleName,
		ibchost.ModuleName,
		genutiltypes.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		ibctransfertypes.ModuleName,
		feegrant.ModuleName,
		gaslesstypes.ModuleName,
		hubgentypes.ModuleName,
		hubtypes.ModuleName,
		wasm.ModuleName,
		callbackTypes.ModuleName,
		cwerrorsTypes.ModuleName,
		rollappparamstypes.ModuleName,
	}
	app.mm.SetOrderInitGenesis(initGenesisList...)

	app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), encodingConfig.Amino)
	app.configurator = module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.mm.RegisterServices(app.configurator)

	// RegisterUpgradeHandlers is used for registering any on-chain upgrades.
	// Make sure it's called after `app.mm` and `app.configurator` are set.
	// app.RegisterUpgradeHandlers()

	// add test gRPC service for testing gRPC queries in isolation
	testdata.RegisterQueryServer(app.GRPCQueryRouter(), testdata.QueryImpl{})

	// create the simulation manager and define the order of the modules for deterministic simulations
	//
	// NOTE: this is not required apps that don't use the simulator for fuzz testing
	// transactions
	overrideModules := map[string]module.AppModuleSimulation{
		authtypes.ModuleName: auth.NewAppModule(app.appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts),
	}
	app.sm = module.NewSimulationManagerFromAppModules(app.mm.Modules, overrideModules)

	app.sm.RegisterStoreDecoders()

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)
	app.setAnteHandler(encodingConfig.TxConfig, wasmConfig)

	if manager := app.SnapshotManager(); manager != nil {
		err := manager.RegisterExtensions(
			wasmkeeper.NewWasmSnapshotter(app.CommitMultiStore(), &app.WasmKeeper),
		)
		if err != nil {
			panic(fmt.Errorf("failed to register snapshot extension: %s", err))
		}
	}
	// In v0.46, the SDK introduces _postHandlers_. PostHandlers are like
	// antehandlers, but are run _after_ the `runMsgs` execution. They are also
	// defined as a chain, and have the same signature as antehandlers.
	//
	// In baseapp, postHandlers are run in the same store branch as `runMsgs`,
	// meaning that both `runMsgs` and `postHandler` state will be committed if
	// both are successful, and both will be reverted if any of the two fails.
	//
	// The SDK exposes a default empty postHandlers chain.
	//
	// Please note that changing any of the anteHandler or postHandler chain is
	// likely to be a state-machine breaking change, which needs a coordinated
	// upgrade.
	app.setPostHandler()

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(err.Error())
		}

		ctx := app.BaseApp.NewUncachedContext(true, tmproto.Header{})

		// Initialize pinned codes in wasmvm as they are not persisted there
		if err := app.WasmKeeper.InitializePinnedCodes(ctx); err != nil {
			panic(fmt.Sprintf("failed initialize pinned codes %s", err))
		}
	}

	app.ScopedIBCKeeper = scopedIBCKeeper
	app.ScopedTransferKeeper = scopedTransferKeeper
	app.ScopedWasmKeeper = scopedWasmKeeper
	return app
}

func (app *App) setAnteHandler(txConfig client.TxConfig, wasmConfig wasmtypes.WasmConfig) {
	anteHandler, err := NewAnteHandler(
		HandlerOptions{
			HandlerOptions: ante.HandlerOptions{
				AccountKeeper:   app.AccountKeeper,
				BankKeeper:      app.BankKeeper,
				FeegrantKeeper:  app.FeeGrantKeeper,
				SignModeHandler: txConfig.SignModeHandler(),
				SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
			},
			IBCKeeper:         app.IBCKeeper,
			WasmConfig:        &wasmConfig,
			TxCounterStoreKey: app.keys[wasmtypes.StoreKey],
			GaslessKeeper:     app.GaslessKeeper,
		},
	)
	if err != nil {
		panic(err)
	}

	app.SetAnteHandler(anteHandler)
}

func (app *App) setPostHandler() {
	postHandler, err := posthandler.NewPostHandler(
		posthandler.HandlerOptions{},
	)
	if err != nil {
		panic(err)
	}

	app.SetPostHandler(postHandler)
}

// Name returns the name of the App
func (app *App) Name() string { return app.BaseApp.Name() }

// BeginBlocker application updates every begin block
func (app *App) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker application updates every end block
func (app *App) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	rollappparams := app.RollappConsensusParamsKeeper.GetParams(ctx)
	abciEndBlockResponse := app.mm.EndBlock(ctx, req)
	abciEndBlockResponse.RollappConsensusParamUpdates = &abci.RollappConsensusParams{
		Da:     rollappparams.Da,
		Commit: rollappparams.Commit,
		Block: &abci.BlockParams{
			MaxGas:   int64(rollappparams.Blockmaxgas),
			MaxBytes: int64(rollappparams.Blockmaxsize),
		},
	}
	return abciEndBlockResponse
}

// InitChainer application update at chain initialization
func (app *App) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	if err := tmjson.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}

	// Passing the dymint sequencers to the sequencer module from RequestInitChain
	if len(req.Validators) == 0 {
		panic(fmt.Sprint("Dymint have no sequencers defined on InitChain, req:", req))
	}
	app.SequencersKeeper.SetDymintSequencers(ctx, req.Validators)

	app.UpgradeKeeper.SetModuleVersionMap(ctx, app.mm.GetVersionMap())
	res := app.mm.InitGenesis(ctx, app.appCodec, genesisState)
	return res
}

// LoadHeight loads a particular height
func (app *App) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *App) ModuleAccountAddrs() map[string]bool {
	ret := make(map[string]bool)
	for acc := range maccPerms {
		ret[authtypes.NewModuleAddress(acc).String()] = true
	}
	return ret
}

// BlockedAddrs returns all the app's module account addresses that are not
// allowed to receive funds. If the map value is true, that account cannot receive funds.
func (app *App) BlockedAddrs() map[string]bool {
	// block all modules by default
	ret := app.ModuleAccountAddrs()
	// delete them if they CAN receive tokens
	for _, acc := range maccCanReceiveTokens {
		delete(ret, authtypes.NewModuleAddress(acc).String())
	}
	return ret
}

// LegacyAmino returns App's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) LegacyAmino() *codec.LegacyAmino {
	return app.cdc
}

// AppCodec returns an app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns an InterfaceRegistry
func (app *App) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetKey(storeKey string) *storetypes.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetTKey(storeKey string) *storetypes.TransientStoreKey {
	return app.tkeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
func (app *App) GetMemKey(storeKey string) *storetypes.MemoryStoreKey {
	return app.memKeys[storeKey]
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// SimulationManager implements the SimulationApp interface
func (app *App) SimulationManager() *module.SimulationManager {
	return app.sm
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *App) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register node gRPC service for grpc-gateway.
	nodeservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register grpc-gateway routes for all modules.
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// register swagger API from root so that other applications can override easily
	if apiConfig.Swagger {
		RegisterSwaggerAPI(clientCtx, apiSvr.Router)
	}
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *App) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *App) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(
		clientCtx,
		app.BaseApp.GRPCQueryRouter(),
		app.interfaceRegistry,
		app.Query,
	)
}

func (app *App) RegisterNodeService(clientCtx client.Context) {
	nodeservice.RegisterNodeService(clientCtx, app.GRPCQueryRouter())
}

// IBC Go TestingApp functions

// GetBaseApp implements the TestingApp interface.
func (app *App) GetBaseApp() *baseapp.BaseApp {
	return app.BaseApp
}

// GetStakingKeeper implements the TestingApp interface.
func (app *App) GetStakingKeeper() ibctestingtypes.StakingKeeper {
	return app.StakingKeeper
}

// GetStakingKeeperSDK implements the TestingApp interface.
func (app *App) GetStakingKeeperSDK() stakingkeeper.Keeper {
	return app.StakingKeeper
}

// GetIBCKeeper implements the TestingApp interface.
func (app *App) GetIBCKeeper() *ibckeeper.Keeper {
	return app.IBCKeeper
}

// GetScopedIBCKeeper implements the TestingApp interface.
func (app *App) GetScopedIBCKeeper() capabilitykeeper.ScopedKeeper {
	return app.ScopedIBCKeeper
}

// GetTxConfig implements the TestingApp interface.
func (app *App) GetTxConfig() client.TxConfig {
	cfg := rollappwasmparams.MakeEncodingConfig()
	return cfg.TxConfig
}

// RegisterSwaggerAPI registers swagger route with API Server
func RegisterSwaggerAPI(_ client.Context, rtr *mux.Router) {
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}

	staticServer := http.FileServer(statikFS)
	rtr.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", staticServer))
}

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey storetypes.StoreKey) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(seqtypes.ModuleName)
	paramsKeeper.Subspace(minttypes.ModuleName)
	paramsKeeper.Subspace(epochstypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govv1.ParamKeyTable())
	paramsKeeper.Subspace(ibctransfertypes.ModuleName)
	paramsKeeper.Subspace(ibchost.ModuleName)
	paramsKeeper.Subspace(hubgentypes.ModuleName)
	paramsKeeper.Subspace(gaslesstypes.ModuleName)

	paramsKeeper.Subspace(wasmtypes.ModuleName)
	paramsKeeper.Subspace(callbackTypes.ModuleName)
	paramsKeeper.Subspace(cwerrorsTypes.ModuleName)
	paramsKeeper.Subspace(rollappparamstypes.ModuleName)
	return paramsKeeper
}

func getAcceptedStargateQueries() wasmkeeper.AcceptedStargateQueries {
	return wasmkeeper.AcceptedStargateQueries{
		"/rollapp.cwerrors.v1.Query/Errors": &cwerrorsTypes.QueryErrorsRequest{},
	}
}
