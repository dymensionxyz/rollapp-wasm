package keeper

import (
	"context"
	
	sdk "github.com/cosmos/cosmos-sdk/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/dymensionxyz/rollapp-wasm/x/cwerrors/types"
)

var _ types.MsgServer = (*MsgServer)(nil)

// MsgServer implements the module gRPC messaging service.
type MsgServer struct {
	keeper Keeper
}

// NewMsgServer creates a new gRPC messaging server.
func NewMsgServer(keeper Keeper) *MsgServer {
	return &MsgServer{
		keeper: keeper,
	}
}

// SubscribeToError implements types.MsgServer.
func (s *MsgServer) SubscribeToError(c context.Context, request *types.MsgSubscribeToError) (*types.MsgSubscribeToErrorResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	sender, err := sdk.AccAddressFromBech32(request.Sender)
	if err != nil {
		return nil, err
	}

	contractAddr, err := sdk.AccAddressFromBech32(request.ContractAddress)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(c)
	subscriptionEndHeight, err := s.keeper.SetSubscription(ctx, sender, contractAddr, request.Fee)
	if err != nil {
		return nil, err
	}

	types.EmitSubscribedToErrorsEvent(
		ctx,
		request.Sender,
		request.ContractAddress,
		request.Fee,
		subscriptionEndHeight,
	)
	return &types.MsgSubscribeToErrorResponse{
		SubscriptionValidTill: subscriptionEndHeight,
	}, nil
}