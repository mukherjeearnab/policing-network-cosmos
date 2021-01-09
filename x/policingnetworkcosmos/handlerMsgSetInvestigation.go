package policingnetworkcosmos

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/types"
	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/keeper"
)

func handleMsgSetInvestigation(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetInvestigation) (*sdk.Result, error) {
	var investigation = types.Investigation{
		Creator: msg.Creator,
		ID:      msg.ID,
    	ID: msg.ID,
    	FirID: msg.FirID,
    	OfficerID: msg.OfficerID,
    	Content: msg.Content,
    	Evidence: msg.Evidence,
    	Complete: msg.Complete,
	}
	if !msg.Creator.Equals(k.GetInvestigationOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetInvestigation(ctx, investigation)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
