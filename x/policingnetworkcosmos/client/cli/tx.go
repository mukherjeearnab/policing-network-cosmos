package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	policingnetworkcosmosTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	policingnetworkcosmosTxCmd.AddCommand(flags.PostCommands(
    // this line is used by starport scaffolding # 1
		GetCmdCreateJudgement(cdc),
		GetCmdSetJudgement(cdc),
		GetCmdDeleteJudgement(cdc),
		GetCmdCreateChargesheet(cdc),
		GetCmdSetChargesheet(cdc),
		GetCmdDeleteChargesheet(cdc),
		GetCmdCreateEvidence(cdc),
		GetCmdSetEvidence(cdc),
		GetCmdDeleteEvidence(cdc),
		GetCmdCreateInvestigation(cdc),
		GetCmdSetInvestigation(cdc),
		GetCmdDeleteInvestigation(cdc),
		GetCmdCreateFir(cdc),
		GetCmdSetFir(cdc),
		GetCmdDeleteFir(cdc),
		GetCmdCreateProfile(cdc),
		GetCmdSetProfile(cdc),
		GetCmdDeleteProfile(cdc),
	)...)

	return policingnetworkcosmosTxCmd
}
