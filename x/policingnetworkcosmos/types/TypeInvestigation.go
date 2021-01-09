package types

import (
    sdk "github.com/cosmos/cosmos-sdk/types"
)

type Investigation struct {
    Creator   sdk.AccAddress `json:"creator" yaml:"creator"`
    ID        string         `json:"id" yaml:"id"`
    FirID     string         `json:"FirID" yaml:"FirID"`
    OfficerID string         `json:"OfficerID" yaml:"OfficerID"`
    Content   string         `json:"Content" yaml:"Content"`
    Evidence  []string       `json:"Evidence" yaml:"Evidence"`
    Complete  bool           `json:"Complete" yaml:"Complete"`
}
