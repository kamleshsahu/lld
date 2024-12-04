package States

import "fmt"

type DispenseProductState struct {
	machine *VendingMachine
}

func (d *DispenseProductState) InsertCoin(coin int) {
	//TODO implement me
	fmt.Println("implement me")
}

func (d *DispenseProductState) SelectProduct(sku int) {
	//TODO implement me
	fmt.Println("implement me")
}

func (d *DispenseProductState) DispenseProduct() {
	fmt.Printf("product rolled out %s\n", d.machine.SelectedProduct.Name)
	d.machine.SelectedProduct = nil
	d.machine.payment = 0
	d.machine.State = getIdleState(d.machine)
}
