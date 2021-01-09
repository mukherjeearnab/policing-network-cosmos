package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Chargesheet struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
    ID string `json:"ID" yaml:"ID"`
    OfficerIDs string `json:"OfficerIDs" yaml:"OfficerIDs"`
    FirIDs string `json:"FirIDs" yaml:"FirIDs"`
    InvestigationIDs string `json:"InvestigationIDs" yaml:"InvestigationIDs"`
    Content string `json:"Content" yaml:"Content"`
    Complete string `json:"Complete" yaml:"Complete"`
}