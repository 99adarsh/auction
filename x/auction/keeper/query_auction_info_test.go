package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "auction/testutil/keeper"
	"auction/testutil/nullify"
	"auction/x/auction/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestAuctionInfoQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.AuctionKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNAuctionInfo(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetAuctionInfoRequest
		response *types.QueryGetAuctionInfoResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetAuctionInfoRequest{
				AuctionId: msgs[0].AuctionId,
			},
			response: &types.QueryGetAuctionInfoResponse{AuctionInfo: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetAuctionInfoRequest{
				AuctionId: msgs[1].AuctionId,
			},
			response: &types.QueryGetAuctionInfoResponse{AuctionInfo: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetAuctionInfoRequest{
				AuctionId: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.AuctionInfo(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestAuctionInfoQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.AuctionKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNAuctionInfo(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllAuctionInfoRequest {
		return &types.QueryAllAuctionInfoRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.AuctionInfoAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.AuctionInfo), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.AuctionInfo),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.AuctionInfoAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.AuctionInfo), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.AuctionInfo),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.AuctionInfoAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.AuctionInfo),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.AuctionInfoAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
