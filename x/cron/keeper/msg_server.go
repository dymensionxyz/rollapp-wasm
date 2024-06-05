package keeper

import (
	"context"
	errorsmod "cosmossdk.io/errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/dymensionxyz/rollapp-wasm/x/cron/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) RegisterContract(goCtx context.Context, msg *types.MsgRegisterContract) (*types.MsgRegisterContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// check if the cron is globally enabled
	params := k.GetParams(ctx)
	if !params.EnableCron {
		return &types.MsgRegisterContractResponse{}, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "Cron is disabled")
	}

	if err := msg.ValidateBasic(); err != nil {
		ctx.Logger().Error(fmt.Sprintf("request invalid: %s", err))
		return &types.MsgRegisterContractResponse{}, err
	}

	// Validation such that only the user who instantiated the contract can register contract
	contractAddr, err := sdk.AccAddressFromBech32(msg.ContractAddress)
	if err != nil {
		return &types.MsgRegisterContractResponse{}, sdkerrors.ErrInvalidAddress
	}
	contractInfo := k.conOps.GetContractInfo(ctx, contractAddr)

	// check if sender is authorized
	exists := k.CheckSecurityAddress(ctx, msg.SecurityAddress)
	if !exists {
		return &types.MsgRegisterContractResponse{}, sdkerrors.ErrUnauthorized
	}

	allContracts := k.GetWhitelistedContracts(ctx)

	for _, data := range allContracts {
		if data.ContractAddress == msg.ContractAddress {
			return &types.MsgRegisterContractResponse{}, errorsmod.Wrapf(sdkerrors.ErrNotFound, "contract already registered")
		}
	}
	gameID := k.GetGameID(ctx)
	contract := types.WhitelistedContract{
		GameId:          gameID + 1,
		SecurityAddress: msg.SecurityAddress,
		ContractAdmin:   contractInfo.Admin,
		GameName:        msg.GameName,
		ContractAddress: msg.ContractAddress,
		GameType:        msg.GameType,
	}

	k.SetContract(ctx, contract)
	k.SetGameID(ctx, gameID+1)

	return &types.MsgRegisterContractResponse{}, nil
}

func (k msgServer) DeregisterContract(goCtx context.Context, msg *types.MsgDeregisterContract) (*types.MsgDeregisterContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)
	if !params.EnableCron {
		return &types.MsgDeregisterContractResponse{}, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "Cron is disabled")
	}

	if err := msg.ValidateBasic(); err != nil {
		ctx.Logger().Error(fmt.Sprintf("request invalid: %s", err))
		return &types.MsgDeregisterContractResponse{}, err
	}

	// Get Game info from Game Id
	gameInfo, found := k.GetWhitelistedContract(ctx, msg.GameId)
	if !found {
		return &types.MsgDeregisterContractResponse{}, errorsmod.Wrapf(sdkerrors.ErrNotFound, "no contract found for this game ID")
	}

	// Validation such that only the user who instantiated the contract can register contract
	contractAddr, err := sdk.AccAddressFromBech32(gameInfo.ContractAddress)
	if err != nil {
		return &types.MsgDeregisterContractResponse{}, sdkerrors.ErrInvalidAddress
	}
	contractInfo := k.conOps.GetContractInfo(ctx, contractAddr)

	// check if sender is authorized
	exists := k.CheckSecurityAddress(ctx, msg.SecurityAddress)
	if !exists && contractInfo.Admin != msg.SecurityAddress {
		return &types.MsgDeregisterContractResponse{}, sdkerrors.ErrUnauthorized
	}

	k.DeleteContract(ctx, msg.GameId)

	return &types.MsgDeregisterContractResponse{}, nil
}
