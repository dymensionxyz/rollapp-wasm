package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/dymensionxyz/rollapp-wasm/x/cron/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = QueryServer{}

type QueryServer struct {
	Keeper
}

func NewQueryServerImpl(k Keeper) types.QueryServer {
	return &QueryServer{
		Keeper: k,
	}
}

func (q QueryServer) QueryWhitelistedContracts(c context.Context, req *types.QueryWhitelistedContractsRequest) (*types.QueryWhitelistedContractsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be empty")
	}

	var (
		items []types.WhitelistedContract
		ctx   = sdk.UnwrapSDKContext(c)
	)

	pagination, err := query.FilteredPaginate(
		prefix.NewStore(q.Store(ctx), types.SetContractKeyPrefix),
		req.Pagination,
		func(_, value []byte, accumulate bool) (bool, error) {
			var item types.WhitelistedContract
			if err := q.cdc.Unmarshal(value, &item); err != nil {
				return false, err
			}

			if accumulate {
				items = append(items, item)
			}

			return true, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryWhitelistedContractsResponse{
		WhilistedContracts: items,
		Pagination:         pagination,
	}, nil
}

func (q QueryServer) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := q.Keeper.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}
