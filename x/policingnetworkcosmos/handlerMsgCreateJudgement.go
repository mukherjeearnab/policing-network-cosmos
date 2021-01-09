package policingnetworkcosmos

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/types"
	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/keeper"
)

func handleMsgCreateJudgement(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateJudgement) (*sdk.Result, error) {
	k.CreateJudgement(ctx, msg)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
