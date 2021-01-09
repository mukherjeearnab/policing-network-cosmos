package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteEvidence{}

type MsgDeleteEvidence struct {
  ID      string         `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteEvidence(id string, creator sdk.AccAddress) MsgDeleteEvidence {
  return MsgDeleteEvidence{
    ID: id,
		Creator: creator,
	}
}

func (msg MsgDeleteEvidence) Route() string {
  return RouterKey
}

func (msg MsgDeleteEvidence) Type() string {
  return "DeleteEvidence"
}

func (msg MsgDeleteEvidence) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteEvidence) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteEvidence) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}