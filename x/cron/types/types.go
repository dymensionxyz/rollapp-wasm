package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
		return fmt.Errorf("invalid GameId: %d", m.GameId)
	}
	// check if the security address is valid
	if _, err := sdk.AccAddressFromBech32(m.SecurityAddress); err != nil {
		return fmt.Errorf("invalid SecurityAddress: %s", m.SecurityAddress)
	}
	// check if the contract admin is valid
	if _, err := sdk.AccAddressFromBech32(m.ContractAdmin); err != nil {
		return fmt.Errorf("invalid ContractAdmin: %s", m.ContractAdmin)
	}
	if m.GameName == "" {
		return fmt.Errorf("invalid GameName: %s", m.GameName)
	}
	// check if the contract address is valid
	if _, err := sdk.AccAddressFromBech32(m.ContractAddress); err != nil {
		return fmt.Errorf("invalid ContractAddress: %s", m.ContractAddress)
	}
	// check if game type does not contain 1,2,3 
	if !slices.Contains([]uint64{1, 2, 3}, m.GameType) {
		return fmt.Errorf("invalid GameType: %d", m.GameType)
	}

	return nil
}
