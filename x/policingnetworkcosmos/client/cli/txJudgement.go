package cli

import (
	"bufio"
    
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/types"
)

func GetCmdCreateJudgement(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-judgement [ID] [ChargeSheetID] [CourtID] [Content] [Complete]",
		Short: "Creates a new judgement",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsID := string(args[0] )
			argsChargeSheetID := string(args[1] )
			argsCourtID := string(args[2] )
			argsContent := string(args[3] )
			argsComplete := string(args[4] )
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateJudgement(cliCtx.GetFromAddress(), string(argsID), string(argsChargeSheetID), string(argsCourtID), string(argsContent), string(argsComplete))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}


func GetCmdSetJudgement(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-judgement [id]  [ID] [ChargeSheetID] [CourtID] [Content] [Complete]",
		Short: "Set a new judgement",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsID := string(args[1])
			argsChargeSheetID := string(args[2])
			argsCourtID := string(args[3])
			argsContent := string(args[4])
			argsComplete := string(args[5])
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgSetJudgement(cliCtx.GetFromAddress(), id, string(argsID), string(argsChargeSheetID), string(argsCourtID), string(argsContent), string(argsComplete))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdDeleteJudgement(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-judgement [id]",
		Short: "Delete a new judgement by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteJudgement(args[0], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
