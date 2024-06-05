package cron

import (
	errorsmod "cosmossdk.io/errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/dymensionxyz/rollapp-wasm/x/cron/keeper"
	"github.com/dymensionxyz/rollapp-wasm/x/cron/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	// this line is used by starport scaffolding # handler/msgServer

	server := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgRegisterContract:
			res, err := server.RegisterContract(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgDeregisterContract:
			res, err := server.DeregisterContract(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		// this line is used by starport scaffolding # 1
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, errorsmod.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
