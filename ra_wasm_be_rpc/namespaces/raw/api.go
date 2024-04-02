package raw

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/server"
	rawberpcbackend "github.com/dymensionxyz/rollapp-wasm/ra_wasm_be_rpc/backend"
	"github.com/tendermint/tendermint/libs/log"
)

// RPC namespaces and API version
const (
	DymRollAppWasmBlockExplorerNamespace = "raw"

	ApiVersion = "1.0"
)

// API is the RollApp Wasm Block Explorer JSON-RPC.
// Developers can create custom API for the chain.
type API struct {
	ctx     *server.Context
	logger  log.Logger
	backend rawberpcbackend.RollAppWasmBackendI
}

// NewRollAppWasmAPI creates an instance of the RollApp Wasm Block Explorer API.
func NewRollAppWasmAPI(
	ctx *server.Context,
	backend rawberpcbackend.RollAppWasmBackendI,
) *API {
	return &API{
		ctx:     ctx,
		logger:  ctx.Logger.With("api", "raw"),
		backend: backend,
	}
}

func (api *API) Echo(text string) string {
	api.logger.Debug("raw_echo")
	return fmt.Sprintf("hello \"%s\" from RollApp Wasm Block Explorer API", text)
}
