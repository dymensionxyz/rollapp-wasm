package keeper

import (
	"encoding/hex"
	"golang.org/x/exp/slices"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	// "github.com/dymensionxyz/dymension-rdk/utils/logger"
	"github.com/dymensionxyz/rollapp-wasm/x/cron/expected"
	"github.com/dymensionxyz/rollapp-wasm/x/cron/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
		conOps     expected.ContractOpsKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	conOps expected.ContractOpsKeeper,

) Keeper {
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{

		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
		conOps:     conOps,
	}
}

//nolint:staticcheck
func (k Keeper) SudoContractCall(ctx sdk.Context, contractAddress string, p []byte) error {

	contractAddr, err := sdk.AccAddressFromBech32(contractAddress)
	if err != nil {
		return sdkerrors.Wrapf(err, "contract")
	}
	data, err := k.conOps.Sudo(ctx, contractAddr, p)
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeContractSudoMsg,
		sdk.NewAttribute(types.AttributeKeyResultDataHex, hex.EncodeToString(data)),
	))
	return nil
}

func (k Keeper) CheckSecurityAddress(ctx sdk.Context, from string) bool {
	params := k.GetParams(ctx)
	return slices.Contains(params.SecurityAddress, from)
}

func (k Keeper) Store(ctx sdk.Context) sdk.KVStore {
	return ctx.KVStore(k.storeKey)
}

func (k Keeper) SinglePlayer(ctx sdk.Context, contractAddress string, ResolveSinglePlayer []byte, gameName string) {
	err := k.SudoContractCall(ctx, contractAddress, ResolveSinglePlayer)
	if err != nil {
		ctx.Logger().Error("Game %s contract call error for single-player", gameName)
	} else {
		ctx.Logger().Info("Game %s contract call for single-player success", gameName)
	}
}

func (k Keeper) MultiPlayer(ctx sdk.Context, contractAddress string, SetupMultiPlayer []byte, ResolveMultiPlayer []byte, gameName string) {
	err := k.SudoContractCall(ctx, contractAddress, SetupMultiPlayer)
	if err != nil {
		ctx.Logger().Error("Game %s contract call error for setup multi-player", gameName)
	} else {
		ctx.Logger().Info("Game %s contract call for setup multi-player success", gameName)
	}

	err = k.SudoContractCall(ctx, contractAddress, ResolveMultiPlayer)
	if err != nil {
		ctx.Logger().Error("Game %s contract call error for resolve multi-player", gameName)
	} else {
		ctx.Logger().Info("Game %s contract call for resolve multi-player success", gameName)
	}
}
