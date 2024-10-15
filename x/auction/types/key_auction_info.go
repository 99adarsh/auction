package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// AuctionInfoKeyPrefix is the prefix to retrieve all AuctionInfo
	AuctionInfoKeyPrefix = "AuctionInfo/value/"
)

// AuctionInfoKey returns the store key to retrieve a AuctionInfo from the index fields
func AuctionInfoKey(
	auctionId string,
) []byte {
	var key []byte

	auctionIdBytes := []byte(auctionId)
	key = append(key, auctionIdBytes...)
	key = append(key, []byte("/")...)

	return key
}
