package types

import (
  sdk "github.com/cosmos/cosmos-sdk/types"
  sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetEvidence{}

type MsgSetEvidence struct {
  ID              string         `json:"id" yaml:"id"`
  Creator         sdk.AccAddress `json:"creator" yaml:"creator"`
  FileExt         string         `json:"FileExt" yaml:"FileExt"`
  InvestigationID string         `json:"InvestigationID" yaml:"InvestigationID"`
}

func NewMsgSetEvidence(creator sdk.AccAddress, id string, FileExt string, InvestigationID string) MsgSetEvidence {
  return MsgSetEvidence{
    ID:              id,
    Creator:         creator,
    FileExt:         FileExt,
    InvestigationID: InvestigationID,
  }
}

func (msg MsgSetEvidence) Route() string {
  return RouterKey
}

func (msg MsgSetEvidence) Type() string {
  return "SetEvidence"
}

func (msg MsgSetEvidence) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetEvidence) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetEvidence) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}
