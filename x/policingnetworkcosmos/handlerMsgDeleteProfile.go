package policingnetworkcosmos

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/keeper"
	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/types"
)

// Handle a message to delete name
func handleMsgDeleteProfile(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteProfile) (*sdk.Result, error) {
	if !k.ProfileExists(ctx, msg.Address) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.Address)
	}
	if !msg.Creator.Equals(k.GetProfileOwner(ctx, msg.Address)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeleteProfile(ctx, msg.Address)
	return &sdk.Result{}, nil
}
