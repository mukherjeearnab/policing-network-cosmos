package types

import (
  sdk "github.com/cosmos/cosmos-sdk/types"
  sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateJudgement{}

type MsgCreateJudgement struct {
  Creator       sdk.AccAddress `json:"creator" yaml:"creator"`
  ChargeSheetID string         `json:"ChargeSheetID" yaml:"ChargeSheetID"`
  CourtID       string         `json:"CourtID" yaml:"CourtID"`
  Content       string         `json:"Content" yaml:"Content"`
}

func NewMsgCreateJudgement(creator sdk.AccAddress, ChargeSheetID string, CourtID string, Content string) MsgCreateJudgement {
  return MsgCreateJudgement{
    Creator:       creator,
    ChargeSheetID: ChargeSheetID,
    CourtID:       CourtID,
    Content:       Content,
  }
}

func (msg MsgCreateJudgement) Route() string {
  return RouterKey
}

func (msg MsgCreateJudgement) Type() string {
  return "CreateJudgement"
}

func (msg MsgCreateJudgement) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateJudgement) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateJudgement) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}
