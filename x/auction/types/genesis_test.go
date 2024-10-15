package types_test

import (
	"testing"

	"auction/x/auction/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				AuctionInfoList: []types.AuctionInfo{
					{
						AuctionId: "0",
					},
					{
						AuctionId: "1",
					},
				},
				ActiveAuctionsListList: []types.ActiveAuctionsList{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				ActiveAuctionsListCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated auctionInfo",
			genState: &types.GenesisState{
				AuctionInfoList: []types.AuctionInfo{
					{
						AuctionId: "0",
					},
					{
						AuctionId: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated activeAuctionsList",
			genState: &types.GenesisState{
				ActiveAuctionsListList: []types.ActiveAuctionsList{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid activeAuctionsList count",
			genState: &types.GenesisState{
				ActiveAuctionsListList: []types.ActiveAuctionsList{
					{
						Id: 1,
					},
				},
				ActiveAuctionsListCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
