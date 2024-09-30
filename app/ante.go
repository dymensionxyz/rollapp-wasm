package app

import (
	errorsmod "cosmossdk.io/errors"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	conntypes "github.com/cosmos/ibc-go/v6/modules/core/03-connection/types"
	ibcante "github.com/cosmos/ibc-go/v6/modules/core/ante"
	ibckeeper "github.com/cosmos/ibc-go/v6/modules/core/keeper"
	"github.com/dymensionxyz/dymension-rdk/x/gasless"
	gaslesskeeper "github.com/dymensionxyz/dymension-rdk/x/gasless/keeper"
	cosmosante "github.com/evmos/evmos/v12/app/ante/cosmos"
)

// HandlerOptions are the options required for constructing a default SDK AnteHandler.
type HandlerOptions struct {
	ante.HandlerOptions

	IBCKeeper         *ibckeeper.Keeper
	WasmConfig        *wasmtypes.WasmConfig
	TxCounterStoreKey storetypes.StoreKey
	GaslessKeeper     gaslesskeeper.Keeper
}

func GetAnteDecorators(options HandlerOptions) []sdk.AnteDecorator {
	sigGasConsumer := options.SigGasConsumer
	if sigGasConsumer == nil {
		sigGasConsumer = ante.DefaultSigVerificationGasConsumer
	}

	anteDecorators := []sdk.AnteDecorator{
		ante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		cosmosante.NewRejectMessagesDecorator(
			[]string{
				sdk.MsgTypeURL(&conntypes.MsgConnectionOpenInit{}), // don't let any connection open from the Rollapp side (it's still possible from the other side)
			},
		),
		wasmkeeper.NewLimitSimulationGasDecorator(options.WasmConfig.SimulationGasLimit), // after setup context to enforce limits early
		wasmkeeper.NewCountTXDecorator(options.TxCounterStoreKey),
		ante.NewExtensionOptionsDecorator(options.ExtensionOptionChecker),

		ante.NewValidateBasicDecorator(),
		ante.NewTxTimeoutHeightDecorator(),

		ante.NewValidateMemoDecorator(options.AccountKeeper),
		NewCreateAccountDecorator(options.AccountKeeper.(accountKeeper)),
		ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		NewBypassIBCFeeDecorator(gasless.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper, options.TxFeeChecker, options.GaslessKeeper)),
		ante.NewSetPubKeyDecorator(options.AccountKeeper), // SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		ante.NewSigGasConsumeDecorator(options.AccountKeeper, sigGasConsumer),
		NewSigCheckDecorator(options.AccountKeeper.(accountKeeper), options.SignModeHandler),
		ante.NewIncrementSequenceDecorator(options.AccountKeeper),
	}

	anteDecorators = append(anteDecorators, ibcante.NewRedundantRelayDecorator(options.IBCKeeper))

	return anteDecorators
}

// NewAnteHandler returns an AnteHandler that checks and increments sequence
// numbers, checks signatures & account numbers, and deducts fees from the first
// signer.
func NewAnteHandler(options HandlerOptions) (sdk.AnteHandler, error) {
	// From x/auth/ante.go
	if options.AccountKeeper == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "account keeper is required for ante builder")
	}

	if options.BankKeeper == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "bank keeper is required for ante builder")
	}

	if options.SignModeHandler == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "sign mode handler is required for ante builder")
	}

	if options.WasmConfig == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "wasm config is required for ante builder")
	}

	return sdk.ChainAnteDecorators(GetAnteDecorators(options)...), nil
}
