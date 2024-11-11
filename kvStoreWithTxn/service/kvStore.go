package service

import (
	"lld/kvStoreWithTxn/entity"
	"lld/kvStoreWithTxn/interfaces"
	"sync"
)

type kvStore struct {
	Store entity.Store
	lock  *sync.Mutex
}

func (k *kvStore) SetStore(store *entity.Store) {
	k.lock.Lock()
	defer k.lock.Unlock()
	k.Store = *store
}

func (k *kvStore) GetStore() *entity.Store {
	return &k.Store
}

func (k *kvStore) DELETE(key string) error {
	k.lock.Lock()
	defer k.lock.Unlock()
	delete(k.Store, key)
	return nil
}

func (k *kvStore) GET(key string) (string, error) {
	k.lock.Lock()
	defer k.lock.Unlock()
	return k.Store[key].Value, nil
}

func (k *kvStore) PUT(key string, value string) error {
	k.lock.Lock()
	defer k.lock.Unlock()
	pair := k.Store[key]
	pair.Value = value
	k.Store[key] = pair
	return nil
}

var (
	store     interfaces.IKVStore
	singleton sync.Once
)

func GetKVStore() interfaces.IKVStore {
	singleton.Do(func() {
		store = &kvStore{Store: make(entity.Store), lock: new(sync.Mutex)}
	})
	return store
}
