package keeper_test

import (
	"strconv"
	"testing"

	keepertest "auction/testutil/keeper"
	"auction/testutil/nullify"
	"auction/x/auction/keeper"
	"auction/x/auction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNAuctionInfo(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.AuctionInfo {
	items := make([]types.AuctionInfo, n)
	for i := range items {
		items[i].AuctionId = strconv.Itoa(i)

		keeper.SetAuctionInfo(ctx, items[i])
	}
	return items
}

func TestAuctionInfoGet(t *testing.T) {
	keeper, ctx := keepertest.AuctionKeeper(t)
	items := createNAuctionInfo(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetAuctionInfo(ctx,
			item.AuctionId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestAuctionInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.AuctionKeeper(t)
	items := createNAuctionInfo(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAuctionInfo(ctx,
			item.AuctionId,
		)
		_, found := keeper.GetAuctionInfo(ctx,
			item.AuctionId,
		)
		require.False(t, found)
	}
}

func TestAuctionInfoGetAll(t *testing.T) {
	keeper, ctx := keepertest.AuctionKeeper(t)
	items := createNAuctionInfo(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAuctionInfo(ctx)),
	)
}
