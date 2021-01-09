package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetChargesheet{}

type MsgSetChargesheet struct {
  ID      string      `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  ID string `json:"ID" yaml:"ID"`
  OfficerIDs string `json:"OfficerIDs" yaml:"OfficerIDs"`
  FirIDs string `json:"FirIDs" yaml:"FirIDs"`
  InvestigationIDs string `json:"InvestigationIDs" yaml:"InvestigationIDs"`
  Content string `json:"Content" yaml:"Content"`
  Complete string `json:"Complete" yaml:"Complete"`
}

func NewMsgSetChargesheet(creator sdk.AccAddress, id string, ID string, OfficerIDs string, FirIDs string, InvestigationIDs string, Content string, Complete string) MsgSetChargesheet {
  return MsgSetChargesheet{
    ID: id,
		Creator: creator,
    ID: ID,
    OfficerIDs: OfficerIDs,
    FirIDs: FirIDs,
    InvestigationIDs: InvestigationIDs,
    Content: Content,
    Complete: Complete,
	}
}

func (msg MsgSetChargesheet) Route() string {
  return RouterKey
}

func (msg MsgSetChargesheet) Type() string {
  return "SetChargesheet"
}

func (msg MsgSetChargesheet) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetChargesheet) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetChargesheet) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}