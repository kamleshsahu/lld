package main

import (
	"lld/vendingMaching"
	"lld/vendingMaching/Machine"
	"lld/vendingMaching/States"
)

func main() {

	Inventory := make(map[int]Machine.Product)
	Inventory[1] = Machine.Product{Sku: 1, Price: 10, Name: "Samosa"}
	Inventory[2] = Machine.Product{Sku: 1, Price: 20, Name: "Kachori"}
	vm := Machine.VendingMachine{}

	vm.States = map[int]vendingMaching.State{
		Machine.SelectProductState:   States.GetSelectProductState(&vm),
		Machine.PaymentState:         States.GetPaymentState(&vm),
		Machine.DispenseProductState: States.GetDispenseProductState(&vm),
	}

	vm.SetState(States.GetSelectProductState(&vm))
	vm.Inventory = Inventory

	vm.SelectProduct(1)
	vm.InsertCoin(15)

	vm.SelectProduct(2)
	vm.SelectProduct(1)
	vm.InsertCoin(100)

	vm.SelectProduct(2)
	vm.InsertCoin(10)
	vm.InsertCoin(20)

	vm.SelectProduct(3)
}
