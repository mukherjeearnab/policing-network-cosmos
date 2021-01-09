package policingnetworkcosmos

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/keeper"
	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
    // this line is used by starport scaffolding # 1
		case types.MsgCreateJudgement:
			return handleMsgCreateJudgement(ctx, k, msg)
		case types.MsgSetJudgement:
			return handleMsgSetJudgement(ctx, k, msg)
		case types.MsgDeleteJudgement:
			return handleMsgDeleteJudgement(ctx, k, msg)
		case types.MsgCreateChargesheet:
			return handleMsgCreateChargesheet(ctx, k, msg)
		case types.MsgSetChargesheet:
			return handleMsgSetChargesheet(ctx, k, msg)
		case types.MsgDeleteChargesheet:
			return handleMsgDeleteChargesheet(ctx, k, msg)
		case types.MsgCreateEvidence:
			return handleMsgCreateEvidence(ctx, k, msg)
		case types.MsgSetEvidence:
			return handleMsgSetEvidence(ctx, k, msg)
		case types.MsgDeleteEvidence:
			return handleMsgDeleteEvidence(ctx, k, msg)
		case types.MsgCreateInvestigation:
			return handleMsgCreateInvestigation(ctx, k, msg)
		case types.MsgSetInvestigation:
			return handleMsgSetInvestigation(ctx, k, msg)
		case types.MsgDeleteInvestigation:
			return handleMsgDeleteInvestigation(ctx, k, msg)
		case types.MsgCreateFir:
			return handleMsgCreateFir(ctx, k, msg)
		case types.MsgSetFir:
			return handleMsgSetFir(ctx, k, msg)
		case types.MsgDeleteFir:
			return handleMsgDeleteFir(ctx, k, msg)
		case types.MsgCreateProfile:
			return handleMsgCreateProfile(ctx, k, msg)
		case types.MsgSetProfile:
			return handleMsgSetProfile(ctx, k, msg)
		case types.MsgDeleteProfile:
			return handleMsgDeleteProfile(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
