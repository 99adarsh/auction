package auction

import (
	"auction/x/auction/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Finalize the active auctions by removing it from the ActiveAuctionList
func EndBlock(ctx sdk.Context, k keeper.Keeper) error {
	all_active_auction_list := k.GetAllActiveAuctionsList(ctx)

	for _, auction := range all_active_auction_list {
		if auction.AuctionEndHeight == ctx.BlockHeight() {
			auction_info, found := k.GetAuctionInfo(ctx, auction.AuctionId)
			if found {
				if auction_info.AuctionEndHeight == uint64(ctx.BlockHeight()) {
					k.RemoveActiveAuctionsList(ctx, auction.Id)
				}
			}
		}
	}

	return nil
}
