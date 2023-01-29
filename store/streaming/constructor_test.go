package streaming_test

import (
	"testing"

	"github.com/cheqd/cosmos-sdk/baseapp"
	"github.com/cheqd/cosmos-sdk/codec"
	codecTypes "github.com/cheqd/cosmos-sdk/codec/types"
	serverTypes "github.com/cheqd/cosmos-sdk/server/types"
	"github.com/cheqd/cosmos-sdk/simapp"
	"github.com/cheqd/cosmos-sdk/store/streaming"
	"github.com/cheqd/cosmos-sdk/store/streaming/file"
	"github.com/cheqd/cosmos-sdk/store/types"
	sdk "github.com/cheqd/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	"github.com/stretchr/testify/require"
)

type fakeOptions struct{}

func (f *fakeOptions) Get(key string) interface{} {
	if key == "streamers.file.write_dir" {
		return "data/file_streamer"

	}
	return nil
}

var (
	mockOptions       = new(fakeOptions)
	mockKeys          = []types.StoreKey{sdk.NewKVStoreKey("mockKey1"), sdk.NewKVStoreKey("mockKey2")}
	interfaceRegistry = codecTypes.NewInterfaceRegistry()
	testMarshaller    = codec.NewProtoCodec(interfaceRegistry)
)

func TestStreamingServiceConstructor(t *testing.T) {
	_, err := streaming.NewServiceConstructor("unexpectedName")
	require.NotNil(t, err)

	constructor, err := streaming.NewServiceConstructor("file")
	require.Nil(t, err)
	var expectedType streaming.ServiceConstructor
	require.IsType(t, expectedType, constructor)

	serv, err := constructor(mockOptions, mockKeys, testMarshaller)
	require.Nil(t, err)
	require.IsType(t, &file.StreamingService{}, serv)
	listeners := serv.Listeners()
	for _, key := range mockKeys {
		_, ok := listeners[key]
		require.True(t, ok)
	}
}

func TestLoadStreamingServices(t *testing.T) {
	db := dbm.NewMemDB()
	encCdc := simapp.MakeTestEncodingConfig()
	keys := sdk.NewKVStoreKeys("mockKey1", "mockKey2")
	bApp := baseapp.NewBaseApp("appName", log.NewNopLogger(), db, nil)

	testCases := map[string]struct {
		appOpts            serverTypes.AppOptions
		activeStreamersLen int
	}{
		"empty app options": {
			appOpts: simapp.EmptyAppOptions{},
		},
		"all StoreKeys exposed": {
			appOpts:            streamingAppOptions{keys: []string{"*"}},
			activeStreamersLen: 1,
		},
		"some StoreKey exposed": {
			appOpts:            streamingAppOptions{keys: []string{"mockKey1"}},
			activeStreamersLen: 1,
		},
		"not exposing anything": {
			appOpts: streamingAppOptions{keys: []string{"mockKey3"}},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			activeStreamers, _, err := streaming.LoadStreamingServices(bApp, tc.appOpts, encCdc.Marshaler, keys)
			require.NoError(t, err)
			require.Equal(t, tc.activeStreamersLen, len(activeStreamers))
		})
	}
}

type streamingAppOptions struct {
	keys []string
}

func (ao streamingAppOptions) Get(o string) interface{} {
	switch o {
	case "store.streamers":
		return []string{"file"}
	case "streamers.file.keys":
		return ao.keys
	case "streamers.file.write_dir":
		return "data/file_streamer"
	default:
		return nil
	}
}
