package app

import (
	"fmt"
	"runtime/debug"

	errorsmod "cosmossdk.io/errors"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/crypto/types/multisig"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	conntypes "github.com/cosmos/ibc-go/v6/modules/core/03-connection/types"
	ibcante "github.com/cosmos/ibc-go/v6/modules/core/ante"
	ibckeeper "github.com/cosmos/ibc-go/v6/modules/core/keeper"
	"github.com/dymensionxyz/dymension-rdk/x/gasless"
	gaslesskeeper "github.com/dymensionxyz/dymension-rdk/x/gasless/keeper"
	cosmosante "github.com/evmos/evmos/v12/app/ante/cosmos"
	"github.com/evmos/evmos/v12/crypto/ethsecp256k1"
	tmlog "github.com/tendermint/tendermint/libs/log"
)

// HandlerOptions are the options required for constructing a default SDK AnteHandler.
type HandlerOptions struct {
	ante.HandlerOptions

	IBCKeeper         *ibckeeper.Keeper
	WasmConfig        *wasmtypes.WasmConfig
	TxCounterStoreKey storetypes.StoreKey
	GaslessKeeper     gaslesskeeper.Keeper
}

func NewAnteHandler(options HandlerOptions) (sdk.AnteHandler, error) {
	if err := options.validate(); err != nil {
		return nil, fmt.Errorf("options validate: %w", err)
	}

	return func(ctx sdk.Context, tx sdk.Tx, sim bool) (_ sdk.Context, err error) {
		defer Recover(ctx.Logger(), &err)

		return cosmosHandler(options)(ctx, tx, sim)
	}, nil
}

func cosmosHandler(options HandlerOptions) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(getAnteDecorators(options)...)
}

type SigGasConsumer func(meter sdk.GasMeter, sig signing.SignatureV2, params types.Params) error

func getAnteDecorators(options HandlerOptions) []sdk.AnteDecorator {
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
		ante.NewSigGasConsumeDecorator(options.AccountKeeper, defaultSigVerificationGasConsumer),
		NewSigCheckDecorator(options.AccountKeeper.(accountKeeper), options.SignModeHandler),
		ante.NewIncrementSequenceDecorator(options.AccountKeeper),
	}

	anteDecorators = append(anteDecorators, ibcante.NewRedundantRelayDecorator(options.IBCKeeper))

	return anteDecorators
}

const (
	secp256k1VerifyCost uint64 = 21000
)

// TODO: check with zero fee relayer
// Copied from github.com/evmos/ethermint
func defaultSigVerificationGasConsumer(meter sdk.GasMeter, sig signing.SignatureV2, params types.Params) error {
	pubkey := sig.PubKey
	switch pubkey := pubkey.(type) {
	case *ethsecp256k1.PubKey:
		meter.ConsumeGas(secp256k1VerifyCost, "ante verify: eth_secp256k1")
		return nil

	case multisig.PubKey:
		// Multisig keys
		multisignature, ok := sig.Data.(*signing.MultiSignatureData)
		if !ok {
			return fmt.Errorf("expected %T, got, %T", &signing.MultiSignatureData{}, sig.Data)
		}
		return consumeMultisignatureVerificationGas(meter, multisignature, pubkey, params, sig.Sequence)

	default:
		return ante.DefaultSigVerificationGasConsumer(meter, sig, params)
	}
}

// Copied from github.com/evmos/ethermint
func consumeMultisignatureVerificationGas(
	meter sdk.GasMeter, sig *signing.MultiSignatureData, pubkey multisig.PubKey,
	params types.Params, accSeq uint64,
) error {
	size := sig.BitArray.Count()
	sigIndex := 0

	for i := 0; i < size; i++ {
		if !sig.BitArray.GetIndex(i) {
			continue
		}
		sigV2 := signing.SignatureV2{
			PubKey:   pubkey.GetPubKeys()[i],
			Data:     sig.Signatures[sigIndex],
			Sequence: accSeq,
		}
		err := defaultSigVerificationGasConsumer(meter, sigV2, params)
		if err != nil {
			return err
		}
		sigIndex++
	}

	return nil
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
