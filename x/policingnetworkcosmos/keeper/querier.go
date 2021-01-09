package keeper

import (
  // this line is used by starport scaffolding # 1
	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/types"
		
	
		
	
		
	
		
	
		
	
		
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for policingnetworkcosmos clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
    // this line is used by starport scaffolding # 2
		case types.QueryListJudgement:
			return listJudgement(ctx, k)
		case types.QueryGetJudgement:
			return getJudgement(ctx, path[1:], k)
		case types.QueryListChargesheet:
			return listChargesheet(ctx, k)
		case types.QueryGetChargesheet:
			return getChargesheet(ctx, path[1:], k)
		case types.QueryListEvidence:
			return listEvidence(ctx, k)
		case types.QueryGetEvidence:
			return getEvidence(ctx, path[1:], k)
		case types.QueryListInvestigation:
			return listInvestigation(ctx, k)
		case types.QueryGetInvestigation:
			return getInvestigation(ctx, path[1:], k)
		case types.QueryListFir:
			return listFir(ctx, k)
		case types.QueryGetFir:
			return getFir(ctx, path[1:], k)
		case types.QueryListProfile:
			return listProfile(ctx, k)
		case types.QueryGetProfile:
			return getProfile(ctx, path[1:], k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown policingnetworkcosmos query endpoint")
		}
	}
}
