package entity

import (
	"github.com/google/uuid"
)

type OrderDetails struct {
	ID       uuid.UUID
	OrderID  string      // shadow field
	Products []*Products // shadow field
	Quantity int
	Price    float32
}

func NewOrderDetails(quantity int, price float32) *OrderDetails {
	return &OrderDetails{
		ID:       uuid.New(),
		Quantity: quantity,
		Price:    price,
	}
}
