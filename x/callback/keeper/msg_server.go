package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/dymensionxyz/rollapp-wasm/x/callback/types"
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

// CancelCallback implements types.MsgServer.
func (s MsgServer) CancelCallback(c context.Context, request *types.MsgCancelCallback) (*types.MsgCancelCallbackResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	// If a callback with same job id does not exist, return error
	callback, err := s.keeper.GetCallback(ctx, request.CallbackHeight, request.ContractAddress, request.JobId)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrCallbackNotFound, "callback with given job id does not exist for given height")
	}

	// Deleting the callback from state
	err = s.keeper.DeleteCallback(ctx, request.Sender, callback)
	if err != nil {
		return nil, err
	}

	// Returning the transaction fees + surplus fees as the callback was never executed
	refundFees := callback.FeeSplit.TransactionFees.Add(*callback.FeeSplit.SurplusFees)
	err = s.keeper.RefundFromCallbackModule(ctx, request.Sender, refundFees)
	if err != nil {
		return nil, err
	}

	// Sending the reservation fees to fee collector
	reservationFees := callback.FeeSplit.BlockReservationFees.Add(*callback.FeeSplit.FutureReservationFees)
	err = s.keeper.SendToFeeCollector(ctx, reservationFees)
	if err != nil {
		return nil, err
	}

	// Emit event
	types.EmitCallbackCancelledEvent(
		ctx,
		request.ContractAddress,
		request.JobId,
		request.CallbackHeight,
		request.Sender,
		refundFees,
	)

	return &types.MsgCancelCallbackResponse{
		Refund: refundFees,
	}, nil
}

// RequestCallback implements types.MsgServer.
func (s MsgServer) RequestCallback(c context.Context, request *types.MsgRequestCallback) (*types.MsgRequestCallbackResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	// Get the expected fees which is to be paid
	futureReservationFee, blockReservationFee, transactionFee, err := s.keeper.EstimateCallbackFees(ctx, request.CallbackHeight)
	if err != nil {
		return nil, err
	}
	expectedFees := transactionFee.Add(blockReservationFee).Add(futureReservationFee)

	// If the fees sent by the sender is less than the expected fees, return error
	if request.Fees.IsLT(expectedFees) {
		return nil, errorsmod.Wrapf(types.ErrInsufficientFees, "expected %s, got %s", expectedFees, request.Fees)
	}
	surplusFees := request.Fees.Sub(expectedFees) // Calculating any surplus user has sent

	// Save the callback in state
	callback := types.NewCallback(
		request.Sender,
		request.ContractAddress,
		request.CallbackHeight,
		request.JobId,
		transactionFee,
		blockReservationFee,
		futureReservationFee,
		surplusFees,
	)
	err = s.keeper.SaveCallback(ctx, callback)
	if err != nil {
		return nil, err
	}

	// Send the fees into module account
	err = s.keeper.SendToCallbackModule(ctx, request.Sender, request.Fees)
	if err != nil {
		return nil, err
	}

	// Emit event
	types.EmitCallbackRegisteredEvent(
		ctx,
		request.ContractAddress,
		request.JobId,
		request.CallbackHeight,
		callback.FeeSplit,
		request.Sender,
	)

	return &types.MsgRequestCallbackResponse{}, nil
}

func (s MsgServer) UpdateParams(goCtx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if s.keeper.authority != msg.Authority {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid authority; expected %s, got %s", s.keeper.authority, msg.Authority)
	}

	err := msg.ValidateBasic()
	if err != nil {
		return nil, err
	}

	err = s.keeper.SetParams(ctx, msg.Params)
	if err != nil {
		return nil, err
	}
	return &types.MsgUpdateParamsResponse{}, nil
}
