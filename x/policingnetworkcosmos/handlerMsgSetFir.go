package policingnetworkcosmos

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/keeper"
	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/types"
)

func handleMsgSetFir(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetFir) (*sdk.Result, error) {
	var fir = types.Fir{
		Creator:         msg.Creator,
		ID:              msg.ID,
		CitizenID:       msg.CitizenID,
		Content:         msg.Content,
		InvestigationID: msg.InvestigationID,
	}
	if !msg.Creator.Equals(k.GetFirOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetFir(ctx, fir)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
