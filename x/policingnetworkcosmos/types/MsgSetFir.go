package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetFir{}

type MsgSetFir struct {
  ID      string      `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  ID string `json:"ID" yaml:"ID"`
  CitizenID string `json:"CitizenID" yaml:"CitizenID"`
  Content string `json:"Content" yaml:"Content"`
  InvestigationID string `json:"InvestigationID" yaml:"InvestigationID"`
}

func NewMsgSetFir(creator sdk.AccAddress, id string, ID string, CitizenID string, Content string, InvestigationID string) MsgSetFir {
  return MsgSetFir{
    ID: id,
		Creator: creator,
    ID: ID,
    CitizenID: CitizenID,
    Content: Content,
    InvestigationID: InvestigationID,
	}
}

func (msg MsgSetFir) Route() string {
  return RouterKey
}

func (msg MsgSetFir) Type() string {
  return "SetFir"
}

func (msg MsgSetFir) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetFir) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetFir) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}