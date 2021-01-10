package types

import (
  sdk "github.com/cosmos/cosmos-sdk/types"
  sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateProfile{}

type MsgCreateProfile struct {
  Creator     sdk.AccAddress `json:"creator" yaml:"creator"`
  ProfileType string         `json:"ProfileType" yaml:"ProfileType"`
  ID          string         `json:"ID" yaml:"ID"`
  Address     string         `json:"Address" yaml:"Address"`
  Name        string         `json:"Name" yaml:"Name"`
  Role        string         `json:"Role" yaml:"Role"`
}

func NewMsgCreateProfile(creator sdk.AccAddress, Address string, ProfileType string, ID string, Name string, Role string) MsgCreateProfile {
  return MsgCreateProfile{
    Creator:     creator,
    ProfileType: ProfileType,
    ID:          ID,
    Address:     Address,
    Name:        Name,
    Role:        Role,
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

  _, err := sdk.AccAddressFromBech32(msg.Address)

  if err != nil {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Invalid Proile Address")
  }
  return nil
}
