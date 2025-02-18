package app

import (
	"fmt"
	"runtime/debug"

	errorsmod "cosmossdk.io/errors"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	conntypes "github.com/cosmos/ibc-go/v6/modules/core/03-connection/types"
	ibckeeper "github.com/cosmos/ibc-go/v6/modules/core/keeper"
	rdkante "github.com/dymensionxyz/dymension-rdk/server/ante"
	distrkeeper "github.com/dymensionxyz/dymension-rdk/x/dist/keeper"
	"github.com/dymensionxyz/dymension-rdk/x/gasless"
	gaslesskeeper "github.com/dymensionxyz/dymension-rdk/x/gasless/keeper"
	rollappparamskeeper "github.com/dymensionxyz/dymension-rdk/x/rollappparams/keeper"
	seqkeeper "github.com/dymensionxyz/dymension-rdk/x/sequencers/keeper"
	cosmosante "github.com/evmos/evmos/v12/app/ante/cosmos"
	evmostypes "github.com/evmos/evmos/v12/types"
	evmtypes "github.com/evmos/evmos/v12/x/evm/types"
	tmlog "github.com/tendermint/tendermint/libs/log"
)

// HandlerOptions are the options required for constructing a default SDK AnteHandler.
type HandlerOptions struct {
	ante.HandlerOptions

	IBCKeeper           *ibckeeper.Keeper
	WasmConfig          *wasmtypes.WasmConfig
	TxCounterStoreKey   storetypes.StoreKey
	GaslessKeeper       gaslesskeeper.Keeper
	DistrKeeper         distrkeeper.Keeper
	SequencersKeeper    seqkeeper.Keeper
	RollappParamsKeeper rollappparamskeeper.Keeper
}

// NewAnteHandler returns an AnteHandler that checks and increments sequence
// numbers, checks signatures & account numbers, and deducts fees from the first
// signer.
func NewAnteHandler(options HandlerOptions) (sdk.AnteHandler, error) {
	if err := options.validate(); err != nil {
		return nil, fmt.Errorf("options validate: %w", err)
	}

	return func(
		ctx sdk.Context, tx sdk.Tx, sim bool,
	) (newCtx sdk.Context, err error) {
		var anteHandler sdk.AnteHandler

		defer Recover(ctx.Logger(), &err)

		txWithExtensions, ok := tx.(ante.HasExtensionOptionsTx)
		if ok {
			opts := txWithExtensions.GetExtensionOptions()
			if len(opts) > 0 {
				switch typeURL := opts[0].GetTypeUrl(); typeURL {
				case "/ethermint.types.v1.ExtensionOptionsWeb3Tx":
					// Deprecated: Handle as normal Cosmos SDK tx, except signature is checked for Legacy EIP712 representation
					options.ExtensionOptionChecker = func(c *codectypes.Any) bool {
						_, ok := c.GetCachedValue().(*evmostypes.ExtensionOptionsWeb3Tx)
						return ok
					}
					anteHandler = cosmosHandler(
						options,
						// nolint:staticcheck
						cosmosante.NewLegacyEip712SigVerificationDecorator(options.AccountKeeper.(evmtypes.AccountKeeper), options.SignModeHandler), // Use old signature verification: uses EIP instead of the cosmos signature validator
					)
				default:
					return ctx, errorsmod.Wrapf(
						sdkerrors.ErrUnknownExtensionOptions,
						"rejecting tx with unsupported extension option: %s", typeURL,
					)
				}

				return anteHandler(ctx, tx, sim)
			}
		}

		// handle as totally normal Cosmos SDK tx
		switch tx.(type) {
		case sdk.Tx:
			// we reject any extension
			anteHandler = cosmosHandler(
				options,
				ante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler), // Use modern signature verification
			)
		default:
			return ctx, errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "invalid transaction type: %T", tx)
		}

		return anteHandler(ctx, tx, sim)
	}, nil
}

func cosmosHandler(options HandlerOptions, sigChecker sdk.AnteDecorator) sdk.AnteHandler {
	sigGasConsumer := options.SigGasConsumer
	if sigGasConsumer == nil {
		sigGasConsumer = ante.DefaultSigVerificationGasConsumer
	}
	// only override the modern sig checker, and preserve the legacy one
	if _, ok := sigChecker.(ante.SigVerificationDecorator); ok {
		sigChecker = NewSigCheckDecorator(options.AccountKeeper.(accountKeeper), options.SignModeHandler)
	}
	return sdk.ChainAnteDecorators(
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
		rdkante.NewBypassIBCFeeDecorator(
			gasless.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper, options.TxFeeChecker, options.GaslessKeeper),
			options.DistrKeeper,
			options.SequencersKeeper,
			options.RollappParamsKeeper,
		),
		ante.NewSetPubKeyDecorator(options.AccountKeeper), // SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		ante.NewSigGasConsumeDecorator(options.AccountKeeper, sigGasConsumer),
		sigChecker,
		ante.NewIncrementSequenceDecorator(options.AccountKeeper),
	)
}

func (o HandlerOptions) validate() error {
	// From x/auth/ante.go
	if o.AccountKeeper == nil {
		return errorsmod.Wrap(sdkerrors.ErrLogic, "account keeper is required for ante builder")
	}

	if o.BankKeeper == nil {
		return errorsmod.Wrap(sdkerrors.ErrLogic, "bank keeper is required for ante builder")
	}

	if o.SignModeHandler == nil {
		return errorsmod.Wrap(sdkerrors.ErrLogic, "sign mode handler is required for ante builder")
	}

	if o.WasmConfig == nil {
		return errorsmod.Wrap(sdkerrors.ErrLogic, "wasm config is required for ante builder")
	}

	if o.TxCounterStoreKey == nil {
		return errorsmod.Wrap(sdkerrors.ErrLogic, "tx counter store key is required for ante builder")
	}

	if o.SigGasConsumer == nil {
		return errorsmod.Wrap(sdkerrors.ErrLogic, "signature gas consumer is required for ante builder")
	}

	return nil
}

func Recover(logger tmlog.Logger, err *error) {
	if r := recover(); r != nil {
		*err = errorsmod.Wrapf(sdkerrors.ErrPanic, "%v", r)

		if e, ok := r.(error); ok {
			logger.Error(
				"ante handler panicked",
				"error", e,
				"stack trace", string(debug.Stack()),
			)
		} else {
			logger.Error(
				"ante handler panicked",
				"recover", fmt.Sprintf("%v", r),
			)
		}
	}
}
