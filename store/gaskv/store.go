package gaskv

import (
	"fmt"
	"io"
	"time"

	"github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
)

var _ types.KVStore = &Store{}

// Store applies gas tracking to an underlying KVStore. It implements the
// KVStore interface.
type Store struct {
	gasMeter  types.GasMeter
	gasConfig types.GasConfig
	parent    types.KVStore
}

// NewStore returns a reference to a new GasKVStore.
func NewStore(parent types.KVStore, gasMeter types.GasMeter, gasConfig types.GasConfig) *Store {
	kvs := &Store{
		gasMeter:  gasMeter,
		gasConfig: gasConfig,
		parent:    parent,
	}
	return kvs
}

// Implements Store.
func (gs *Store) GetStoreType() types.StoreType {
	return gs.parent.GetStoreType()
}

func (gs *Store) GetWorkingHash() []byte {
	return gs.parent.GetWorkingHash()
}

// Implements KVStore.
func (gs *Store) Get(key []byte) (value []byte) {
	gs.gasMeter.ConsumeGas(gs.gasConfig.ReadCostFlat, types.GasReadCostFlatDesc)
	value = gs.parent.Get(key)

	keyGas := 0 * types.Gas(len(key))
	valueGas := 0 * types.Gas(len(value))

	println(fmt.Printf("GasFee:READ:key key=%s gas=%d bytes=%d configuredGasPerByte=%d actualGas=%d", string(key), keyGas, len(key), gs.gasConfig.ReadCostPerByte, gs.gasConfig.ReadCostPerByte * types.Gas(len(key))))
	println(fmt.Printf("GasFee:READ:value gas=%d bytes=%d configuredGasPerByte=%d actualGas=%d", keyGas, len(value), gs.gasConfig.ReadCostPerByte, gs.gasConfig.ReadCostPerByte * types.Gas(len(value))))

	// TODO overflow-safe math?
	gs.gasMeter.ConsumeGas(keyGas, types.GasReadPerByteDesc)
	gs.gasMeter.ConsumeGas(valueGas, types.GasReadPerByteDesc)

	return value
}

// Implements KVStore.
func (gs *Store) Set(key []byte, value []byte) {
	types.AssertValidKey(key)
	types.AssertValidValue(value)


	keyGas := 0 * types.Gas(len(key))
	valueGas := 0 * types.Gas(len(value))

	gs.gasMeter.ConsumeGas(gs.gasConfig.WriteCostFlat, types.GasWriteCostFlatDesc)

	// TODO overflow-safe math?
	gs.gasMeter.ConsumeGas(keyGas, types.GasWritePerByteDesc)
	gs.gasMeter.ConsumeGas(valueGas, types.GasWritePerByteDesc)

	println(fmt.Printf("GasFee:WRITE:key key=%s gas=%d bytes=%d configuredGasPerByte=%d actualGas=%d", string(key), keyGas, len(key), gs.gasConfig.WriteCostPerByte, gs.gasConfig.WriteCostPerByte * types.Gas(len(key))))
	println(fmt.Printf("GasFee:WRITE:value gas=%d bytes=%d configuredGasPerByte=%d actualGas=%d", keyGas, len(value), gs.gasConfig.WriteCostPerByte, gs.gasConfig.WriteCostPerByte * types.Gas(len(value))))

	gs.parent.Set(key, value)
}

// Implements KVStore.
func (gs *Store) Has(key []byte) bool {
	defer telemetry.MeasureSince(time.Now(), "store", "gaskv", "has")
	gs.gasMeter.ConsumeGas(gs.gasConfig.HasCost, types.GasHasDesc)
	return gs.parent.Has(key)
}

// Implements KVStore.
func (gs *Store) Delete(key []byte) {
	defer telemetry.MeasureSince(time.Now(), "store", "gaskv", "delete")
	// charge gas to prevent certain attack vectors even though space is being freed
	gs.gasMeter.ConsumeGas(gs.gasConfig.DeleteCost, types.GasDeleteDesc)
	gs.parent.Delete(key)
}

// Iterator implements the KVStore interface. It returns an iterator which
// incurs a flat gas cost for seeking to the first key/value pair and a variable
// gas cost based on the current value's length if the iterator is valid.
func (gs *Store) Iterator(start, end []byte) types.Iterator {
	return gs.iterator(start, end, true)
}

// ReverseIterator implements the KVStore interface. It returns a reverse
// iterator which incurs a flat gas cost for seeking to the first key/value pair
// and a variable gas cost based on the current value's length if the iterator
// is valid.
func (gs *Store) ReverseIterator(start, end []byte) types.Iterator {
	return gs.iterator(start, end, false)
}

// Implements KVStore.
func (gs *Store) CacheWrap() types.CacheWrap {
	panic("cannot CacheWrap a GasKVStore")
}

// CacheWrapWithTrace implements the KVStore interface.
func (gs *Store) CacheWrapWithTrace(_ io.Writer, _ types.TraceContext) types.CacheWrap {
	panic("cannot CacheWrapWithTrace a GasKVStore")
}

// CacheWrapWithListeners implements the CacheWrapper interface.
func (gs *Store) CacheWrapWithListeners(_ types.StoreKey, _ []types.WriteListener) types.CacheWrap {
	panic("cannot CacheWrapWithListeners a GasKVStore")
}

func (gs *Store) iterator(start, end []byte, ascending bool) types.Iterator {
	var parent types.Iterator
	if ascending {
		parent = gs.parent.Iterator(start, end)
	} else {
		parent = gs.parent.ReverseIterator(start, end)
	}

	gi := newGasIterator(gs.gasMeter, gs.gasConfig, parent)
	gi.(*gasIterator).consumeSeekGas()

	return gi
}

type gasIterator struct {
	gasMeter  types.GasMeter
	gasConfig types.GasConfig
	parent    types.Iterator
}

func newGasIterator(gasMeter types.GasMeter, gasConfig types.GasConfig, parent types.Iterator) types.Iterator {
	return &gasIterator{
		gasMeter:  gasMeter,
		gasConfig: gasConfig,
		parent:    parent,
	}
}

// Implements Iterator.
func (gi *gasIterator) Domain() (start []byte, end []byte) {
	return gi.parent.Domain()
}

// Implements Iterator.
func (gi *gasIterator) Valid() bool {
	return gi.parent.Valid()
}

// Next implements the Iterator interface. It seeks to the next key/value pair
// in the iterator. It incurs a flat gas cost for seeking and a variable gas
// cost based on the current value's length if the iterator is valid.
func (gi *gasIterator) Next() {
	gi.consumeSeekGas()
	gi.parent.Next()
}

// Key implements the Iterator interface. It returns the current key and it does
// not incur any gas cost.
func (gi *gasIterator) Key() (key []byte) {
	key = gi.parent.Key()
	return key
}

// Value implements the Iterator interface. It returns the current value and it
// does not incur any gas cost.
func (gi *gasIterator) Value() (value []byte) {
	value = gi.parent.Value()
	return value
}

// Implements Iterator.
func (gi *gasIterator) Close() error {
	return gi.parent.Close()
}

// Error delegates the Error call to the parent iterator.
func (gi *gasIterator) Error() error {
	return gi.parent.Error()
}

// consumeSeekGas consumes on each iteration step a flat gas cost and a variable gas cost
// based on the current value's length.
func (gi *gasIterator) consumeSeekGas() {
	if gi.Valid() {
		key := gi.Key()
		value := gi.Value()

		keyGas := 0 * types.Gas(len(key))
		valueGas := 0 * types.Gas(len(value))
		println(fmt.Printf("GasFee:SEEK:key key=%s gas=%d bytes=%d configuredGasPerByte=%d actualGas=%d", string(key), keyGas, len(key), gi.gasConfig.ReadCostPerByte, gi.gasConfig.ReadCostPerByte*types.Gas(len(key))))
		println(fmt.Printf("GasFee:SEEK:value gas=%d bytes=%d configuredGasPerByte=%d actualGas=%d", keyGas, len(value), gi.gasConfig.ReadCostPerByte, gi.gasConfig.ReadCostPerByte*types.Gas(len(value))))

		gi.gasMeter.ConsumeGas(keyGas, types.GasValuePerByteDesc)
		gi.gasMeter.ConsumeGas(valueGas, types.GasValuePerByteDesc)
	}
	gi.gasMeter.ConsumeGas(gi.gasConfig.IterNextCostFlat, types.GasIterNextCostFlatDesc)
}
