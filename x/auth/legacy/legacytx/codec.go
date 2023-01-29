package legacytx

import (
	"github.com/cheqd/cosmos-sdk/codec"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(StdTx{}, "cosmos-sdk/StdTx", nil)
}
