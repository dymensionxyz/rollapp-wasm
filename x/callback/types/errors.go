package types

import (
	errorsmod "cosmossdk.io/errors"

	cwerrortypes "github.com/dymensionxyz/rollapp-wasm/x/cwerrors/types"
)

var (
	DefaultCodespace                = ModuleName
	ErrContractNotFound             = errorsmod.Register(DefaultCodespace, 2, "contract with given address not found")
	ErrCallbackJobIDExists          = errorsmod.Register(DefaultCodespace, 3, "callback with given job id already exists for given height")
	ErrCallbackHeightNotInFuture    = errorsmod.Register(DefaultCodespace, 4, "callback request height is not in the future")
	ErrUnauthorized                 = errorsmod.Register(DefaultCodespace, 5, "sender not authorized to register callback")
	ErrCallbackNotFound             = errorsmod.Register(DefaultCodespace, 6, "callback with given job id does not exist for given height")
	ErrInsufficientFees             = errorsmod.Register(DefaultCodespace, 7, "insufficient fees to register callback")
	ErrCallbackExists               = errorsmod.Register(DefaultCodespace, 8, "callback with given job id already exists for given height")
	ErrCallbackHeightTooFarInFuture = errorsmod.Register(DefaultCodespace, 9, "callback request height is too far in the future")
	ErrBlockFilled                  = errorsmod.Register(DefaultCodespace, 10, "block filled with max capacity of callbacks")
)

// NewSudoError creates a new sudo error instance to pass on to the errors module
func NewSudoError(errorCode ModuleErrors, contractAddr string, inputPayload string, errMsg string) cwerrortypes.SudoError {
	return cwerrortypes.SudoError{
		ModuleName:      ModuleName,
		ErrorCode:       int32(errorCode),
		ContractAddress: contractAddr,
		InputPayload:    inputPayload,
		ErrorMessage:    errMsg,
	}
}
