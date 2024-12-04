package States

import (
	"fmt"
	"lld/vendingMaching"
	"lld/vendingMaching/Machine"
)

type SelectProductState struct {
	machine *Machine.VendingMachine
}

func GetSelectProductState(machine *Machine.VendingMachine) vendingMaching.State {
	return &SelectProductState{machine: machine}
}

func (i *SelectProductState) InsertCoin(coin int) {
	fmt.Println("First select a product")
}

func (i *SelectProductState) SelectProduct(sku int) {
	if product, found := i.machine.Inventory[sku]; found {
		i.machine.SelectedProduct = &product
		i.machine.State = getPaymentState(i.machine)
		fmt.Printf("%s selected\n", product.Name)
		return
	}
	fmt.Println("SKU not found")
}

func (i *SelectProductState) DispenseProduct() {
	fmt.Println("First Select a Product")
}
