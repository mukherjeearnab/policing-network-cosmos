package policingnetworkcosmos

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/types"
	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/keeper"
)

func handleMsgSetJudgement(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetJudgement) (*sdk.Result, error) {
	var judgement = types.Judgement{
		Creator: msg.Creator,
		ID:      msg.ID,
    	ID: msg.ID,
    	ChargeSheetID: msg.ChargeSheetID,
    	CourtID: msg.CourtID,
    	Content: msg.Content,
    	Complete: msg.Complete,
	}
	if !msg.Creator.Equals(k.GetJudgementOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetJudgement(ctx, judgement)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
