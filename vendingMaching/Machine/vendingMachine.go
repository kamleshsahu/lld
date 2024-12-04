package Machine

import "fmt"
import "States"

type VendingMachine struct {
	SelectedProduct *Product
	payment         int
	State           States.
	inventory       map[int]Product
}

type Product struct {
	Sku   int
	Name  string
	Price int
}

type SelectProductState struct {
	machine *VendingMachine
}



func (s *SelectProductState) InsertCoin(coin int) {
	if s.machine.SelectedProduct.Price <= coin {
		s.machine.State = getDispenseProductState(s.machine)
	} else {
		fmt.Println("insufficient payment")
	}
}

func (s *SelectProductState) SelectProduct(sku int) {
	//TODO implement me
	fmt.Println("implement me")
}

func (s *SelectProductState) DispenseProduct() {
	//TODO implement me
	fmt.Println("implement me")
}

func (i *IdleState) InsertCoin(coin int) {
	panic("first select product")
}

func (i *IdleState) SelectProduct(sku int) {
	if product, found := i.machine.inventory[sku]; found {
		i.machine.SelectedProduct = &product
		i.machine.State = getSelectProductState(i.machine)
		return
	}
	fmt.Println("sku not found")
}

func (i *IdleState) DispenseProduct() {
	fmt.Println("implement me")
}

func getIdleState(machine *VendingMachine) State {
	return &IdleState{machine: machine}
}

func getSelectProductState(machine *VendingMachine) State {
	return &SelectProductState{machine: machine}
}

func getDispenseProductState(machine *VendingMachine) State {
	return &DispenseProductState{machine: machine}
}


