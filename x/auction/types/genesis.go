package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AuctionInfoList:        []AuctionInfo{},
		ActiveAuctionsListList: []ActiveAuctionsList{},
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
	// Check for duplicated ID in activeAuctionsList
	activeAuctionsListIdMap := make(map[uint64]bool)
	activeAuctionsListCount := gs.GetActiveAuctionsListCount()
	for _, elem := range gs.ActiveAuctionsListList {
		if _, ok := activeAuctionsListIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for activeAuctionsList")
		}
		if elem.Id >= activeAuctionsListCount {
			return fmt.Errorf("activeAuctionsList id should be lower or equal than the last id")
		}
		activeAuctionsListIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
