package domain

type OrderDetail struct {
	ID        int
	OrderID   int // shadow field
	ProductID int // shadow field
	Quantity  int
	Price     float32
}

func NewOrderDetail(quantity int, OrderId int, productId int, price float32) *OrderDetails {
	return &OrderDetails{
		OrderID:   OrderId,
		ProductID: productId,
		Quantity:  quantity,
		Price:     price,
	}
}
