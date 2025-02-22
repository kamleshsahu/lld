package repo

import (
	"errors"
	"lld/inventoryService/entity"
)

type orderRepo struct {
	orderId  int
	ordermap map[int]entity.Order
}

func (o *orderRepo) GetOrder(orderId int) (entity.Order, error) {
	if _, ok := o.ordermap[orderId]; !ok {
		return entity.Order{}, errors.New("order not found")
	}
	return o.ordermap[orderId], nil
}

func (o *orderRepo) AddOrder(order entity.Order) (int, error) {
	o.orderId++
	o.ordermap[o.orderId] = order
	return o.orderId, nil
}

func (o *orderRepo) UpdateStatus(orderId int, status entity.OrderStatus) error {
	if _, ok := o.ordermap[orderId]; !ok {
		return errors.New("order not found")
	}
	order := o.ordermap[orderId]
	order.OrderStatus = status
	return nil
}

func NewOrderRepo() IOrderRepo {
	return &orderRepo{ordermap: make(map[int]entity.Order)}
}
