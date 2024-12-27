package types

import (
	errorsmod "cosmossdk.io/errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgUpdateParams = "update_params"
)

var (
	_ sdk.Msg = &MsgRequestCallback{}
	_ sdk.Msg = &MsgCancelCallback{}
	_ sdk.Msg = &MsgUpdateParams{}
)

// NewMsgRequestCallback creates a new MsgRequestCallback instance.
func NewMsgRequestCallback(
	senderAddr sdk.AccAddress,
	contractAddr sdk.AccAddress,
	jobId uint64,
	callbackHeight int64,
	fees sdk.Coin,
) *MsgRequestCallback {
	msg := &MsgRequestCallback{
		Sender:          senderAddr.String(),
		ContractAddress: contractAddr.String(),
		JobId:           jobId,
		CallbackHeight:  callbackHeight,
		Fees:            fees,
	}

	return msg
}

// GetSigners implements the sdk.Msg interface.
func (m MsgRequestCallback) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Sender)}
}

// ValidateBasic implements the sdk.Msg interface.
func (m MsgRequestCallback) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return errorsmod.Wrapf(sdkErrors.ErrInvalidAddress, "invalid sender address: %v", err)
	}
	if _, err := sdk.AccAddressFromBech32(m.ContractAddress); err != nil {
		return errorsmod.Wrapf(sdkErrors.ErrInvalidAddress, "invalid contract address: %v", err)
	}

	return nil
}

// NewMsgCancelCallback creates a new MsgCancelCallback instance.
func NewMsgCancelCallback(
	senderAddr sdk.AccAddress,
	contractAddr sdk.AccAddress,
	jobId uint64,
	callbackHeight int64,
) *MsgCancelCallback {
	msg := &MsgCancelCallback{
		Sender:          senderAddr.String(),
		ContractAddress: contractAddr.String(),
		JobId:           jobId,
		CallbackHeight:  callbackHeight,
	}

	return msg
}

// GetSigners implements the sdk.Msg interface.
func (m MsgCancelCallback) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Sender)}
}

// ValidateBasic implements the sdk.Msg interface.
func (m MsgCancelCallback) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return errorsmod.Wrapf(sdkErrors.ErrInvalidAddress, "invalid sender address: %v", err)
	}
	if _, err := sdk.AccAddressFromBech32(m.ContractAddress); err != nil {
		return errorsmod.Wrapf(sdkErrors.ErrInvalidAddress, "invalid contract address: %v", err)
	}

	return nil
}

// GetSigners implements types.Msg.
func (m *MsgUpdateParams) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

// ValidateBasic implements types.Msg.
func (m *MsgUpdateParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return fmt.Errorf("invalid authority address: %w", err)
	}

	if err := m.Params.Validate(); err != nil {
		return err
	}

	return nil
}

func (m *MsgUpdateParams) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgUpdateParams) Route() string {
	return RouterKey
}

func (m *MsgUpdateParams) Type() string {
	return TypeMsgUpdateParams
}
