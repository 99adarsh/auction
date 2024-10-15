package keeper_test

import (
	"testing"

	keepertest "auction/testutil/keeper"
	"auction/testutil/nullify"
	"auction/x/auction/keeper"
	"auction/x/auction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNActiveAuctionsList(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ActiveAuctionsList {
	items := make([]types.ActiveAuctionsList, n)
	for i := range items {
		items[i].Id = keeper.AppendActiveAuctionsList(ctx, items[i])
	}
	return items
}

func TestActiveAuctionsListGet(t *testing.T) {
	keeper, ctx := keepertest.AuctionKeeper(t)
	items := createNActiveAuctionsList(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetActiveAuctionsList(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestActiveAuctionsListRemove(t *testing.T) {
	keeper, ctx := keepertest.AuctionKeeper(t)
	items := createNActiveAuctionsList(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveActiveAuctionsList(ctx, item.Id)
		_, found := keeper.GetActiveAuctionsList(ctx, item.Id)
		require.False(t, found)
	}
}

func TestActiveAuctionsListGetAll(t *testing.T) {
	keeper, ctx := keepertest.AuctionKeeper(t)
	items := createNActiveAuctionsList(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllActiveAuctionsList(ctx)),
	)
}

func TestActiveAuctionsListCount(t *testing.T) {
	keeper, ctx := keepertest.AuctionKeeper(t)
	items := createNActiveAuctionsList(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetActiveAuctionsListCount(ctx))
}
