package policingnetworkcosmos

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/types"
	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/keeper"
)

func handleMsgSetEvidence(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetEvidence) (*sdk.Result, error) {
	var evidence = types.Evidence{
		Creator: msg.Creator,
		ID:      msg.ID,
    	ID: msg.ID,
    	FileExt: msg.FileExt,
    	InvestigationID: msg.InvestigationID,
	}
	if !msg.Creator.Equals(k.GetEvidenceOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetEvidence(ctx, evidence)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
