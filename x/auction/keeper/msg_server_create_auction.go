package keeper

import (
	"context"
	"strconv"

	"auction/x/auction/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateAuction(goCtx context.Context, msg *types.MsgCreateAuction) (*types.MsgCreateAuctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check starting price > 0
	if msg.StartingPrice == 0 {
		return nil, types.ErrZeroStartingAuctionPrice
	}
	// check duration should be > 0
	if msg.DurationBlocks == 0 {
		return nil, types.ErrZeroAuctionBlockDuration
	}

	auctionEndHeight := ctx.BlockHeight() + int64(msg.DurationBlocks)
	// create the auction info and save it in the memory
	// create new auction id, check for existing auction against that id
	// TODO: Create more appropriate id
	newAuctionId := msg.ItemName + strconv.Itoa(int(ctx.BlockHeight()))
	// TODO: verify if their is any other auction store against newAuctionId

	// Initially, bidder is empty string and the bid amount is 0
	newAuctionInfo := types.AuctionInfo{
		AuctionId:            newAuctionId,
		ItemName:             msg.ItemName,
		StartingPrice:        msg.StartingPrice,
		AuctionEndHeight:     uint64(auctionEndHeight),
		CurrentHighestBid:    0,
		CurrentHighestBidder: "",
	}

	k.Keeper.SetAuctionInfo(ctx, newAuctionInfo)

	return &types.MsgCreateAuctionResponse{}, nil
}
