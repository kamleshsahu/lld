package entity

import "sync"

type Transaction struct {
	Store
	transactionId int
	Logs          []Pair
	lock          *sync.Mutex
}

func (t *Transaction) GET(key string) Pair {
	t.lock.Lock()
	defer t.lock.Unlock()
	return t.Store[key]
}

func (t *Transaction) PUT(key string, value string) {
	t.lock.Lock()
	defer t.lock.Unlock()

	pair := t.Store[key]
	pair.Value = value
	t.Logs = append(t.Logs, Pair{key, value, pair.Version})
	pair.Version++
	t.Store[key] = pair
}

func (t *Transaction) DELETE(key string) {
	t.lock.Lock()
	defer t.lock.Unlock()

	pair := t.Store[key]
	t.Logs = append(t.Logs, Pair{key, "", pair.Version})
	delete(t.Store, key)
}

func NewTransaction(store Store, transactionId int) *Transaction {
	return &Transaction{
		store.SnapShot(),
		transactionId,
		make([]Pair, 0),
		&sync.Mutex{},
	}
}
