package States

type State interface {
	SelectProduct(sku int)
	InsertCoin(coin int)
	DispenseProduct()
}
