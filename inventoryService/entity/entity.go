package entity

type Product struct {
	ID              int
	Name            string
	Quantity        int
	BlockedQuantity int
}

type OrderItem struct {
	ProductID int
	Quantity  int
	Name      string
}

type OrderStatus int

const (
	OrderPlaced OrderStatus = iota
	OrderShipped
	OrderDelivered
	OrderCancelled
	OrderCompleted
)

type Order struct {
	ID          int
	OrderItems  []OrderItem
	OrderStatus OrderStatus
}
