package types

import "fmt"

func NewGenesisState(whitelistedContracts []WhitelistedContract, params Params) *GenesisState {
	return &GenesisState{
		WhitelistedContracts: whitelistedContracts,
		Params:               params,
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(
		[]WhitelistedContract{},
		DefaultParams(),
	)
}

func (genState *GenesisState) Validate() error {
	if err := genState.Params.Validate(); err != nil {
		return fmt.Errorf("invalid params: %w", err)
	}
	// validates all the whitelisted contracts
	for i, contract := range genState.WhitelistedContracts {
		if err := contract.Validate(); err != nil {
			return fmt.Errorf("invalid whitelisted contract %d: %w", i, err)
		}
	}

	return nil
}
