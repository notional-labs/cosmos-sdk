package v042

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/kv"
	v042auth "github.com/cosmos/cosmos-sdk/x/auth/migrations/v042"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/golang/protobuf/proto"
)

const (
	// ModuleName defines the module name
	ModuleName = "bank"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName
)

// KVStore keys
var (
	BalancesPrefix      = []byte("balances")
	SupplyKey           = []byte{0x00}
	DenomMetadataPrefix = []byte{0x1}
)

// DenomMetadataKey returns the denomination metadata key.
func DenomMetadataKey(denom string) []byte {
	d := []byte(denom)
	return append(DenomMetadataPrefix, d...)
}

// AddressFromBalancesStore returns an account address from a balances prefix
// store. The key must not contain the perfix BalancesPrefix as the prefix store
// iterator discards the actual prefix.
func AddressFromBalancesStore(key []byte) sdk.AccAddress {
	kv.AssertKeyAtLeastLength(key, 1+v042auth.AddrLen)
	addr := key[:v042auth.AddrLen]
	kv.AssertKeyLength(addr, v042auth.AddrLen)
	return sdk.AccAddress(addr)
}

// SupplyI defines an inflationary supply interface for modules that handle
// token supply.
// It is copy-pasted from:
// https://github.com/cosmos/cosmos-sdk/blob/v042.3/x/bank/exported/exported.go
// where we stripped off the unnecessary methods.
//
// It is used in the migration script, because we save this interface as an Any
// in the supply state.
//
// Deprecated.
type SupplyI interface {
	proto.Message
}

// RegisterInterfaces registers interfaces required for the v0.40 migrations.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterInterface(
		"cosmos.bank.v1beta1.SupplyI",
		(*SupplyI)(nil),
		&types.Supply{},
	)
}
