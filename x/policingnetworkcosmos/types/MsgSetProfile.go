package types

import (
  sdk "github.com/cosmos/cosmos-sdk/types"
  sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetProfile{}

type MsgSetProfile struct {
  ID          string         `json:"id" yaml:"id"`
  Creator     sdk.AccAddress `json:"creator" yaml:"creator"`
  ProfileType string         `json:"ProfileType" yaml:"ProfileType"`
  Address     string         `json:"Address" yaml:"Address"`
  Name        string         `json:"Name" yaml:"Name"`
  Role        string         `json:"Role" yaml:"Role"`
}

func NewMsgSetProfile(creator sdk.AccAddress, Address string, id string, ProfileType string, Name string, Role string) MsgSetProfile {
  return MsgSetProfile{
    ID:          id,
    Creator:     creator,
    ProfileType: ProfileType,
    Address:     Address,
    Name:        Name,
    Role:        Role,
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

  _, err := sdk.AccAddressFromBech32(msg.Address)

  if err != nil {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Invalid Proile Address")
  }
  return nil
}
