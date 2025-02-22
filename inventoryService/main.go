package main

import (
	"fmt"
	"lld/inventoryService/entity"
	"lld/inventoryService/repo"
	"lld/inventoryService/service"
	"time"
)

func main() {

	ir := repo.NewInventoryRepo()
	or := repo.NewOrderRepo()
	os := service.NewOrderService(or, ir)

	p1 := entity.Product{ID: 1, Name: "Product 1", Quantity: 10}
	ir.AddProduct(p1)

	order := entity.Order{
		OrderItems:  []entity.OrderItem{{ProductID: 1, Quantity: 1, Name: "Product 1"}},
		OrderStatus: 0,
	}
	placedOrderId, err := os.PlaceOrder(order)
	if err != nil {
		return
	}
	fmt.Println(placedOrderId)
	os.ConfirmOrder(placedOrderId)
	qty, _ := os.GetQuantity(1)
	fmt.Println(qty)

	order2 := entity.Order{
		OrderItems:  []entity.OrderItem{{ProductID: 1, Quantity: 5, Name: "Product 1"}},
		OrderStatus: 0,
	}
	_, err = os.PlaceOrder(order2)
	if err != nil {
		return
	}
	qty, _ = os.GetQuantity(1)
	fmt.Println(qty)

	time.Sleep(2 * time.Second)
	qty, _ = os.GetQuantity(1)
	fmt.Println(qty)

}
