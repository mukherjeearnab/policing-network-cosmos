package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/types"
	"github.com/spf13/cobra"
)

func GetCmdListInvestigation(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-investigation",
		Short: "list all investigation",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryListInvestigation, queryRoute), nil)
			if err != nil {
				fmt.Printf("could not list Investigation\n%s\n", err.Error())
				return nil
			}
			var out []types.Investigation
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdGetInvestigation(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-investigation [key]",
		Short: "Query a investigation by key",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			key := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryGetInvestigation, key), nil)
			if err != nil {
				fmt.Printf("could not resolve investigation %s \n%s\n", key, err.Error())

				return nil
			}

			var out types.Investigation
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
