package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "auction/testutil/keeper"
	"auction/testutil/nullify"
	"auction/x/auction/types"
)

func TestActiveAuctionsListQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.AuctionKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNActiveAuctionsList(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetActiveAuctionsListRequest
		response *types.QueryGetActiveAuctionsListResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetActiveAuctionsListRequest{Id: msgs[0].Id},
			response: &types.QueryGetActiveAuctionsListResponse{ActiveAuctionsList: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetActiveAuctionsListRequest{Id: msgs[1].Id},
			response: &types.QueryGetActiveAuctionsListResponse{ActiveAuctionsList: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetActiveAuctionsListRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ActiveAuctionsList(wctx, tc.request)
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

func TestActiveAuctionsListQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.AuctionKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNActiveAuctionsList(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllActiveAuctionsListRequest {
		return &types.QueryAllActiveAuctionsListRequest{
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
			resp, err := keeper.ActiveAuctionsListAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ActiveAuctionsList), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ActiveAuctionsList),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ActiveAuctionsListAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ActiveAuctionsList), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ActiveAuctionsList),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ActiveAuctionsListAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ActiveAuctionsList),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ActiveAuctionsListAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
