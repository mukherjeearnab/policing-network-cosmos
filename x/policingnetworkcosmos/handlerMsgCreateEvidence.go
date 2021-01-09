package policingnetworkcosmos

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/types"
	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/keeper"
)

func handleMsgCreateEvidence(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateEvidence) (*sdk.Result, error) {
	k.CreateEvidence(ctx, msg)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
