package types

import (
  sdk "github.com/cosmos/cosmos-sdk/types"
  sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateChargesheet{}

type MsgCreateChargesheet struct {
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Content string         `json:"Content" yaml:"Content"`
}

func NewMsgCreateChargesheet(creator sdk.AccAddress, Content string) MsgCreateChargesheet {
  return MsgCreateChargesheet{
    Creator: creator,
    Content: Content,
  }
}

func (msg MsgCreateChargesheet) Route() string {
  return RouterKey
}

func (msg MsgCreateChargesheet) Type() string {
  return "CreateChargesheet"
}

func (msg MsgCreateChargesheet) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateChargesheet) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateChargesheet) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}
