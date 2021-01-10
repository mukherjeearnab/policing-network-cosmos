package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/types"
	"github.com/spf13/cobra"
)

func GetCmdListJudgement(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-judgement",
		Short: "list all judgement",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryListJudgement, queryRoute), nil)
			if err != nil {
				fmt.Printf("could not list Judgement\n%s\n", err.Error())
				return nil
			}
			var out []types.Judgement
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdGetJudgement(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-judgement [key]",
		Short: "Query a judgement by key",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			key := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryGetJudgement, key), nil)
			if err != nil {
				fmt.Printf("could not resolve judgement %s \n%s\n", key, err.Error())

				return nil
			}

			var out types.Judgement
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
