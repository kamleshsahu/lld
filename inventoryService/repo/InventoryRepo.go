package repo

import (
	"errors"
	"lld/inventoryService/entity"
)

type inventoryRepo struct {
	productId  int
	productMap map[int]entity.Product
}

func (i *inventoryRepo) GetBlockedQuantity(itemID int) (int, error) {
	if _, ok := i.productMap[itemID]; !ok {
		return 0, errors.New("product not found")
	}
	return i.productMap[itemID].BlockedQuantity, nil
}

func (i *inventoryRepo) UnblockBlockedQuantity(itemID int, quantity int) error {
	if _, ok := i.productMap[itemID]; !ok {
		return errors.New("product not found")
	}
	product := i.productMap[itemID]
	product.Quantity += quantity
	product.BlockedQuantity -= quantity
	i.productMap[itemID] = product
	return nil
}

func (i *inventoryRepo) BlockQuantity(itemID int, quantity int) error {
	if _, ok := i.productMap[itemID]; !ok {
		return errors.New("product not found")
	}
	product := i.productMap[itemID]
	product.Quantity -= quantity
	product.BlockedQuantity += quantity
	i.productMap[itemID] = product
	return nil
}

func (i *inventoryRepo) AddProduct(product entity.Product) (int, error) {
	if _, ok := i.productMap[product.ID]; ok {
		i.productMap[product.ID] = product
		return product.ID, nil
	}
	i.productId++
	i.productMap[i.productId] = product
	return i.productId, nil
}

func (i *inventoryRepo) AddQuantity(itemID int, quantity int) error {
	if _, ok := i.productMap[itemID]; !ok {
		return errors.New("product not found")
	}
	product := i.productMap[itemID]
	product.Quantity += quantity
	i.productMap[itemID] = product
	return nil
}

func (i *inventoryRepo) ConsumeBlockedQuantity(itemID int, quantity int) error {
	if _, ok := i.productMap[itemID]; !ok {
		return errors.New("product not found")
	}
	product := i.productMap[itemID]
	product.BlockedQuantity -= quantity
	i.productMap[itemID] = product
	return nil
}

func (i *inventoryRepo) GetQuantity(itemID int) (int, error) {
	if _, ok := i.productMap[itemID]; !ok {
		return 0, errors.New("product not found")
	}
	return i.productMap[itemID].Quantity, nil
}

func NewInventoryRepo() IInventoryRepo {
	return &inventoryRepo{productMap: make(map[int]entity.Product)}
}
