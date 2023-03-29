package entity

import (
	"github.com/google/uuid"
)

type Products struct {
	ID          uuid.UUID
	ProductName string // shadow field
	Category    string
	Price       float32
}

func NewProducts(productName string, price float32, category string) *Products {
	return &Products{
		ID:          uuid.New(),
		ProductName: productName,
		Category:    category,
		Price:       price,
	}
}
