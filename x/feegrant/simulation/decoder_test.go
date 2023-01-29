package simulation_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cheqd/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cheqd/cosmos-sdk/simapp"
	sdk "github.com/cheqd/cosmos-sdk/types"
	"github.com/cheqd/cosmos-sdk/types/kv"
	"github.com/cheqd/cosmos-sdk/x/feegrant"
	"github.com/cheqd/cosmos-sdk/x/feegrant/simulation"
)

var (
	granterPk   = ed25519.GenPrivKey().PubKey()
	granterAddr = sdk.AccAddress(granterPk.Address())
	granteePk   = ed25519.GenPrivKey().PubKey()
	granteeAddr = sdk.AccAddress(granterPk.Address())
)

func TestDecodeStore(t *testing.T) {
	cdc := simapp.MakeTestEncodingConfig().Marshaler
	dec := simulation.NewDecodeStore(cdc)

	grant, err := feegrant.NewGrant(granterAddr, granteeAddr, &feegrant.BasicAllowance{
		SpendLimit: sdk.NewCoins(sdk.NewCoin("foo", sdk.NewInt(100))),
	})

	require.NoError(t, err)

	grantBz, err := cdc.Marshal(&grant)
	require.NoError(t, err)

	kvPairs := kv.Pairs{
		Pairs: []kv.Pair{
			{Key: []byte(feegrant.FeeAllowanceKeyPrefix), Value: grantBz},
			{Key: []byte{0x99}, Value: []byte{0x99}},
		},
	}

	tests := []struct {
		name        string
		expectedLog string
	}{
		{"Grant", fmt.Sprintf("%v\n%v", grant, grant)},
		{"other", ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			switch i {
			case len(tests) - 1:
				require.Panics(t, func() { dec(kvPairs.Pairs[i], kvPairs.Pairs[i]) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, dec(kvPairs.Pairs[i], kvPairs.Pairs[i]), tt.name)
			}
		})
	}
}
