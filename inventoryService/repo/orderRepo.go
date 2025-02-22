package repo

import (
	"errors"
	"lld/inventoryService/entity"
	"sync"
)

type IOrderRepo interface {
	AddOrder(order *entity.Order) (int, error)
	UpdateStatus(orderId int, status entity.OrderStatus) error
	GetOrder(orderId int) (*entity.Order, error)
	UnblockExpiredOrder(orderId int) (map[int]int, error)
}

type orderRepo struct {
	orderId  int
	ordermap map[int]*entity.Order
	mutex    sync.Mutex
}

func (o *orderRepo) UnblockExpiredOrder(orderId int) (map[int]int, error) {
	qty := make(map[int]int)
	o.mutex.Lock()
	defer o.mutex.Unlock()

	if _, ok := o.ordermap[orderId]; !ok {
		return nil, errors.New("order not found")
	}
	if o.ordermap[orderId].OrderStatus == entity.OrderPlaced {
		for _, orderItem := range o.ordermap[orderId].OrderItems {
			qty[orderItem.ProductID] += orderItem.Quantity
		}
		order := o.ordermap[orderId]
		order.OrderStatus = entity.OrderCancelled
	}

	return qty, nil
}

func (o *orderRepo) GetOrder(orderId int) (*entity.Order, error) {
	if _, ok := o.ordermap[orderId]; !ok {
		return nil, errors.New("order not found")
	}
	return o.ordermap[orderId], nil
}

func (o *orderRepo) AddOrder(order *entity.Order) (int, error) {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	o.orderId++
	order.ID = o.orderId
	o.ordermap[o.orderId] = order
	return o.orderId, nil
}

func (o *orderRepo) UpdateStatus(orderId int, status entity.OrderStatus) error {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	if _, ok := o.ordermap[orderId]; !ok {
		return errors.New("order not found")
	}
	order := o.ordermap[orderId]
	order.OrderStatus = status
	return nil
}

func NewOrderRepo() IOrderRepo {
	return &orderRepo{ordermap: make(map[int]*entity.Order)}
}
