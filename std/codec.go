package std

import (
	"github.com/cheqd/cosmos-sdk/codec"
	"github.com/cheqd/cosmos-sdk/codec/types"
	cryptocodec "github.com/cheqd/cosmos-sdk/crypto/codec"
	sdk "github.com/cheqd/cosmos-sdk/types"
	txtypes "github.com/cheqd/cosmos-sdk/types/tx"
)

// RegisterLegacyAminoCodec registers types with the Amino codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptocodec.RegisterCrypto(cdc)
	codec.RegisterEvidences(cdc)
}

// RegisterInterfaces registers Interfaces from sdk/types, vesting, crypto, tx.
func RegisterInterfaces(interfaceRegistry types.InterfaceRegistry) {
	sdk.RegisterInterfaces(interfaceRegistry)
	txtypes.RegisterInterfaces(interfaceRegistry)
	cryptocodec.RegisterInterfaces(interfaceRegistry)
}
