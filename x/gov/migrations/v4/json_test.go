package v4_test

import (
	"encoding/json"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/cosmos/cosmos-sdk/x/gov"
	v4 "github.com/cosmos/cosmos-sdk/x/gov/migrations/v4"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/stretchr/testify/require"
)

func TestMigrateJSON(t *testing.T) {
	encodingConfig := moduletestutil.MakeTestEncodingConfig(gov.AppModuleBasic{})
	clientCtx := client.Context{}.
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithCodec(encodingConfig.Codec)

	govGenState := v1.DefaultGenesisState()

	migrated, err := v4.MigrateJSON(govGenState)
	require.NoError(t, err)

	bz, err := clientCtx.Codec.MarshalJSON(migrated)
	require.NoError(t, err)

	// Indent the JSON bz correctly.
	var jsonObj map[string]interface{}
	err = json.Unmarshal(bz, &jsonObj)
	require.NoError(t, err)
	indentedBz, err := json.MarshalIndent(jsonObj, "", "\t")
	require.NoError(t, err)

	// Make sure about:
	// - Proposals use MsgExecLegacyContent
	expected := `{
	"constitution": "",
	"deposit_params": null,
	"deposits": [],
	"params": {
		"max_deposit_period": "172800s",
		"min_deposit": [
			{
				"amount": "10000000",
				"denom": "stake"
			}
		],
		"min_initial_deposit_ratio": "0.000000000000000000",
		"quorum": "0.334000000000000000",
		"threshold": "0.500000000000000000",
		"veto_threshold": "0.334000000000000000",
		"voting_period": "172800s"
	},
	"proposals": [],
	"starting_proposal_id": "1",
	"tally_params": null,
	"votes": [],
	"voting_params": null
}`

	require.Equal(t, expected, string(indentedBz))
}
