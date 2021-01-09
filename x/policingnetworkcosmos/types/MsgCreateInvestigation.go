package types

import (
  sdk "github.com/cosmos/cosmos-sdk/types"
  sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateInvestigation{}

type MsgCreateInvestigation struct {
  Creator   sdk.AccAddress `json:"creator" yaml:"creator"`
  FirID     string         `json:"FirID" yaml:"FirID"`
  OfficerID string         `json:"OfficerID" yaml:"OfficerID"`
  Content   string         `json:"Content" yaml:"Content"`
}

func NewMsgCreateInvestigation(creator sdk.AccAddress, FirID string, OfficerID string, Content string) MsgCreateInvestigation {
  return MsgCreateInvestigation{
    Creator:   creator,
    FirID:     FirID,
    OfficerID: OfficerID,
    Content:   Content,
  }
}

func (msg MsgCreateInvestigation) Route() string {
  return RouterKey
}

func (msg MsgCreateInvestigation) Type() string {
  return "CreateInvestigation"
}

func (msg MsgCreateInvestigation) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateInvestigation) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateInvestigation) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}
