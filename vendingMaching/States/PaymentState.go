package States

import (
	"fmt"
	"lld/vendingMaching"
)
import "lld/vendingMaching/Machine"

type PaymentState struct {
	machine *Machine.VendingMachine
}

func GetPaymentState(machine *Machine.VendingMachine) vendingMaching.State {
	return &PaymentState{machine: machine}
}

func (s *PaymentState) InsertCoin(coin int) {
	if s.machine.SelectedProduct.Price <= coin {
		fmt.Printf("Rs.%d recieved for %s\n", coin, s.machine.SelectedProduct.Name)
		s.machine.SetState(s.machine.States[Machine.DispenseProductState])
		s.machine.DispenseProduct()
	} else {
		fmt.Printf("(Rs.%d) Insufficient Payment, Rs.%d required for %s\n", coin, s.machine.SelectedProduct.Price, s.machine.SelectedProduct.Name)
	}
}

func (s *PaymentState) SelectProduct(sku int) {
	s.machine.SetState(GetSelectProductState(s.machine))
	s.machine.SelectProduct(sku)
}

func (s *PaymentState) DispenseProduct() {
	fmt.Println("Please complete the payment first")
}
