package keeper

import (
	"auction/x/auction/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetAuctionInfo set a specific auctionInfo in the store from its index
func (k Keeper) SetAuctionInfo(ctx sdk.Context, auctionInfo types.AuctionInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionInfoKeyPrefix))
	b := k.cdc.MustMarshal(&auctionInfo)
	store.Set(types.AuctionInfoKey(
		auctionInfo.AuctionId,
	), b)
}

// GetAuctionInfo returns a auctionInfo from its index
func (k Keeper) GetAuctionInfo(
	ctx sdk.Context,
	auctionId string,

) (val types.AuctionInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionInfoKeyPrefix))

	b := store.Get(types.AuctionInfoKey(
		auctionId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAuctionInfo removes a auctionInfo from the store
func (k Keeper) RemoveAuctionInfo(
	ctx sdk.Context,
	auctionId string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionInfoKeyPrefix))
	store.Delete(types.AuctionInfoKey(
		auctionId,
	))
}

// GetAllAuctionInfo returns all auctionInfo
func (k Keeper) GetAllAuctionInfo(ctx sdk.Context) (list []types.AuctionInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionInfoKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AuctionInfo
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
