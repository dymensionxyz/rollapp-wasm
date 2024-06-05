package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeregisterContract = "de_register_contract"

var _ sdk.Msg = &MsgDeregisterContract{}

func NewMsgDeregisterContract(
	securityAddress string,
	gameID uint64,
) *MsgDeregisterContract {
	return &MsgDeregisterContract{
		SecurityAddress: securityAddress,
		GameId:          gameID,
	}
}

func (msg *MsgDeregisterContract) Route() string {
	return RouterKey
}

func (msg *MsgDeregisterContract) Type() string {
	return TypeMsgDeregisterContract
}

func (msg *MsgDeregisterContract) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.SecurityAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeregisterContract) GetSignBytes() []byte {
	bz := moduleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeregisterContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.SecurityAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid security address (%s)", err)
	}

	return nil
}
