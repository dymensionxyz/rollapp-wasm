package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	DefaultSecurityAddress []string
	// KeySecurityAddress is store's key for SecurityAddress Params
	KeySecurityAddress = []byte("SecurityAddress")
	KeyEnableCron      = []byte("EnableCron")
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(securityAddress []string, enableCron bool) Params {
	return Params{
		SecurityAddress: securityAddress,
		EnableCron:      enableCron,
	}
}

// DefaultParams default minting module parameters
func DefaultParams() Params {
	return NewParams(DefaultSecurityAddress, true)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeySecurityAddress, &p.SecurityAddress, validateSecurityAddress),
		paramtypes.NewParamSetPair(KeyEnableCron, &p.EnableCron, validateBool),
	}
}

// validateSecurityAddress validates that the security addresses are valid
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

func validateBool(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

// Validate all params
func (p Params) Validate() error {
	for _, field := range []struct {
		val          interface{}
		validateFunc func(i interface{}) error
	}{
		{p.SecurityAddress, validateSecurityAddress},
		{p.EnableCron, validateBool},
	} {
		if err := field.validateFunc(field.val); err != nil {
			return err
		}
	}

	return nil
}
