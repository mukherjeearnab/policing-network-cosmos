package policingnetworkcosmos

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/keeper"
	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/types"
)

func handleMsgSetChargesheet(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetChargesheet) (*sdk.Result, error) {
	var chargesheet = types.Chargesheet{
		Creator:          msg.Creator,
		ID:               msg.ID,
		OfficerIDs:       msg.OfficerIDs,
		FirIDs:           msg.FirIDs,
		InvestigationIDs: msg.InvestigationIDs,
		Content:          msg.Content,
		Complete:         msg.Complete,
	}
	if !msg.Creator.Equals(k.GetChargesheetOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetChargesheet(ctx, chargesheet)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
