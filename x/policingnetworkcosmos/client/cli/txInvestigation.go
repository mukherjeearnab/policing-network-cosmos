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
		Use:   "create-investigation [FirID] [OfficerID] [Content]",
		Short: "Creates a new investigation",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsFirID := string(args[0])
			argsOfficerID := string(args[1])
			argsContent := string(args[2])

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateInvestigation(cliCtx.GetFromAddress(), string(argsFirID), string(argsOfficerID), string(argsContent))
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
		Use:   "set-investigation [id] [FirID] [OfficerID] [Content] [Complete]",
		Short: "Set a new investigation",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsFirID := string(args[1])
			argsOfficerID := string(args[2])
			argsContent := string(args[3])
			argsComplete := string(args[4])

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgSetInvestigation(cliCtx.GetFromAddress(), id, string(argsFirID), string(argsOfficerID), string(argsContent), string(argsComplete))
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
