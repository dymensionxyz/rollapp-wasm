package types

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/tx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	epochstypes "github.com/dymensionxyz/dymension-rdk/x/epochs/types"
	sequencerstypes "github.com/dymensionxyz/dymension-rdk/x/sequencers/types"
)

// QueryClient defines a gRPC Client used for:
//   - Transaction simulation
type QueryClient struct {
	tx.ServiceClient

	BankQueryClient       banktypes.QueryClient
	SequencersQueryClient sequencerstypes.QueryClient
	EpochQueryClient      epochstypes.QueryClient
}

// NewQueryClient creates a new gRPC query client
func NewQueryClient(clientCtx client.Context) *QueryClient {
	return &QueryClient{
		ServiceClient:         tx.NewServiceClient(clientCtx),
		BankQueryClient:       banktypes.NewQueryClient(clientCtx),
		SequencersQueryClient: sequencerstypes.NewQueryClient(clientCtx),
		EpochQueryClient:      epochstypes.NewQueryClient(clientCtx),
	}
}
