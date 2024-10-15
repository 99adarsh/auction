package auction_test

import (
	"testing"

	keepertest "auction/testutil/keeper"
	"auction/testutil/nullify"
	"auction/x/auction"
	"auction/x/auction/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		AuctionInfoList: []types.AuctionInfo{
			{
				AuctionId: "0",
			},
			{
				AuctionId: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AuctionKeeper(t)
	auction.InitGenesis(ctx, *k, genesisState)
	got := auction.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.AuctionInfoList, got.AuctionInfoList)
	// this line is used by starport scaffolding # genesis/test/assert
}
