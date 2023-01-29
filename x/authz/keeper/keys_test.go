package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cheqd/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cheqd/cosmos-sdk/types"
	"github.com/cheqd/cosmos-sdk/types/address"
	bank "github.com/cheqd/cosmos-sdk/x/bank/types"
)

var (
	granter = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	grantee = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	msgType = bank.SendAuthorization{}.MsgTypeURL()
)

func TestGrantkey(t *testing.T) {
	require := require.New(t)
	key := grantStoreKey(grantee, granter, msgType)
	require.Len(key, len(GrantKey)+len(address.MustLengthPrefix(grantee))+len(address.MustLengthPrefix(granter))+len([]byte(msgType)))

	granter1, grantee1 := addressesFromGrantStoreKey(grantStoreKey(grantee, granter, msgType))
	require.Equal(granter, granter1)
	require.Equal(grantee, grantee1)
}
