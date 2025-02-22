package service

import (
	"errors"
	"fmt"
	"lld/inventoryService/entity"
	"lld/inventoryService/repo"
	"sync"
	"time"
)

type IOrderService interface {
	PlaceOrder(order entity.Order) (int, error)
	ConfirmOrder(orderId int) error
	CancelOrder(orderID int) error
	GetOrder(orderID int) (entity.Order, error)
	GetQuantity(itemID int) (int, error)
}

type OrderService struct {
	orderRepo     repo.IOrderRepo
	inventoryRepo repo.IInventoryRepo
	mutex         sync.Mutex
}

func (o *OrderService) GetQuantity(itemID int) (int, error) {
	return o.inventoryRepo.GetQuantity(itemID)
}

func (o *OrderService) PlaceOrder(order entity.Order) (int, error) {
	order.Time = time.Now()
	for _, orderItem := range order.OrderItems {
		quantity, err := o.inventoryRepo.GetQuantity(orderItem.ProductID)
		if err != nil {
			return 0, err
		}

		if quantity < orderItem.Quantity {
			return 0, errors.New(fmt.Sprintf("insufficient quantity for item %d", orderItem.ProductID))
		}

	}

	for _, orderItem := range order.OrderItems {
		err := o.inventoryRepo.BlockQuantity(orderItem.ProductID, orderItem.Quantity)
		if err != nil {
			return 0, err
		}
	}

	orderId, err := o.orderRepo.AddOrder(&order)
	if err != nil {
		return 0, err
	}

	order.Timer = time.AfterFunc(time.Second, func() {
		o.checkExpiry(orderId)
	})
	return orderId, nil
}

func (o *OrderService) checkExpiry(orderId int) {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	qty, err := o.orderRepo.UnblockExpiredOrder(orderId)
	if err != nil {
		return
	}
	for productID, quantity := range qty {
		err = o.inventoryRepo.AddQuantity(productID, quantity)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (o *OrderService) ConfirmOrder(orderId int) error {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	order, err := o.orderRepo.GetOrder(orderId)
	if err != nil {
		return err
	}
	if order.OrderStatus != entity.OrderPlaced {
		return errors.New("order cannot be completed, invalid status")
	}
	err = o.orderRepo.UpdateStatus(orderId, entity.OrderCompleted)
	if err != nil {
		return err
	}
	order.Timer.Stop()
	return nil
}

func (o *OrderService) CancelOrder(orderID int) error {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	order, err := o.orderRepo.GetOrder(orderID)
	if err != nil {
		return err
	}

	err = o.orderRepo.UpdateStatus(orderID, entity.OrderCancelled)
	if err != nil {
		return err
	}

	for _, orderItem := range order.OrderItems {
		err = o.inventoryRepo.AddQuantity(orderItem.ProductID, orderItem.Quantity)
		if err != nil {
			return err
		}
	}

	return nil
}

func (o *OrderService) GetOrder(orderID int) (entity.Order, error) {
	order, _ := o.orderRepo.GetOrder(orderID)
	return *order, nil
}

func NewOrderService(or repo.IOrderRepo, ir repo.IInventoryRepo) IOrderService {
	return &OrderService{or, ir, sync.Mutex{}}
}
