package keeper

import (
	"context"

	"auction/x/auction/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AuctionInfoAll(goCtx context.Context, req *types.QueryAllAuctionInfoRequest) (*types.QueryAllAuctionInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var auctionInfos []types.AuctionInfo
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	auctionInfoStore := prefix.NewStore(store, types.KeyPrefix(types.AuctionInfoKeyPrefix))

	pageRes, err := query.Paginate(auctionInfoStore, req.Pagination, func(key []byte, value []byte) error {
		var auctionInfo types.AuctionInfo
		if err := k.cdc.Unmarshal(value, &auctionInfo); err != nil {
			return err
		}

		auctionInfos = append(auctionInfos, auctionInfo)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAuctionInfoResponse{AuctionInfo: auctionInfos, Pagination: pageRes}, nil
}

func (k Keeper) AuctionInfo(goCtx context.Context, req *types.QueryGetAuctionInfoRequest) (*types.QueryGetAuctionInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetAuctionInfo(
		ctx,
		req.AuctionId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetAuctionInfoResponse{AuctionInfo: val}, nil
}
