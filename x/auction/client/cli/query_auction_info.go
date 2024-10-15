package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"auction/x/auction/types"
)

func CmdListAuctionInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-auction-info",
		Short: "list all auction-info",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllAuctionInfoRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.AuctionInfoAll(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowAuctionInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-auction-info [auction-id]",
		Short: "shows a auction-info",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argAuctionId := args[0]

			params := &types.QueryGetAuctionInfoRequest{
				AuctionId: argAuctionId,
			}

			res, err := queryClient.AuctionInfo(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
