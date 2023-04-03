package domain

import (
	"github.com/google/uuid"
)

type OrderDetails struct {
	ID        uuid.UUID
	OrderID   int // shadow field
	ProductID int // shadow field
	Quantity  int
	Price     float32
}

func NewOrderDetails(quantity int, OrderId int, productId int, price float32) *OrderDetails {
	return &OrderDetails{
		ID:        uuid.New(),
		OrderID:   OrderId,
		ProductID: productId,
		Quantity:  quantity,
		Price:     price,
	}
}
