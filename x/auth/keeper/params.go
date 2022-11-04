package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// SetParams sets the auth module's parameters.
func (ak AccountKeeper) SetParams(ctx sdk.Context, params types.Params) {
	ak.paramSubspace.SetParamSet(ctx, &params)
}

// // GetParams gets the auth module's parameters.
// func (ak AccountKeeper) GetParams(ctx sdk.Context) (params types.Params) {
// 	ak.paramSubspace.GetParamSet(ctx, &params)
// 	return
// }

// GetParams gets the auth module's parameters.
// Get all parameters as types.Params
func (ak AccountKeeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		ak.MaxMemoCharacters(ctx),
		ak.TxSigLimit(ctx),
		ak.TxSizeCostPerByte(ctx),
		ak.SigVerifyCostED25519(ctx),
		ak.SigVerifyCostSecp256k1(ctx),
		ak.TxFeeBurnPercent(ctx),
	)
}

// MaxMemoCharacters - Maximum memo charactes
func (ak AccountKeeper) MaxMemoCharacters(ctx sdk.Context) (res uint64) {
	ak.paramSubspace.Get(ctx, types.KeyMaxMemoCharacters, &res)
	return
}

// TxSigLimit - Transaction sig limit
func (ak AccountKeeper) TxSigLimit(ctx sdk.Context) (res uint64) {
	ak.paramSubspace.Get(ctx, types.KeyTxSigLimit, &res)
	return
}

// TxSizeCostPerByte - Transaction size cost per byte
func (ak AccountKeeper) TxSizeCostPerByte(ctx sdk.Context) (res uint64) {
	ak.paramSubspace.Get(ctx, types.KeyTxSizeCostPerByte, &res)
	return
}

// SigVerifyCostED25519 -
func (ak AccountKeeper) SigVerifyCostED25519(ctx sdk.Context) (res uint64) {
	ak.paramSubspace.Get(ctx, types.KeySigVerifyCostED25519, &res)
	return
}

// SigVerifyCostSecp256k1 -
func (ak AccountKeeper) SigVerifyCostSecp256k1(ctx sdk.Context) (res uint64) {
	ak.paramSubspace.Get(ctx, types.KeySigVerifyCostSecp256k1, &res)
	return
}

// TxFeeBurnPercent - Transaction fee burn percentage
// The percentage of each transaction fee that gets burned
func (ak AccountKeeper) TxFeeBurnPercent(ctx sdk.Context) (res sdk.Int) {
	ak.paramSubspace.Get(ctx, types.KeyTxFeeBurnPercent, &res)
	return
}
