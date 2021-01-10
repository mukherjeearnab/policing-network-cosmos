package types

import (
    sdk "github.com/cosmos/cosmos-sdk/types"
)

type Profile struct {
    Creator     sdk.AccAddress `json:"creator" yaml:"creator"`
    ID          string         `json:"id" yaml:"id"`
    Address     sdk.AccAddress `json:"Address" yaml:"Address"`
    ProfileType string         `json:"ProfileType" yaml:"ProfileType"`
    Name        string         `json:"Name" yaml:"Name"`
    Role        string         `json:"Role" yaml:"Role"`
    FirList     []string       `json:"FirList" yaml:"FirList"`
}
