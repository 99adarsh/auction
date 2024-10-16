package keeper

import (
	"context"

	"auction/x/auction/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) PlaceBid(goCtx context.Context, msg *types.MsgPlaceBid) (*types.MsgPlaceBidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	auction_info, found := k.Keeper.GetAuctionInfo(ctx, msg.AuctionId)

	// Auction validity check
	if !found {
		return nil, types.ErrAuctionInfoNotFound
	}
	if ctx.BlockHeight() >= int64(auction_info.AuctionEndHeight) {
		return nil, types.ErrAuctionBiddingAlreadyEnded
	}

	// Bid Validity check
	if msg.BidAmount >= auction_info.StartingPrice && msg.BidAmount > auction_info.CurrentHighestBid {
		auction_info.CurrentHighestBid = msg.BidAmount
		auction_info.CurrentHighestBidder = msg.Bidder
		k.Keeper.SetAuctionInfo(ctx, auction_info)
	}

	return &types.MsgPlaceBidResponse{}, nil
}
