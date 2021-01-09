package types

import (
  sdk "github.com/cosmos/cosmos-sdk/types"
  sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetChargesheet{}

type MsgSetChargesheet struct {
  ID       string         `json:"id" yaml:"id"`
  Creator  sdk.AccAddress `json:"creator" yaml:"creator"`
  Content  string         `json:"Content" yaml:"Content"`
  Complete bool           `json:"Complete" yaml:"Complete"`
}

func NewMsgSetChargesheet(creator sdk.AccAddress, id string, Content string, Complete string) MsgSetChargesheet {
  return MsgSetChargesheet{
    ID:       id,
    Creator:  creator,
    Content:  Content,
    Complete: true,
  }
}

func (msg MsgSetChargesheet) Route() string {
  return RouterKey
}

func (msg MsgSetChargesheet) Type() string {
  return "SetChargesheet"
}

func (msg MsgSetChargesheet) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetChargesheet) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetChargesheet) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}
