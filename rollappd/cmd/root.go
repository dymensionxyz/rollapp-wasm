package cmd

import (
	"errors"
	"io"
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/config"
	"github.com/cosmos/cosmos-sdk/client/debug"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/server"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	ethermintclient "github.com/evmos/evmos/v12/client"
	evmosconfig "github.com/evmos/evmos/v12/cmd/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	tmcfg "github.com/tendermint/tendermint/config"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	tmlog "github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	etherminthd "github.com/evmos/evmos/v12/crypto/hd"

	berpcconfig "github.com/bcdevtools/block-explorer-rpc-cosmos/be_rpc/config"
	rdkserver "github.com/dymensionxyz/dymension-rdk/server"
	rdkserverconfig "github.com/dymensionxyz/dymension-rdk/server/config"
	"github.com/dymensionxyz/dymension-rdk/utils"
	dymintconf "github.com/dymensionxyz/dymint/config"

	"github.com/dymensionxyz/rollapp-wasm/app"
	"github.com/dymensionxyz/rollapp-wasm/app/params"
)

const rollappAscii = `
██████   ██████  ██      ██       █████  ██████  ██████      ██     ██  █████  ███████ ███    ███ 
██   ██ ██    ██ ██      ██      ██   ██ ██   ██ ██   ██     ██     ██ ██   ██ ██      ████  ████ 
██████  ██    ██ ██      ██      ███████ ██████  ██████      ██  █  ██ ███████ ███████ ██ ████ ██ 
██   ██ ██    ██ ██      ██      ██   ██ ██      ██          ██ ███ ██ ██   ██      ██ ██  ██  ██ 
██   ██  ██████  ███████ ███████ ██   ██ ██      ██           ███ ███  ██   ██ ███████ ██      ██                 
`

// NewRootCmd creates a new root rollapp-wasm command. It is called once in the main function.
func NewRootCmd() (*cobra.Command, params.EncodingConfig) {
	encodingConfig := app.MakeEncodingConfig()

	// TODO: refactor to use depinject

	initClientCtx := client.Context{}.
		WithCodec(encodingConfig.Codec).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithInput(os.Stdin).
		WithAccountRetriever(types.AccountRetriever{}).
		WithHomeDir(app.DefaultNodeHome).
		WithKeyringOptions(etherminthd.EthSecp256k1Option()).
		WithViper("ROLLAPP")

	rootCmd := &cobra.Command{
		// TODO: set by code, not in Makefile
		Use:   version.AppName,
		Short: rollappAscii,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			// set the default command outputs
			cmd.SetOut(cmd.OutOrStdout())
			cmd.SetErr(cmd.ErrOrStderr())

			initClientCtx, err := client.ReadPersistentCommandFlags(initClientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			initClientCtx, err = config.ReadFromClientConfig(initClientCtx)
			if err != nil {
				return err
			}

			if err := client.SetCmdClientContextHandler(initClientCtx, cmd); err != nil {
				return err
			}

			customTMConfig := initTendermintConfig()
			customAppTemplate, customAppConfig := initAppConfig()
			err = server.InterceptConfigsPreRunHandler(
				cmd, customAppTemplate, customAppConfig, customTMConfig,
			)
			if err != nil {
				return err
			}
			serverCtx := server.GetServerContextFromCmd(cmd)

			// create dymint toml config file
			home := serverCtx.Viper.GetString(tmcli.HomeFlag)
			dymintconf.EnsureRoot(home, dymintconf.DefaultConfig(home))

			// create Block Explorer Json-RPC toml config file
			berpcconfig.EnsureRoot(home, berpcconfig.DefaultBeJsonRpcConfig())

			return nil
		},
	}

	initRootCmd(rootCmd, encodingConfig)
	return rootCmd, encodingConfig
}

// initTendermintConfig helps to override default Tendermint Config values.
// return tmcfg.DefaultConfig if no custom configuration is required for the application.
func initTendermintConfig() *tmcfg.Config {
	cfg := tmcfg.DefaultConfig()

	// these values put a higher strain on node memory
	// cfg.P2P.MaxNumInboundPeers = 100
	// cfg.P2P.MaxNumOutboundPeers = 40

	return cfg
}

// initAppConfig helps to override default appConfig template and configs.
// return "", nil if no custom configuration is required for the application.
func initAppConfig() (string, interface{}) {
	customAppTemplate := serverconfig.DefaultConfigTemplate
	srvCfg := serverconfig.DefaultConfig()

	rdkserverconfig.SetDefaultPruningSettings(srvCfg)

	return customAppTemplate, srvCfg
}

func initRootCmd(
	rootCmd *cobra.Command,
	encodingConfig params.EncodingConfig,
) {
	// Set config
	sdkconfig := sdk.GetConfig()
	// Add ability to override bech32 prefix from env variable
	prefix := app.AccountAddressPrefix
	if overridePrefix := os.Getenv("OVERRIDE_BECH32"); overridePrefix != "" {
		prefix = overridePrefix
	}
	utils.SetPrefixes(sdkconfig, prefix)
	evmosconfig.SetBip44CoinType(sdkconfig)
	sdkconfig.Seal()

	ac := appCreator{
		encCfg: encodingConfig,
	}
	rootCmd.AddCommand(
		genutilcli.InitCmd(app.ModuleBasics, app.DefaultNodeHome),
		genutilcli.CollectGenTxsCmd(banktypes.GenesisBalancesIterator{}, app.DefaultNodeHome),
		genutilcli.MigrateGenesisCmd(),
		genutilcli.GenTxCmd(app.ModuleBasics, encodingConfig.TxConfig, banktypes.GenesisBalancesIterator{}, app.DefaultNodeHome),

		genutilcli.ValidateGenesisCmd(app.ModuleBasics),
		AddGenesisAccountCmd(app.DefaultNodeHome),
		tmcli.NewCompletionCmd(rootCmd, true),
		debug.Cmd(),
		config.Cmd(),
	)

	rdkserver.AddRollappCommands(rootCmd, app.DefaultNodeHome, ac.newApp, ac.appExport)
	rootCmd.AddCommand(StartCmd(ac.newApp, app.DefaultNodeHome))

	// add keybase, auxiliary RPC, query, and tx child commands
	rootCmd.AddCommand(
		rpc.StatusCommand(),
		queryCommand(),
		txCommand(),
		ethermintclient.KeyCommands(app.DefaultNodeHome),
	)
}

// queryCommand returns the sub-command to send queries to the app
func queryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "query",
		Aliases:                    []string{"q"},
		Short:                      "Querying subcommands",
		DisableFlagParsing:         false,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcmd.GetAccountCmd(),
		rpc.ValidatorCommand(),
		rpc.BlockCommand(),
		authcmd.QueryTxsByEventsCmd(),
		authcmd.QueryTxCmd(),
		genesisChecksumCmd(),
	)

	app.ModuleBasics.AddQueryCommands(cmd)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

// txCommand returns the sub-command to send transactions to the app
func txCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "tx",
		Short:                      "Transactions subcommands",
		DisableFlagParsing:         false,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcmd.GetSignCommand(),
		authcmd.GetSignBatchCommand(),
		authcmd.GetMultiSignCommand(),
		authcmd.GetMultiSignBatchCmd(),
		authcmd.GetValidateSignaturesCommand(),
		authcmd.GetBroadcastCommand(),
		authcmd.GetEncodeCommand(),
		authcmd.GetDecodeCommand(),
		authcmd.GetAuxToFeeCommand(),
	)

	app.ModuleBasics.AddTxCommands(cmd)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

type appCreator struct {
	encCfg params.EncodingConfig
}

func (ac appCreator) newApp(
	logger tmlog.Logger,
	db dbm.DB,
	traceStore io.Writer,
	appOpts servertypes.AppOptions,
) servertypes.Application {
	baseappOptions := server.DefaultBaseappOptions(appOpts)

	skipUpgradeHeights := make(map[int64]bool)
	for _, h := range cast.ToIntSlice(appOpts.Get(server.FlagUnsafeSkipUpgrades)) {
		skipUpgradeHeights[int64(h)] = true
	}

	var wasmOpts []wasm.Option
	if cast.ToBool(appOpts.Get("telemetry.enabled")) {
		wasmOpts = append(wasmOpts, wasmkeeper.WithVMCacheMetrics(prometheus.DefaultRegisterer))
	}

	return app.NewRollapp(
		logger,
		db,
		traceStore,
		true,
		skipUpgradeHeights,
		cast.ToString(appOpts.Get(flags.FlagHome)),
		cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod)),
		ac.encCfg,
		app.GetEnabledProposals(),
		appOpts,
		wasmOpts,
		baseappOptions...)
}

func (ac appCreator) appExport(
	logger tmlog.Logger,
	db dbm.DB,
	traceStore io.Writer,
	height int64,
	forZeroHeight bool,
	jailAllowedAddrs []string,
	appOpts servertypes.AppOptions,
) (servertypes.ExportedApp, error) {
	var rollapp *app.App
	homePath, ok := appOpts.Get(flags.FlagHome).(string)
	if !ok || homePath == "" {
		return servertypes.ExportedApp{}, errors.New("application home not set")
	}

	loadLatest := height == -1
	var emptyWasmOpts []wasm.Option
	rollapp = app.NewRollapp(
		logger,
		db,
		traceStore,
		loadLatest,
		map[int64]bool{},
		homePath,
		cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod)),
		ac.encCfg,
		app.GetEnabledProposals(),
		appOpts,
		emptyWasmOpts,
	)

	if height != -1 {
		if err := rollapp.LoadHeight(height); err != nil {
			return servertypes.ExportedApp{}, err
		}
	}

	return rollapp.ExportAppStateAndValidators(forZeroHeight, jailAllowedAddrs)
}
