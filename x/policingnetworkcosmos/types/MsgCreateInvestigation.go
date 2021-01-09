package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateInvestigation{}

type MsgCreateInvestigation struct {
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  ID string `json:"ID" yaml:"ID"`
  FirID string `json:"FirID" yaml:"FirID"`
  OfficerID string `json:"OfficerID" yaml:"OfficerID"`
  Content string `json:"Content" yaml:"Content"`
  Evidence string `json:"Evidence" yaml:"Evidence"`
  Complete string `json:"Complete" yaml:"Complete"`
}

func NewMsgCreateInvestigation(creator sdk.AccAddress, ID string, FirID string, OfficerID string, Content string, Evidence string, Complete string) MsgCreateInvestigation {
  return MsgCreateInvestigation{
		Creator: creator,
    ID: ID,
    FirID: FirID,
    OfficerID: OfficerID,
    Content: Content,
    Evidence: Evidence,
    Complete: Complete,
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