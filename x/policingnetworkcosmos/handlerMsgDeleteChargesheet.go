package policingnetworkcosmos

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/keeper"
	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/types"
)

// Handle a message to delete name
func handleMsgDeleteChargesheet(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteChargesheet) (*sdk.Result, error) {
	if !k.ChargesheetExists(ctx, msg.ID) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.ID)
	}
	if !msg.Creator.Equals(k.GetChargesheetOwner(ctx, msg.ID)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeleteChargesheet(ctx, msg.ID)
	return &sdk.Result{}, nil
}
