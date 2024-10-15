package keeper

import (
	"context"

	"auction/x/auction/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ActiveAuctionsListAll(goCtx context.Context, req *types.QueryAllActiveAuctionsListRequest) (*types.QueryAllActiveAuctionsListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var activeAuctionsLists []types.ActiveAuctionsList
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	activeAuctionsListStore := prefix.NewStore(store, types.KeyPrefix(types.ActiveAuctionsListKey))

	pageRes, err := query.Paginate(activeAuctionsListStore, req.Pagination, func(key []byte, value []byte) error {
		var activeAuctionsList types.ActiveAuctionsList
		if err := k.cdc.Unmarshal(value, &activeAuctionsList); err != nil {
			return err
		}

		activeAuctionsLists = append(activeAuctionsLists, activeAuctionsList)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllActiveAuctionsListResponse{ActiveAuctionsList: activeAuctionsLists, Pagination: pageRes}, nil
}

func (k Keeper) ActiveAuctionsList(goCtx context.Context, req *types.QueryGetActiveAuctionsListRequest) (*types.QueryGetActiveAuctionsListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	activeAuctionsList, found := k.GetActiveAuctionsList(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetActiveAuctionsListResponse{ActiveAuctionsList: activeAuctionsList}, nil
}
