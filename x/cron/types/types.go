package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errors "github.com/cosmos/cosmos-sdk/types/errors"
	"golang.org/x/exp/slices"
)

var (
	ResolveSinglePlayer = []byte(`{"resolve_bet":{}}`)
	SetupMultiPlayer    = []byte(`{"setup_multiplayer":{}}`)
	ResolveMultiPlayer  = []byte(`{"resolve_multiplayer":{}}`)
)

func NewWhitelistContract(gameId uint64, securityAddress, contractAdmin sdk.AccAddress, gameName string, contractAddress sdk.AccAddress, gameType uint64) WhitelistedContract {
	return WhitelistedContract{
		GameId:          gameId,
		SecurityAddress: securityAddress.String(),
		ContractAdmin:   contractAdmin.String(),
		GameName:        gameName,
		ContractAddress: contractAddress.String(),
		GameType:        gameType,
	}
}

func (m WhitelistedContract) Validate() error {
	if m.GameId == 0 {
		return errorsmod.Wrap(errors.ErrInvalidRequest, "game id must not be 0")
	}
	// check if the security address is valid
	if _, err := sdk.AccAddressFromBech32(m.SecurityAddress); err != nil {
		return errorsmod.Wrapf(errors.ErrInvalidAddress, "invalid security address: %v", err)
	}
	// check if the contract admin is valid
	if _, err := sdk.AccAddressFromBech32(m.ContractAdmin); err != nil {
		return errorsmod.Wrapf(errors.ErrInvalidAddress, "invalid contract admin: %v", err)
	}
	if m.GameName == "" {
		return errorsmod.Wrap(errors.ErrInvalidRequest, "game name must not be empty")
	}
	// check if the contract address is valid
	if _, err := sdk.AccAddressFromBech32(m.ContractAddress); err != nil {
		return errorsmod.Wrapf(errors.ErrInvalidAddress, "invalid ContractAddress: %v", err)
	}
	// check if game type does not contain 1,2,3
	if !slices.Contains([]uint64{1, 2, 3}, m.GameType) {
		return errorsmod.Wrapf(errors.ErrInvalidRequest, "invalid game type, should be 1, 2 or 3")
	}

	return nil
}
