package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateProfile{}

type MsgCreateProfile struct {
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Type string `json:"Type" yaml:"Type"`
  ID string `json:"ID" yaml:"ID"`
  Name string `json:"Name" yaml:"Name"`
  Role string `json:"Role" yaml:"Role"`
  FirList string `json:"FirList" yaml:"FirList"`
}

func NewMsgCreateProfile(creator sdk.AccAddress, Type string, ID string, Name string, Role string, FirList string) MsgCreateProfile {
  return MsgCreateProfile{
		Creator: creator,
    Type: Type,
    ID: ID,
    Name: Name,
    Role: Role,
    FirList: FirList,
	}
}

func (msg MsgCreateProfile) Route() string {
  return RouterKey
}

func (msg MsgCreateProfile) Type() string {
  return "CreateProfile"
}

func (msg MsgCreateProfile) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateProfile) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateProfile) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}