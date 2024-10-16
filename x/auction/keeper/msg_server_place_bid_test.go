package keeper_test

import (
	"auction/x/auction/types"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// TestPlaceBidSuccess tests successful placement of a bid.
func (suite *KeeperTestSuite) TestPlaceBidSuccess() {
	ctx := suite.ctx
	msgServer := suite.msgServer

	// Create an auction to bid on
	createMsg := &types.MsgCreateAuction{
		Creator:        types.Alice,
		ItemName:       "TestItem",
		StartingPrice:  100,
		DurationBlocks: 50,
	}
	_, err := msgServer.CreateAuction(sdk.WrapSDKContext(ctx), createMsg)
	require.NoError(suite.T(), err)

	auctionId := createMsg.ItemName + strconv.Itoa(int(ctx.BlockHeight()))

	// Place a valid bid
	bidMsg := &types.MsgPlaceBid{
		AuctionId: auctionId,
		Bidder:    types.Bob, // Mocked bidder address
		BidAmount: 150,       // Valid bid amount > starting price
	}
	response, err := msgServer.PlaceBid(sdk.WrapSDKContext(ctx), bidMsg)

	// Validate the response
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), response)

	// Check the auction info is updated correctly
	auctionInfo, found := suite.app.AuctionKeeper.GetAuctionInfo(ctx, auctionId)
	require.True(suite.T(), found)
	require.Equal(suite.T(), uint64(150), auctionInfo.CurrentHighestBid)
	require.Equal(suite.T(), types.Bob, auctionInfo.CurrentHighestBidder)
}

// TestPlaceBidAuctionNotFound tests error when trying to bid on a non-existent auction.
func (suite *KeeperTestSuite) TestPlaceBidAuctionNotFound() {
	ctx := suite.ctx
	msgServer := suite.msgServer

	// Try to place a bid on a non-existent auction
	bidMsg := &types.MsgPlaceBid{
		AuctionId: "non_existent_auction",
		Bidder:    types.Bob,
		BidAmount: 150,
	}
	_, err := msgServer.PlaceBid(sdk.WrapSDKContext(ctx), bidMsg)

	// Validate the error
	require.Error(suite.T(), err)
	require.Equal(suite.T(), types.ErrAuctionInfoNotFound, err)
}

// TestPlaceBidAuctionEnded tests error when trying to bid after the auction has ended.
func (suite *KeeperTestSuite) TestPlaceBidAuctionEnded() {
	ctx := suite.ctx
	msgServer := suite.msgServer

	// Create an auction and set it to end immediately
	createMsg := &types.MsgCreateAuction{
		Creator:        types.Alice,
		ItemName:       "TestItem",
		StartingPrice:  100,
		DurationBlocks: 10, // Set the duration to 0, so it ends immediately
	}
	_, err := msgServer.CreateAuction(sdk.WrapSDKContext(ctx), createMsg)
	require.NoError(suite.T(), err)

	auctionId := createMsg.ItemName + strconv.Itoa(int(ctx.BlockHeight()))

	// increase the block height by 20( but auction ended 10 blocks before)
	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 20)

	// Try to place a bid after the auction has ended
	bidMsg := &types.MsgPlaceBid{
		AuctionId: auctionId,
		Bidder:    types.Bob,
		BidAmount: 150,
	}
	_, err = msgServer.PlaceBid(sdk.WrapSDKContext(ctx), bidMsg)

	// Validate the error
	require.Error(suite.T(), err)
	require.Equal(suite.T(), types.ErrAuctionBiddingAlreadyEnded, err)
}

// TestPlaceBidAmountTooLow tests error when the bid amount is lower than the highest bid or starting price.
func (suite *KeeperTestSuite) TestPlaceBidAmountTooLow() {
	ctx := suite.ctx
	msgServer := suite.msgServer

	// Create an auction to bid on
	createMsg := &types.MsgCreateAuction{
		Creator:        types.Alice,
		ItemName:       "TestItem",
		StartingPrice:  100,
		DurationBlocks: 50,
	}
	_, err := msgServer.CreateAuction(sdk.WrapSDKContext(ctx), createMsg)
	require.NoError(suite.T(), err)

	auctionId := createMsg.ItemName + strconv.Itoa(int(ctx.BlockHeight()))

	// Place a bid lower than the starting price
	bidMsg := &types.MsgPlaceBid{
		AuctionId: auctionId,
		Bidder:    types.Bob,
		BidAmount: 90, // Bid amount is less than starting price
	}
	_, err = msgServer.PlaceBid(sdk.WrapSDKContext(ctx), bidMsg)

	// Validate the error
	require.Error(suite.T(), err)
	require.Equal(suite.T(), types.ErrAuctionBidIsLesser, err)

	// Place a valid bid first, then try a lower bid
	validBidMsg := &types.MsgPlaceBid{
		AuctionId: auctionId,
		Bidder:    types.Bob,
		BidAmount: 150, // Valid bid
	}
	_, err = msgServer.PlaceBid(sdk.WrapSDKContext(ctx), validBidMsg)
	require.NoError(suite.T(), err)

	// Now try to place a lower bid
	lowerBidMsg := &types.MsgPlaceBid{
		AuctionId: auctionId,
		Bidder:    types.Dummy,
		BidAmount: 140, // Bid amount is less than the current highest bid
	}
	_, err = msgServer.PlaceBid(sdk.WrapSDKContext(ctx), lowerBidMsg)

	// Validate the error
	require.Error(suite.T(), err)
	require.Equal(suite.T(), types.ErrAuctionBidIsLesser, err)
}
