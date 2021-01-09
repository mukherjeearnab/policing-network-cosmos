package policingnetworkcosmos

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/types"
	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/keeper"
)

// Handle a message to delete name
func handleMsgDeleteEvidence(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteEvidence) (*sdk.Result, error) {
	if !k.EvidenceExists(ctx, msg.ID) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.ID)
	}
	if !msg.Creator.Equals(k.GetEvidenceOwner(ctx, msg.ID)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeleteEvidence(ctx, msg.ID)
	return &sdk.Result{}, nil
}
