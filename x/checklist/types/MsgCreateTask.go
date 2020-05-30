package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgCreateTask
var _ sdk.Msg = &MsgCreateTask{}

// MsgCreateTask ...
type MsgCreateTask struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Title   string         `json:"title" yaml:"title"`
}

// NewMsgCreateTask ...
func NewMsgCreateTask(creator sdk.AccAddress, title string) MsgCreateTask {
	return MsgCreateTask{
		Creator: creator,
		Title:   title,
	}
}

// Route ...
func (msg MsgCreateTask) Route() string { return RouterKey }

// Type ...
func (msg MsgCreateTask) Type() string { return "CreateTask" }

// GetSigners ...
func (msg MsgCreateTask) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

// GetSignBytes ...
func (msg MsgCreateTask) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic ...
func (msg MsgCreateTask) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
