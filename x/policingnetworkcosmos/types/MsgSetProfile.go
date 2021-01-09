package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetProfile{}

type MsgSetProfile struct {
  ID      string      `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Type string `json:"Type" yaml:"Type"`
  ID string `json:"ID" yaml:"ID"`
  Name string `json:"Name" yaml:"Name"`
  Role string `json:"Role" yaml:"Role"`
  FirList string `json:"FirList" yaml:"FirList"`
}

func NewMsgSetProfile(creator sdk.AccAddress, id string, Type string, ID string, Name string, Role string, FirList string) MsgSetProfile {
  return MsgSetProfile{
    ID: id,
		Creator: creator,
    Type: Type,
    ID: ID,
    Name: Name,
    Role: Role,
    FirList: FirList,
	}
}

func (msg MsgSetProfile) Route() string {
  return RouterKey
}

func (msg MsgSetProfile) Type() string {
  return "SetProfile"
}

func (msg MsgSetProfile) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetProfile) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetProfile) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}