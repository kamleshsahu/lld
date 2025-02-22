package main

import (
	"fmt"
	"lld/inventoryService/entity"
	"lld/inventoryService/repo"
	"lld/inventoryService/service"
)

func main() {

	ir := repo.NewInventoryRepo()
	or := repo.NewOrderRepo()
	os := service.NewOrderService(or, ir)

	p1 := entity.Product{ID: 1, Name: "Product 1", Quantity: 10, BlockedQuantity: 0}
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
	qty, _ := ir.GetBlockedQuantity(1)
	fmt.Println(qty)
	os.ConfirmOrder(placedOrderId)
	qty, _ = ir.GetBlockedQuantity(1)
	fmt.Println(qty)
	qty, _ = ir.GetQuantity(1)

	fmt.Println(qty)
}
