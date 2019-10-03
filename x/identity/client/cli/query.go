package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/lcnem/identity/x/identity/internal/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns query commands
func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	coinQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the idnetity module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	coinQueryCmd.AddCommand(client.GetCommands(
		getCmdGet(storeKey, cdc),
	)...)
	return coinQueryCmd
}

func getCmdGet(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get [address]",
		Short: "get address data",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryAddress), nil)
			if err != nil {
				fmt.Printf(err.Error())
				return nil
			}

			var out map[string]string
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
