package repo

import (
	"errors"
	"lld/inventoryService/entity"
	"sync"
)

type IInventoryRepo interface {
	AddProduct(product entity.Product) (int, error)
	AddQuantity(itemID int, quantity int) error
	BlockQuantity(itemID int, quantity int) error
	GetQuantity(itemID int) (int, error)
}

type inventoryRepo struct {
	nextId     int
	productMap map[int]entity.Product
	mutex      sync.Mutex
}

func (i *inventoryRepo) BlockQuantity(itemID int, quantity int) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	if _, ok := i.productMap[itemID]; !ok {
		return errors.New("product not found")
	}
	product := i.productMap[itemID]
	product.Quantity -= quantity
	i.productMap[itemID] = product
	return nil
}

func (i *inventoryRepo) AddProduct(product entity.Product) (int, error) {
	if _, ok := i.productMap[product.ID]; ok {
		i.productMap[product.ID] = product
		return product.ID, nil
	}
	i.mutex.Lock()
	i.nextId++
	i.mutex.Unlock()

	i.productMap[i.nextId] = product
	return i.nextId, nil
}

func (i *inventoryRepo) AddQuantity(itemID int, quantity int) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	if _, ok := i.productMap[itemID]; !ok {
		return errors.New("product not found")
	}
	product := i.productMap[itemID]
	product.Quantity += quantity
	i.productMap[itemID] = product
	return nil
}

func (i *inventoryRepo) GetQuantity(itemID int) (int, error) {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	if _, ok := i.productMap[itemID]; !ok {
		return 0, errors.New("product not found")
	}
	return i.productMap[itemID].Quantity, nil
}

func NewInventoryRepo() IInventoryRepo {
	return &inventoryRepo{productMap: make(map[int]entity.Product), nextId: 0, mutex: sync.Mutex{}}
}
