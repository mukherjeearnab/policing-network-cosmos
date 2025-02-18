package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	// "github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group policingnetworkcosmos queries under a subcommand
	policingnetworkcosmosQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	policingnetworkcosmosQueryCmd.AddCommand(
		flags.GetCommands(
      // this line is used by starport scaffolding # 1
			GetCmdListJudgement(queryRoute, cdc),
			GetCmdGetJudgement(queryRoute, cdc),
			GetCmdListChargesheet(queryRoute, cdc),
			GetCmdGetChargesheet(queryRoute, cdc),
			GetCmdListEvidence(queryRoute, cdc),
			GetCmdGetEvidence(queryRoute, cdc),
			GetCmdListInvestigation(queryRoute, cdc),
			GetCmdGetInvestigation(queryRoute, cdc),
			GetCmdListFir(queryRoute, cdc),
			GetCmdGetFir(queryRoute, cdc),
			GetCmdListProfile(queryRoute, cdc),
			GetCmdGetProfile(queryRoute, cdc),
		)...,
	)

	return policingnetworkcosmosQueryCmd
}
