package backend

import (
	"context"
	"github.com/bcdevtools/block-explorer-rpc-cosmos/be_rpc/config"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	sequencerstypes "github.com/dymensionxyz/dymension-rdk/x/sequencers/types"
	rawberpctypes "github.com/dymensionxyz/rollapp-wasm/ra_wasm_be_rpc/types"
	"github.com/tendermint/tendermint/libs/log"
)

type RollAppWasmBackendI interface {
	// Misc

	GetSequencersModuleParams() (*sequencerstypes.Params, error)
}

var _ RollAppWasmBackendI = (*RollAppWasmBackend)(nil)

// RollAppWasmBackend implements the RollAppWasmBackendI interface
type RollAppWasmBackend struct {
	ctx         context.Context
	clientCtx   client.Context
	queryClient *rawberpctypes.QueryClient // gRPC query client
	logger      log.Logger
	cfg         config.BeJsonRpcConfig
}

// NewRollAppWasmBackend creates a new RollAppWasmBackend instance for RollApp EVM Block Explorer
func NewRollAppWasmBackend(
	ctx *server.Context,
	logger log.Logger,
	clientCtx client.Context,
) *RollAppWasmBackend {
	appConf, err := config.GetConfig(ctx.Viper)
	if err != nil {
		panic(err)
	}

	return &RollAppWasmBackend{
		ctx:         context.Background(),
		clientCtx:   clientCtx,
		queryClient: rawberpctypes.NewQueryClient(clientCtx),
		logger:      logger.With("module", "raw_be_rpc"),
		cfg:         appConf,
	}
}
