package keeper

import (
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
)

// Keeper manages history of public keys per account
type Keeper struct {
	key storetypes.StoreKey
	cdc codec.BinaryCodec
}

func NewKeeper(storeKey storetypes.StoreKey, cdc codec.BinaryCodec) Keeper {
	return Keeper{
		key: storeKey,
		cdc: cdc,
	}
}
