package types

import (
  sdk "github.com/cosmos/cosmos-sdk/types"
  sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetJudgement{}

type MsgSetJudgement struct {
  ID            string         `json:"id" yaml:"id"`
  Creator       sdk.AccAddress `json:"creator" yaml:"creator"`
  ChargeSheetID string         `json:"ChargeSheetID" yaml:"ChargeSheetID"`
  CourtID       string         `json:"CourtID" yaml:"CourtID"`
  Content       string         `json:"Content" yaml:"Content"`
  Complete      string         `json:"Complete" yaml:"Complete"`
}

func NewMsgSetJudgement(creator sdk.AccAddress, id string, ChargeSheetID string, CourtID string, Content string, Complete string) MsgSetJudgement {
  return MsgSetJudgement{
    ID:            id,
    Creator:       creator,
    ChargeSheetID: ChargeSheetID,
    CourtID:       CourtID,
    Content:       Content,
    Complete:      Complete,
  }
}

func (msg MsgSetJudgement) Route() string {
  return RouterKey
}

func (msg MsgSetJudgement) Type() string {
  return "SetJudgement"
}

func (msg MsgSetJudgement) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetJudgement) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetJudgement) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}
