package main

import (
	"lld/vendingMaching/Machine"
	"lld/vendingMaching/States"
)

func main() {

	Inventory := make(map[int]Machine.Product)
	Inventory[1] = Machine.Product{Sku: 1, Price: 10, Name: "Samosa"}
	Inventory[2] = Machine.Product{Sku: 1, Price: 20, Name: "Kachori"}
	vm := Machine.VendingMachine{}

	vm.State = States.GetSelectProductState(&vm)
	vm.Inventory = Inventory

	vm.State.SelectProduct(1)
	vm.State.InsertCoin(15)
	vm.State.DispenseProduct()

	vm.State.SelectProduct(2)
	vm.State.SelectProduct(1)
	vm.State.InsertCoin(100)
	vm.State.DispenseProduct()

	vm.State.SelectProduct(2)
	vm.State.InsertCoin(10)
	vm.State.InsertCoin(20)
	vm.State.DispenseProduct()

	vm.State.SelectProduct(3)

}
