package auction

import (
	"auction/x/auction/keeper"
	"auction/x/auction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the auctionInfo
	for _, elem := range genState.AuctionInfoList {
		k.SetAuctionInfo(ctx, elem)
	}
	// Set all the activeAuctionsList
	for _, elem := range genState.ActiveAuctionsListList {
		k.SetActiveAuctionsList(ctx, elem)
	}

	// Set activeAuctionsList count
	k.SetActiveAuctionsListCount(ctx, genState.ActiveAuctionsListCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.AuctionInfoList = k.GetAllAuctionInfo(ctx)
	genesis.ActiveAuctionsListList = k.GetAllActiveAuctionsList(ctx)
	genesis.ActiveAuctionsListCount = k.GetActiveAuctionsListCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
