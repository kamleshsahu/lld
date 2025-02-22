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
	go func() {
		placedOrderId, err := os.PlaceOrder(order)
		if err != nil {
			fmt.Println("placed status : ", placedOrderId, err)
		}
		fmt.Println("orderID:", placedOrderId)
		err = os.ConfirmOrder(placedOrderId)
		if err != nil {
			fmt.Println("confirmed status : ", placedOrderId, err)
		}
		qty, _ := os.GetQuantity(1)
		fmt.Println(qty)
	}()

	go func() {
		order2 := entity.Order{
			OrderItems:  []entity.OrderItem{{ProductID: 1, Quantity: 5, Name: "Product 1"}},
			OrderStatus: 0,
		}
		placedOrderId, err := os.PlaceOrder(order2)
		if err != nil {
			fmt.Println("placed status : ", placedOrderId, err)
		}
		fmt.Println("orderID:", placedOrderId)
		err = os.ConfirmOrder(placedOrderId)
		if err != nil {
			fmt.Println("confirmed status : ", placedOrderId, err)
		}
		fmt.Println("orderID:", placedOrderId)

		qty, _ := os.GetQuantity(1)
		fmt.Println(qty)
	}()

	go func() {
		time.Sleep(1 * time.Second)
		order3 := entity.Order{
			OrderItems:  []entity.OrderItem{{ProductID: 1, Quantity: 6, Name: "Product 1"}},
			OrderStatus: 0,
		}
		placedOrderId, err := os.PlaceOrder(order3)
		if err != nil {
			fmt.Println("placed status : ", placedOrderId, err)
		}
		fmt.Println("orderID:", placedOrderId)
		qty, _ := os.GetQuantity(1)
		fmt.Println(qty)
	}()

	time.Sleep(3 * time.Second)
	qty, _ := os.GetQuantity(1)
	fmt.Println(qty)

}
