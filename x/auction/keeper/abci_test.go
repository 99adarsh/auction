package keeper_test

import (
	"auction/x/auction/types"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func (suite *KeeperTestSuite) TestEndBlock() {
	ctx := suite.ctx
	msgServer := suite.msgServer

	// Step 1: Create an auction with a defined duration (e.g., 10 blocks)
	createMsg := &types.MsgCreateAuction{
		Creator:        types.Alice,
		ItemName:       "TestItem",
		StartingPrice:  100,
		DurationBlocks: 10, // Auction ends in 10 blocks
	}
	_, err := msgServer.CreateAuction(sdk.WrapSDKContext(ctx), createMsg)
	require.NoError(suite.T(), err)

	auctionId := createMsg.ItemName + strconv.Itoa(int(ctx.BlockHeight()))

	// Step 2: Verify that the auction is added to the active auctions list
	activeAuctions := suite.app.AuctionKeeper.GetAllActiveAuctionsList(ctx)
	require.Len(suite.T(), activeAuctions, 1)
	require.Equal(suite.T(), auctionId, activeAuctions[0].AuctionId)

	// Step 3: Simulate block height increase to auction end height
	auctionEndHeight := ctx.BlockHeight() + 10
	ctx = ctx.WithBlockHeight(auctionEndHeight)

	// Step 4: Call EndBlock function
	err = suite.app.AuctionKeeper.EndBlock(ctx)
	require.NoError(suite.T(), err)

	// Step 5: Verify that the auction is removed from the active auctions list
	activeAuctions = suite.app.AuctionKeeper.GetAllActiveAuctionsList(ctx)
	require.Len(suite.T(), activeAuctions, 0)

	// Step 6: Verify that the auction info is still available (i.e., only removed from active list)
	auctionInfo, found := suite.app.AuctionKeeper.GetAuctionInfo(ctx, auctionId)
	require.True(suite.T(), found)
	require.Equal(suite.T(), auctionId, auctionInfo.AuctionId)
}
