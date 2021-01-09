package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Fir struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
    ID string `json:"ID" yaml:"ID"`
    CitizenID string `json:"CitizenID" yaml:"CitizenID"`
    Content string `json:"Content" yaml:"Content"`
    InvestigationID string `json:"InvestigationID" yaml:"InvestigationID"`
}