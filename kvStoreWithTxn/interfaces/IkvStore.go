package interfaces

import "lld/kvStoreWithTxn/entity"

type IKVStore interface {
	GET(key string) (string, error)
	PUT(key string, value string) error
	DELETE(key string) error
	GetStore() *entity.Store
	SetStore(store *entity.Store)
}
