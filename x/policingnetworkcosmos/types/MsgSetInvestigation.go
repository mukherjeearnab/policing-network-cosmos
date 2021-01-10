package types

import (
  sdk "github.com/cosmos/cosmos-sdk/types"
  sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetInvestigation{}

type MsgSetInvestigation struct {
  ID        string         `json:"id" yaml:"id"`
  Creator   sdk.AccAddress `json:"creator" yaml:"creator"`
  FirID     string         `json:"FirID" yaml:"FirID"`
  OfficerID string         `json:"OfficerID" yaml:"OfficerID"`
  Content   string         `json:"Content" yaml:"Content"`
  Complete  string         `json:"Complete" yaml:"Complete"`
}

func NewMsgSetInvestigation(creator sdk.AccAddress, id string, FirID string, OfficerID string, Content string, Complete string) MsgSetInvestigation {
  return MsgSetInvestigation{
    ID:        id,
    Creator:   creator,
    FirID:     FirID,
    OfficerID: OfficerID,
    Content:   Content,
    Complete:  Complete,
  }
}

func (msg MsgSetInvestigation) Route() string {
  return RouterKey
}

func (msg MsgSetInvestigation) Type() string {
  return "SetInvestigation"
}

func (msg MsgSetInvestigation) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetInvestigation) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetInvestigation) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}
