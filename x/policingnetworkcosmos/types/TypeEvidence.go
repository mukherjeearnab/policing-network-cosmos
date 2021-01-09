package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Evidence struct {
	Creator         sdk.AccAddress `json:"creator" yaml:"creator"`
	ID              string         `json:"id" yaml:"id"`
	FileExt         string         `json:"FileExt" yaml:"FileExt"`
	InvestigationID string         `json:"InvestigationID" yaml:"InvestigationID"`
}
