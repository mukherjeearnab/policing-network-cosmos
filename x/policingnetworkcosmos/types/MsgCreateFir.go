package types

import (
  sdk "github.com/cosmos/cosmos-sdk/types"
  sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateFir{}

type MsgCreateFir struct {
  Creator         sdk.AccAddress `json:"creator" yaml:"creator"`
  CitizenID       string         `json:"CitizenID" yaml:"CitizenID"`
  Content         string         `json:"Content" yaml:"Content"`
  InvestigationID string         `json:"InvestigationID" yaml:"InvestigationID"`
}

func NewMsgCreateFir(creator sdk.AccAddress, CitizenID string, Content string) MsgCreateFir {
  return MsgCreateFir{
    Creator:         creator,
    CitizenID:       CitizenID,
    Content:         Content,
    InvestigationID: "NOT SET",
  }
}

func (msg MsgCreateFir) Route() string {
  return RouterKey
}

func (msg MsgCreateFir) Type() string {
  return "CreateFir"
}

func (msg MsgCreateFir) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateFir) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateFir) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}
