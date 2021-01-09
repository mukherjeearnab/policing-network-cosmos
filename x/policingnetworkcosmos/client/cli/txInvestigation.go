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

func GetCmdCreateInvestigation(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-investigation [ID] [FirID] [OfficerID] [Content] [Evidence] [Complete]",
		Short: "Creates a new investigation",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsID := string(args[0] )
			argsFirID := string(args[1] )
			argsOfficerID := string(args[2] )
			argsContent := string(args[3] )
			argsEvidence := string(args[4] )
			argsComplete := string(args[5] )
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateInvestigation(cliCtx.GetFromAddress(), string(argsID), string(argsFirID), string(argsOfficerID), string(argsContent), string(argsEvidence), string(argsComplete))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}


func GetCmdSetInvestigation(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-investigation [id]  [ID] [FirID] [OfficerID] [Content] [Evidence] [Complete]",
		Short: "Set a new investigation",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsID := string(args[1])
			argsFirID := string(args[2])
			argsOfficerID := string(args[3])
			argsContent := string(args[4])
			argsEvidence := string(args[5])
			argsComplete := string(args[6])
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgSetInvestigation(cliCtx.GetFromAddress(), id, string(argsID), string(argsFirID), string(argsOfficerID), string(argsContent), string(argsEvidence), string(argsComplete))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdDeleteInvestigation(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-investigation [id]",
		Short: "Delete a new investigation by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteInvestigation(args[0], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
