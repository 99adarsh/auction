package keeper_test

import (
	"auction/x/auction/types"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// TestCreateAuctionSuccess tests successful creation of an auction.
func (suite *KeeperTestSuite) TestCreateAuctionSuccess() {
	ctx := suite.ctx
	msgServer := suite.msgServer

	// Define a valid MsgCreateAuction request
	msg := &types.MsgCreateAuction{
		Creator:        types.Alice,       // Mocked creator address
		ItemName:       "TestSuccessItem", // Example item name
		StartingPrice:  100,               // Starting price > 0
		DurationBlocks: 50,                // Auction duration > 0
	}

	// Call CreateAuction
	response, err := msgServer.CreateAuction(sdk.WrapSDKContext(ctx), msg)

	// Validate the response
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), response)

	// Check if the auction is properly stored
	auctionId := msg.ItemName + strconv.Itoa(int(ctx.BlockHeight()))
	auctionInfo, found := suite.app.AuctionKeeper.GetAuctionInfo(ctx, auctionId)
	require.True(suite.T(), found)
	require.Equal(suite.T(), msg.ItemName, auctionInfo.ItemName)
	require.Equal(suite.T(), msg.StartingPrice, auctionInfo.StartingPrice)
	require.Equal(suite.T(), uint64(ctx.BlockHeight()+int64(msg.DurationBlocks)), auctionInfo.AuctionEndHeight)
	require.Equal(suite.T(), "", auctionInfo.CurrentHighestBidder)
	require.Equal(suite.T(), uint64(0), auctionInfo.CurrentHighestBid)
}

// TestCreateAuctionInvalidStartingPrice tests error when starting price is 0.
func (suite *KeeperTestSuite) TestCreateAuctionInvalidStartingPrice() {
	ctx := suite.ctx
	msgServer := suite.msgServer

	// Define a MsgCreateAuction request with zero starting price
	msg := &types.MsgCreateAuction{
		Creator:        types.Alice,
		ItemName:       "TestInvalidStartingPriceItem",
		StartingPrice:  0, // Invalid starting price
		DurationBlocks: 50,
	}

	// Call CreateAuction
	_, err := msgServer.CreateAuction(sdk.WrapSDKContext(ctx), msg)

	// Validate the error
	require.Error(suite.T(), err)
	require.Equal(suite.T(), types.ErrZeroStartingAuctionPrice, err)
}

// TestCreateAuctionInvalidDuration tests error when duration blocks are 0.
func (suite *KeeperTestSuite) TestCreateAuctionInvalidDuration() {
	ctx := suite.ctx
	msgServer := suite.msgServer

	// Define a MsgCreateAuction request with zero duration blocks
	msg := &types.MsgCreateAuction{
		Creator:        types.Alice,
		ItemName:       "TestInvalidDurationItem",
		StartingPrice:  100,
		DurationBlocks: 0, // Invalid duration blocks
	}

	// Call CreateAuction
	_, err := msgServer.CreateAuction(sdk.WrapSDKContext(ctx), msg)

	// Validate the error
	require.Error(suite.T(), err)
	require.Equal(suite.T(), types.ErrZeroAuctionBlockDuration, err)
}
