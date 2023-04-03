package entity

type Products struct {
	ProductName string // shadow field
	Category    string
	Price       float32
}

func NewProducts(productName string, price float32, category string) *Products {
	return &Products{
		ProductName: productName,
		Category:    category,
		Price:       price,
	}
}
