package keeper

import (
	"encoding/binary"

	"auction/x/auction/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetActiveAuctionsListCount get the total number of activeAuctionsList
func (k Keeper) GetActiveAuctionsListCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ActiveAuctionsListCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetActiveAuctionsListCount set the total number of activeAuctionsList
func (k Keeper) SetActiveAuctionsListCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ActiveAuctionsListCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendActiveAuctionsList appends a activeAuctionsList in the store with a new id and update the count
func (k Keeper) AppendActiveAuctionsList(
	ctx sdk.Context,
	activeAuctionsList types.ActiveAuctionsList,
) uint64 {
	// Create the activeAuctionsList
	count := k.GetActiveAuctionsListCount(ctx)

	// Set the ID of the appended value
	activeAuctionsList.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveAuctionsListKey))
	appendedValue := k.cdc.MustMarshal(&activeAuctionsList)
	store.Set(GetActiveAuctionsListIDBytes(activeAuctionsList.Id), appendedValue)

	// Update activeAuctionsList count
	k.SetActiveAuctionsListCount(ctx, count+1)

	return count
}

// SetActiveAuctionsList set a specific activeAuctionsList in the store
func (k Keeper) SetActiveAuctionsList(ctx sdk.Context, activeAuctionsList types.ActiveAuctionsList) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveAuctionsListKey))
	b := k.cdc.MustMarshal(&activeAuctionsList)
	store.Set(GetActiveAuctionsListIDBytes(activeAuctionsList.Id), b)
}

// GetActiveAuctionsList returns a activeAuctionsList from its id
func (k Keeper) GetActiveAuctionsList(ctx sdk.Context, id uint64) (val types.ActiveAuctionsList, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveAuctionsListKey))
	b := store.Get(GetActiveAuctionsListIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveActiveAuctionsList removes a activeAuctionsList from the store
func (k Keeper) RemoveActiveAuctionsList(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveAuctionsListKey))
	store.Delete(GetActiveAuctionsListIDBytes(id))
}

// GetAllActiveAuctionsList returns all activeAuctionsList
func (k Keeper) GetAllActiveAuctionsList(ctx sdk.Context) (list []types.ActiveAuctionsList) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveAuctionsListKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ActiveAuctionsList
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetActiveAuctionsListIDBytes returns the byte representation of the ID
func GetActiveAuctionsListIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetActiveAuctionsListIDFromBytes returns ID in uint64 format from a byte array
func GetActiveAuctionsListIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
