package service

import (
	"fmt"
	"lld/kvStoreWithTxn/entity"
	"lld/kvStoreWithTxn/interfaces"
	"sync"
)

type transactionManager struct {
	storeManager  interfaces.IKVStore
	transactions  map[int]*entity.Transaction
	transactionId int
	lock          *sync.Mutex
}

func (k *transactionManager) GET(transactionId int, key string) (string, error) {
	k.lock.Lock()
	defer k.lock.Unlock()
	transaction := k.transactions[transactionId]
	return transaction.GET(key).Value, nil
}

func (k *transactionManager) PUT(transactionId int, key string, value string) error {
	k.lock.Lock()
	defer k.lock.Unlock()
	txn := k.transactions[transactionId]
	txn.PUT(key, value)
	return nil
}

func (k *transactionManager) DELETE(transactionId int, key string) error {
	k.lock.Lock()
	defer k.lock.Unlock()
	k.transactions[transactionId].DELETE(key)
	return nil
}

func (k *transactionManager) BEGIN() (int, error) {
	k.lock.Lock()
	defer k.lock.Unlock()
	k.transactionId++
	k.transactions[k.transactionId] = entity.NewTransaction(*k.storeManager.GetStore(), k.transactionId)
	return k.transactionId, nil
}

func (k *transactionManager) COMMIT(transactionId int) error {
	k.lock.Lock()
	defer k.lock.Unlock()
	storeInstance := k.storeManager.GetStore()
	snapShot := (*storeInstance).SnapShot()
	txn := k.transactions[transactionId]
	for _, log := range txn.Logs {
		pair, found := (*storeInstance)[log.Key]
		if found && pair.Version != log.Version {
			storeInstance = &snapShot
			return fmt.Errorf("transaction %d failed, version mismatch", transactionId)
		}
		pair.Version++
		pair.Value = log.Value
		(*storeInstance)[log.Key] = pair
	}
	k.storeManager.SetStore(storeInstance)
	err := k.ROLLBACK(transactionId)
	if err != nil {
		return err
	}
	return nil
}

func (k *transactionManager) ROLLBACK(transactionId int) error {
	delete(k.transactions, transactionId)
	return nil
}

func NewTxn(store interfaces.IKVStore) interfaces.ITransaction {
	return &transactionManager{
		storeManager: store,
		transactions: make(map[int]*entity.Transaction),
		lock:         &sync.Mutex{},
	}
}
