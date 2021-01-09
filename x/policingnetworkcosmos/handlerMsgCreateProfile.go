package policingnetworkcosmos

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/types"
	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/keeper"
)

func handleMsgCreateProfile(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateProfile) (*sdk.Result, error) {
	k.CreateProfile(ctx, msg)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
