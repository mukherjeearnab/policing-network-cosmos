package policingnetworkcosmos

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/types"
	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/keeper"
)

func handleMsgCreateFir(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateFir) (*sdk.Result, error) {
	k.CreateFir(ctx, msg)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
