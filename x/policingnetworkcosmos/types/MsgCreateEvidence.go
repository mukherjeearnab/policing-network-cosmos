package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateEvidence{}

type MsgCreateEvidence struct {
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  ID string `json:"ID" yaml:"ID"`
  FileExt string `json:"FileExt" yaml:"FileExt"`
  InvestigationID string `json:"InvestigationID" yaml:"InvestigationID"`
}

func NewMsgCreateEvidence(creator sdk.AccAddress, ID string, FileExt string, InvestigationID string) MsgCreateEvidence {
  return MsgCreateEvidence{
		Creator: creator,
    ID: ID,
    FileExt: FileExt,
    InvestigationID: InvestigationID,
	}
}

func (msg MsgCreateEvidence) Route() string {
  return RouterKey
}

func (msg MsgCreateEvidence) Type() string {
  return "CreateEvidence"
}

func (msg MsgCreateEvidence) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateEvidence) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateEvidence) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}