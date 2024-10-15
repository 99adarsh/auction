package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AuctionInfoList: []AuctionInfo{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in auctionInfo
	auctionInfoIndexMap := make(map[string]struct{})

	for _, elem := range gs.AuctionInfoList {
		index := string(AuctionInfoKey(elem.AuctionId))
		if _, ok := auctionInfoIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for auctionInfo")
		}
		auctionInfoIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
