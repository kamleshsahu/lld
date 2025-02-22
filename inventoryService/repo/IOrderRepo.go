package repo

import "lld/inventoryService/entity"

type IOrderRepo interface {
	AddOrder(order entity.Order) (int, error)
	UpdateStatus(orderId int, status entity.OrderStatus) error
	GetOrder(orderId int) (entity.Order, error)
}
