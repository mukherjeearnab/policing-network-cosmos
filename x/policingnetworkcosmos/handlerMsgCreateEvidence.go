package policingnetworkcosmos

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/keeper"
	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/types"
)

func handleMsgCreateEvidence(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateEvidence) (*sdk.Result, error) {
	k.CreateEvidence(ctx, msg)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
