package types

import (
	"fmt"
	// errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	DefaultSecurityAddress = []string{"aib1sdggcl7eanaanvcsmvars7l0unsge65wzjm3dc"} //assuming BECH32_PREFIX=aib, will need a Default admin to whitelist contract for cron operations

	DefaultContractGasLimit uint64 = 1000000000
	// KeySecurityAddress is store's key for SecurityAddress Params
	KeySecurityAddress = []byte("SecurityAddress")
	// KeyContractGasLimit is store's key for ContractGasLimit Params
	KeyContractGasLimit = []byte("ContractGasLimit")
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(
	securityAddress []string, contractGasLimit uint64,
) Params {
	return Params{
		SecurityAddress:  securityAddress,
		ContractGasLimit: contractGasLimit,
	}
}

// default minting module parameters
func DefaultParams() Params {
	return NewParams(
		DefaultSecurityAddress, DefaultContractGasLimit,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeySecurityAddress, &p.SecurityAddress, validateSecurityAddress),
		paramtypes.NewParamSetPair(KeyContractGasLimit, &p.ContractGasLimit, validateContractGasLimit),
	}
}

// validateSecurityAddress validates that the security addressess are valid
func validateSecurityAddress(i interface{}) error {
	v, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	for _, addr := range v {
		if _, err := sdk.AccAddressFromBech32(addr); err != nil {
			return fmt.Errorf("invalid security address: %s", err.Error())
		}
	}
	return nil
}

func validateContractGasLimit(i interface{}) error {

	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v < 100_000 {
		return fmt.Errorf("invalid contract gas limit must be above 100_000: %d", v)
	}
	return nil
}

// Validate all params
func (p Params) Validate() error {
	for _, field := range []struct {
		val          interface{}
		validateFunc func(i interface{}) error
	}{
		{p.ContractGasLimit, validateContractGasLimit},
		{p.SecurityAddress, validateSecurityAddress},
	} {
		if err := field.validateFunc(field.val); err != nil {
			return err
		}
	}

	return nil
}

