package app

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/types/multisig"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/evmos/evmos/v12/crypto/ethsecp256k1"
)

const (
	secp256k1VerifyCost uint64 = 21000
)

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
