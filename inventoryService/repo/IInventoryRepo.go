package repo

import "lld/inventoryService/entity"

type IInventoryRepo interface {
	AddProduct(product entity.Product) (int, error)
	AddQuantity(itemID int, quantity int) error
	BlockQuantity(itemID int, quantity int) error
	ConsumeBlockedQuantity(itemID int, quantity int) error
	UnblockBlockedQuantity(itemID int, quantity int) error
	GetQuantity(itemID int) (int, error)
	GetBlockedQuantity(itemID int) (int, error)
}
