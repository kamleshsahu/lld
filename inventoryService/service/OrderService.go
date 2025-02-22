package service

import (
	"errors"
	"fmt"
	"lld/inventoryService/entity"
	"lld/inventoryService/repo"
)

type IOrderService interface {
	PlaceOrder(order entity.Order) (int, error)
	ConfirmOrder(orderId int) error
	CancelOrder(orderID int) error
	GetOrder(orderID int) (entity.Order, error)
}

type OrderService struct {
	orderRepo     repo.IOrderRepo
	inventoryRepo repo.IInventoryRepo
}

func (o *OrderService) PlaceOrder(order entity.Order) (int, error) {

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

	orderId, err := o.orderRepo.AddOrder(order)
	if err != nil {
		return 0, err
	}
	return orderId, nil
}

func (o *OrderService) ConfirmOrder(orderId int) error {
	order, err := o.orderRepo.GetOrder(orderId)
	if err != nil {
		return err
	}
	for _, orderItem := range order.OrderItems {
		err = o.inventoryRepo.ConsumeBlockedQuantity(orderItem.ProductID, orderItem.Quantity)
		if err != nil {
			return err
		}
	}

	err = o.orderRepo.UpdateStatus(orderId, entity.OrderCompleted)
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderService) CancelOrder(orderID int) error {
	order, err := o.orderRepo.GetOrder(orderID)
	if err != nil {
		return err
	}
	for _, orderItem := range order.OrderItems {
		err = o.inventoryRepo.UnblockBlockedQuantity(orderItem.ProductID, orderItem.Quantity)
		if err != nil {
			return err
		}
	}

	err = o.orderRepo.UpdateStatus(orderID, entity.OrderCancelled)
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderService) GetOrder(orderID int) (entity.Order, error) {
	return o.orderRepo.GetOrder(orderID)
}

func NewOrderService(or repo.IOrderRepo, ir repo.IInventoryRepo) IOrderService {
	return &OrderService{or, ir}
}
