package store

import (
	"fmt"
	"io"

	"github.com/cosmos/cosmos-sdk/store/cachemulti"
	"github.com/cosmos/cosmos-sdk/store/dbadapter"
	"github.com/cosmos/cosmos-sdk/store/mem"
	"github.com/cosmos/cosmos-sdk/store/tracekv"
	"github.com/cosmos/cosmos-sdk/store/transient"
	"github.com/cosmos/cosmos-sdk/types"
	dbm "github.com/tendermint/tm-db"
)

type Store struct {
	db           dbm.DB
	storesParams map[types.StoreKey]storeParams
	stores       map[types.StoreKey]types.CommitKVStore
	keysByName   map[string]types.StoreKey

	traceWriter  io.Writer
	traceContext types.TraceContext
}

func NewStore(db dbm.DB) *Store {
	return &Store{db: db}
}

// GetStoreType implements Store.
func (rs *Store) GetStoreType() types.StoreType {
	return types.StoreTypeMulti
}

// MountStoreWithDB implements CommitMultiStore.
func (rs *Store) MountStoreWithDB(key types.StoreKey, typ types.StoreType, db dbm.DB) {
	if key == nil {
		panic("MountIAVLStore() key cannot be nil")
	}
	if _, ok := rs.storesParams[key]; ok {
		panic(fmt.Sprintf("store duplicate store key %v", key))
	}
	if _, ok := rs.keysByName[key.Name()]; ok {
		panic(fmt.Sprintf("store duplicate store key name %v", key))
	}
	rs.storesParams[key] = storeParams{
		key: key,
		typ: typ,
		db:  db,
	}
	rs.keysByName[key.Name()] = key
}

// GetCommitStore returns a mounted CommitStore for a given StoreKey. If the
// store is wrapped in an inter-block cache, it will be unwrapped before returning.
func (rs *Store) GetCommitStore(key types.StoreKey) types.CommitStore {
	return rs.GetCommitKVStore(key)
}

// GetCommitKVStore returns a mounted CommitKVStore for a given StoreKey. If the
// store is wrapped in an inter-block cache, it will be unwrapped before returning.
func (rs *Store) GetCommitKVStore(key types.StoreKey) types.CommitKVStore {
	return rs.stores[key]
}

func (rs *Store) LoadStores() error {
	// load each Store (note this doesn't panic on unmounted keys now)
	var newStores = make(map[types.StoreKey]types.CommitKVStore)
	for key, storeParams := range rs.storesParams {
		store, err := rs.loadStoreFromParams(key, storeParams)
		if err != nil {
			return err
		}
		newStores[key] = store
	}
	rs.stores = newStores
	return nil
}

func (rs *Store) loadStoreFromParams(key types.StoreKey, params storeParams) (types.CommitKVStore, error) {
	var db dbm.DB

	if params.db != nil {
		db = dbm.NewPrefixDB(params.db, []byte("s/_/"))
	} else {
		prefix := "s/k:" + params.key.Name() + "/"
		db = dbm.NewPrefixDB(rs.db, []byte(prefix))
	}

	switch params.typ {
	case types.StoreTypeMulti:
		panic("recursive MultiStores not yet supported")
	case types.StoreTypeDB:
		return commitDBStoreAdapter{Store: dbadapter.Store{DB: db}}, nil
	case types.StoreTypeTransient:
		_, ok := key.(*types.TransientStoreKey)
		if !ok {
			return nil, fmt.Errorf("invalid StoreKey for StoreTypeTransient: %s", key.String())
		}

		return transient.NewStore(), nil

	case types.StoreTypeMemory:
		if _, ok := key.(*types.MemoryStoreKey); !ok {
			return nil, fmt.Errorf("unexpected key type for a MemoryStoreKey; got: %s", key.String())
		}

		return mem.NewStore(), nil

	default:
		panic(fmt.Sprintf("unrecognized store type %v", params.typ))
	}
}

// SetTracer sets the tracer for the MultiStore that the underlying
// stores will utilize to trace operations. A MultiStore is returned.
func (rs *Store) SetTracer(w io.Writer) *Store {
	rs.traceWriter = w
	return rs
}

// SetTracingContext updates the tracing context for the MultiStore by merging
// the given context with the existing context by key. Any existing keys will
// be overwritten. It is implied that the caller should update the context when
// necessary between tracing operations. It returns a modified MultiStore.
func (rs *Store) SetTracingContext(tc types.TraceContext) *Store {
	if rs.traceContext != nil {
		for k, v := range tc {
			rs.traceContext[k] = v
		}
	} else {
		rs.traceContext = tc
	}

	return rs
}

// TracingEnabled returns if tracing is enabled for the MultiStore.
func (rs *Store) TracingEnabled() bool {
	return rs.traceWriter != nil
}

// CacheWrap implements CacheWrapper/Store/CommitStore.
func (rs *Store) CacheWrap() types.CacheWrap {
	return rs.CacheMultiStore().(types.CacheWrap)
}

//----------------------------------------
// +MultiStore

// CacheMultiStore cache-wraps the multi-store and returns a CacheMultiStore.
// It implements the MultiStore interface.
func (rs *Store) CacheMultiStore() types.CacheMultiStore {
	stores := make(map[types.StoreKey]types.CacheWrapper)
	for k, v := range rs.stores {
		stores[k] = v
	}

	return cachemulti.NewStore(rs.db, stores, rs.keysByName, rs.traceWriter, rs.traceContext)
}

// GetStore returns a mounted Store for a given StoreKey. If the StoreKey does
// not exist, it will panic. If the Store is wrapped in an inter-block cache, it
// will be unwrapped prior to being returned.
//
// TODO: This isn't used directly upstream. Consider returning the Store as-is
// instead of unwrapping.
func (rs *Store) GetStore(key types.StoreKey) types.Store {
	store := rs.GetCommitKVStore(key)
	if store == nil {
		panic(fmt.Sprintf("store does not exist for key: %s", key.Name()))
	}

	return store
}

// GetKVStore returns a mounted KVStore for a given StoreKey. If tracing is
// enabled on the KVStore, a wrapped TraceKVStore will be returned with the root
// store's tracer, otherwise, the original KVStore will be returned.
//
// NOTE: The returned KVStore may be wrapped in an inter-block cache if it is
// set on the root store.
func (rs *Store) GetKVStore(key types.StoreKey) types.KVStore {
	store := rs.stores[key].(types.KVStore)

	if rs.TracingEnabled() {
		store = tracekv.NewStore(store, rs.traceWriter, rs.traceContext)
	}

	return store
}

// getStoreByName performs a lookup of a StoreKey given a store name typically
// provided in a path. The StoreKey is then used to perform a lookup and return
// a Store. If the Store is wrapped in an inter-block cache, it will be unwrapped
// prior to being returned. If the StoreKey does not exist, nil is returned.
func (rs *Store) getStoreByName(name string) types.Store {
	key := rs.keysByName[name]
	if key == nil {
		return nil
	}

	return rs.GetCommitKVStore(key)
}

//----------------------------------------
// storeParams

type storeParams struct {
	key types.StoreKey
	db  dbm.DB
	typ types.StoreType
}
