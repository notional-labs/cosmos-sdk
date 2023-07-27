package types_test

import (
	"github.com/cosmos/cosmos-sdk/simapp"
)

var (
	app      = simapp.Setup(false)
	ecdc     = simapp.MakeTestEncodingConfig()
	appCodec = ecdc.Marshaler
)
