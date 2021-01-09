package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteFir{}

type MsgDeleteFir struct {
  ID      string         `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteFir(id string, creator sdk.AccAddress) MsgDeleteFir {
  return MsgDeleteFir{
    ID: id,
		Creator: creator,
	}
}

func (msg MsgDeleteFir) Route() string {
  return RouterKey
}

func (msg MsgDeleteFir) Type() string {
  return "DeleteFir"
}

func (msg MsgDeleteFir) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteFir) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteFir) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}