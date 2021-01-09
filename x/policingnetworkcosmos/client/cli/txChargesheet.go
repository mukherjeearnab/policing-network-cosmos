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

func GetCmdCreateChargesheet(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-chargesheet [ID] [OfficerIDs] [FirIDs] [InvestigationIDs] [Content] [Complete]",
		Short: "Creates a new chargesheet",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsID := string(args[0] )
			argsOfficerIDs := string(args[1] )
			argsFirIDs := string(args[2] )
			argsInvestigationIDs := string(args[3] )
			argsContent := string(args[4] )
			argsComplete := string(args[5] )
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateChargesheet(cliCtx.GetFromAddress(), string(argsID), string(argsOfficerIDs), string(argsFirIDs), string(argsInvestigationIDs), string(argsContent), string(argsComplete))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}


func GetCmdSetChargesheet(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-chargesheet [id]  [ID] [OfficerIDs] [FirIDs] [InvestigationIDs] [Content] [Complete]",
		Short: "Set a new chargesheet",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsID := string(args[1])
			argsOfficerIDs := string(args[2])
			argsFirIDs := string(args[3])
			argsInvestigationIDs := string(args[4])
			argsContent := string(args[5])
			argsComplete := string(args[6])
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgSetChargesheet(cliCtx.GetFromAddress(), id, string(argsID), string(argsOfficerIDs), string(argsFirIDs), string(argsInvestigationIDs), string(argsContent), string(argsComplete))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdDeleteChargesheet(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-chargesheet [id]",
		Short: "Delete a new chargesheet by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteChargesheet(args[0], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
