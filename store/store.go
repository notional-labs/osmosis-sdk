package store

import (
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/osmosis-sdk/store/cache"
	"github.com/cosmos/osmosis-sdk/store/rootmulti"
	"github.com/cosmos/osmosis-sdk/store/types"
)

func NewCommitMultiStore(db dbm.DB) types.CommitMultiStore {
	return rootmulti.NewStore(db)
}

func NewCommitKVStoreCacheManager() types.MultiStorePersistentCache {
	return cache.NewCommitKVStoreCacheManager(cache.DefaultCommitKVStoreCacheSize)
}
