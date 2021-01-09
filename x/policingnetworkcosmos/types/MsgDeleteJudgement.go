package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteJudgement{}

type MsgDeleteJudgement struct {
  ID      string         `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteJudgement(id string, creator sdk.AccAddress) MsgDeleteJudgement {
  return MsgDeleteJudgement{
    ID: id,
		Creator: creator,
	}
}

func (msg MsgDeleteJudgement) Route() string {
  return RouterKey
}

func (msg MsgDeleteJudgement) Type() string {
  return "DeleteJudgement"
}

func (msg MsgDeleteJudgement) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteJudgement) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteJudgement) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}