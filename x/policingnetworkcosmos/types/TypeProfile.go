package types

import (
    sdk "github.com/cosmos/cosmos-sdk/types"
)

type Profile struct {
    Creator sdk.AccAddress `json:"creator" yaml:"creator"`
    ID      string         `json:"id" yaml:"id"`
    Type    string         `json:"Type" yaml:"Type"`
    ID      string         `json:"ID" yaml:"ID"`
    Name    string         `json:"Name" yaml:"Name"`
    Role    string         `json:"Role" yaml:"Role"`
    FirList string         `json:"FirList" yaml:"FirList"`
}
