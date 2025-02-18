package types

import (
  sdk "github.com/cosmos/cosmos-sdk/types"
  sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteProfile{}

type MsgDeleteProfile struct {
  Address string         `json:"Address" yaml:"Address"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteProfile(Address string, creator sdk.AccAddress) MsgDeleteProfile {
  return MsgDeleteProfile{
    Address: Address,
    Creator: creator,
  }
}

func (msg MsgDeleteProfile) Route() string {
  return RouterKey
}

func (msg MsgDeleteProfile) Type() string {
  return "DeleteProfile"
}

func (msg MsgDeleteProfile) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteProfile) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteProfile) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}
