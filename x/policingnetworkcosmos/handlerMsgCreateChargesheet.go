package policingnetworkcosmos

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/types"
	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/keeper"
)

func handleMsgCreateChargesheet(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateChargesheet) (*sdk.Result, error) {
	k.CreateChargesheet(ctx, msg)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
