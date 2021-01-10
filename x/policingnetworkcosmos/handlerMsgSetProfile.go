package policingnetworkcosmos

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/keeper"
	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/types"
)

func handleMsgSetProfile(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetProfile) (*sdk.Result, error) {
	address, _ := sdk.AccAddressFromBech32(string(msg.Address))

	var profile = types.Profile{
		Creator:     msg.Creator,
		ID:          msg.ID,
		ProfileType: msg.ProfileType,
		Address:     address,
		Name:        msg.Name,
		Role:        msg.Role,
		FirList:     msg.FirList,
	}
	if !msg.Creator.Equals(k.GetProfileOwner(ctx, msg.Address)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetProfile(ctx, profile)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
