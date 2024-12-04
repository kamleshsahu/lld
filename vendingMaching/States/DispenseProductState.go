package States

import (
	"fmt"
	"lld/vendingMaching"
	"lld/vendingMaching/Machine"
)

type DispenseProductState struct {
	machine *Machine.VendingMachine
}

func getDispenseProductState(machine *Machine.VendingMachine) vendingMaching.State {
	return &DispenseProductState{machine: machine}
}

func (d *DispenseProductState) InsertCoin(coin int) {
	//TODO implement me
	fmt.Println("First Dispense your Product")
}

func (d *DispenseProductState) SelectProduct(sku int) {
	//TODO implement me
	fmt.Println("First Dispense your Product")
}

func (d *DispenseProductState) DispenseProduct() {
	fmt.Printf("%s dispensed, please collect\n", d.machine.SelectedProduct.Name)
	d.machine.SelectedProduct = nil
	d.machine.Payment = 0
	d.machine.State = GetSelectProductState(d.machine)
}
