package Machine

import (
	"lld/vendingMaching"
)

type VendingMachine struct {
	SelectedProduct *Product
	Payment         int
	State           vendingMaching.State
	Inventory       map[int]Product
}

type Product struct {
	Sku   int
	Name  string
	Price int
}
