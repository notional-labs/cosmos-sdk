package types

import (
	"github.com/cheqd/cosmos-sdk/codec"
	cryptocodec "github.com/cheqd/cosmos-sdk/crypto/codec"
)

var amino = codec.NewLegacyAmino()

func init() {
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
