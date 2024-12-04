package Machine

import (
	"lld/vendingMaching"
)

const (
	SelectProductState = iota
	PaymentState
	DispenseProductState
)

type VendingMachine struct {
	SelectedProduct *Product
	Payment         int
	state           vendingMaching.State
	Inventory       map[int]Product
	States          map[int]vendingMaching.State
}

type Product struct {
	Sku   int
	Name  string
	Price int
	Stock int
}

func (v *VendingMachine) SelectProduct(sku int) {
	v.state.SelectProduct(sku)
}

func (v *VendingMachine) InsertCoin(coin int) {
	v.state.InsertCoin(coin)
}

func (v *VendingMachine) DispenseProduct() {
	v.state.DispenseProduct()
}

func (v *VendingMachine) SetState(state vendingMaching.State) {
	v.state = state
}
