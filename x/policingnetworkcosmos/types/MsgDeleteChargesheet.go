package types

import (
  sdk "github.com/cosmos/cosmos-sdk/types"
  sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteChargesheet{}

type MsgDeleteChargesheet struct {
  ID      string         `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteChargesheet(id string, creator sdk.AccAddress) MsgDeleteChargesheet {
  return MsgDeleteChargesheet{
    ID:      id,
    Creator: creator,
  }
}

func (msg MsgDeleteChargesheet) Route() string {
  return RouterKey
}

func (msg MsgDeleteChargesheet) Type() string {
  return "DeleteChargesheet"
}

func (msg MsgDeleteChargesheet) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteChargesheet) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteChargesheet) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}
