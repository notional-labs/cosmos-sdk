package rootmulti_test

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/store/iavl"
	"github.com/cosmos/cosmos-sdk/store/rootmulti"
	"github.com/cosmos/cosmos-sdk/store/types"
	dbm "github.com/tendermint/tm-db"
)

func newMultiStoreWithGeneratedData(db dbm.DB, stores uint8, storeKeys uint64) *rootmulti.Store {
	multiStore := rootmulti.NewStore(db, log.NewNopLogger())
	r := rand.New(rand.NewSource(49872768940)) // Fixed seed for deterministic tests

	keys := []*types.KVStoreKey{}
	for i := uint8(0); i < stores; i++ {
		key := types.NewKVStoreKey(fmt.Sprintf("store%v", i))
		multiStore.MountStoreWithDB(key, types.StoreTypeIAVL, nil)
		keys = append(keys, key)
	}
	multiStore.LoadLatestVersion()

	for _, key := range keys {
		store := multiStore.GetCommitKVStore(key).(*iavl.Store)
		for i := uint64(0); i < storeKeys; i++ {
			k := make([]byte, 8)
			v := make([]byte, 1024)
			binary.BigEndian.PutUint64(k, i)
			_, err := r.Read(v)
			if err != nil {
				panic(err)
			}
			store.Set(k, v)
		}
	}

	multiStore.Commit()
	multiStore.LoadLatestVersion()

	return multiStore
}

func newMultiStoreWithMixedMounts(db dbm.DB) *rootmulti.Store {
	store := rootmulti.NewStore(db, log.NewNopLogger())
	store.MountStoreWithDB(types.NewKVStoreKey("iavl1"), types.StoreTypeIAVL, nil)
	store.MountStoreWithDB(types.NewKVStoreKey("iavl2"), types.StoreTypeIAVL, nil)
	store.MountStoreWithDB(types.NewKVStoreKey("iavl3"), types.StoreTypeIAVL, nil)
	store.MountStoreWithDB(types.NewTransientStoreKey("trans1"), types.StoreTypeTransient, nil)
	store.LoadLatestVersion()

	return store
}

func newMultiStoreWithMixedMountsAndBasicData(db dbm.DB) *rootmulti.Store {
	store := newMultiStoreWithMixedMounts(db)
	store1 := store.GetStoreByName("iavl1").(types.CommitKVStore)
	store2 := store.GetStoreByName("iavl2").(types.CommitKVStore)
	trans1 := store.GetStoreByName("trans1").(types.KVStore)

	store1.Set([]byte("a"), []byte{1})
	store1.Set([]byte("b"), []byte{1})
	store2.Set([]byte("X"), []byte{255})
	store2.Set([]byte("A"), []byte{101})
	trans1.Set([]byte("x1"), []byte{91})
	store.Commit()

	store1.Set([]byte("b"), []byte{2})
	store1.Set([]byte("c"), []byte{3})
	store2.Set([]byte("B"), []byte{102})
	store.Commit()

	store2.Set([]byte("C"), []byte{103})
	store2.Delete([]byte("X"))
	trans1.Set([]byte("x2"), []byte{92})
	store.Commit()

	return store
}

func assertStoresEqual(t *testing.T, expect, actual types.CommitKVStore, msgAndArgs ...interface{}) {
	assert.Equal(t, expect.LastCommitID(), actual.LastCommitID())
	expectIter := expect.Iterator(nil, nil)
	expectMap := map[string][]byte{}
	for ; expectIter.Valid(); expectIter.Next() {
		expectMap[string(expectIter.Key())] = expectIter.Value()
	}
	require.NoError(t, expectIter.Error())

	actualIter := expect.Iterator(nil, nil)
	actualMap := map[string][]byte{}
	for ; actualIter.Valid(); actualIter.Next() {
		actualMap[string(actualIter.Key())] = actualIter.Value()
	}
	require.NoError(t, actualIter.Error())

	assert.Equal(t, expectMap, actualMap, msgAndArgs...)
}

func TestMultistoreSnapshot_Errors(t *testing.T) {
	store := newMultiStoreWithMixedMountsAndBasicData(dbm.NewMemDB())

	testcases := map[string]struct {
		height     uint64
		expectType error
	}{
		"0 height":       {0, nil},
		"unknown height": {9, nil},
	}
	for name, tc := range testcases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			_, err := store.Snapshot(tc.height, 1)
			require.Error(t, err)
			if tc.expectType != nil {
				assert.True(t, errors.Is(err, tc.expectType))
			}
		})
	}
}
