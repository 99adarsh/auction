package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/auction module sentinel errors
var (
	ErrZeroStartingAuctionPrice = sdkerrors.Register(ModuleName, 1101, "Auction starting price cannot be zero")
	ErrZeroAuctionBlockDuration = sdkerrors.Register(ModuleName, 1102, "Auction duration height cannot be zero")
)
