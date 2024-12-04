package vendingMaching

func main() {
	vm := VendingMachine{}
	vm.State = getIdleState(&vm)

	vm.inventory = make(map[int]Product)
	vm.inventory[1] = Product{Sku: 1, Price: 10, Name: "Samosa"}
	vm.inventory[2] = Product{Sku: 1, Price: 20, Name: "Kachori"}

	vm.State.SelectProduct(1)
	vm.State.InsertCoin(15)
	vm.State.DispenseProduct()

	vm.State.SelectProduct(2)
	vm.State.InsertCoin(15)
	vm.State.DispenseProduct()
}
