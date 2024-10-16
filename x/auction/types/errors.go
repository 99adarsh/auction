package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/auction module sentinel errors
var (
	ErrZeroStartingAuctionPrice   = sdkerrors.Register(ModuleName, 1101, "Auction starting price cannot be zero")
	ErrZeroAuctionBlockDuration   = sdkerrors.Register(ModuleName, 1102, "Auction duration height cannot be zero")
	ErrNewAuctionIdAlreadyExists  = sdkerrors.Register(ModuleName, 1103, "This AuctionId already exists")
	ErrAuctionInfoNotFound        = sdkerrors.Register(ModuleName, 1104, "AuctionInfo does not exist this auction id")
	ErrAuctionBiddingAlreadyEnded = sdkerrors.Register(ModuleName, 1105, "Auction bidding already ended for this auction id")
)
