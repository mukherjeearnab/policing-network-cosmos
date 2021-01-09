package types

import (
  sdk "github.com/cosmos/cosmos-sdk/types"
  sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteInvestigation{}

type MsgDeleteInvestigation struct {
  ID      string         `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteInvestigation(id string, creator sdk.AccAddress) MsgDeleteInvestigation {
  return MsgDeleteInvestigation{
    ID:      id,
    Creator: creator,
  }
}

func (msg MsgDeleteInvestigation) Route() string {
  return RouterKey
}

func (msg MsgDeleteInvestigation) Type() string {
  return "DeleteInvestigation"
}

func (msg MsgDeleteInvestigation) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteInvestigation) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteInvestigation) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}
