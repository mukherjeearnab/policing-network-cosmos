package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Judgement struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
    ID string `json:"ID" yaml:"ID"`
    ChargeSheetID string `json:"ChargeSheetID" yaml:"ChargeSheetID"`
    CourtID string `json:"CourtID" yaml:"CourtID"`
    Content string `json:"Content" yaml:"Content"`
    Complete string `json:"Complete" yaml:"Complete"`
}