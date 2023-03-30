package domain

type ItemTypeDto struct {
	Name  string  `json:"name"`
	Type  int     `json:"type"`
	Price float64 `json:"price"`
	Image string  `json:"image"`
}

type ItemDto struct {
	Price float64 `json:"price"`
	Type  int     `json:"type"`
}
type OrderDto struct {
	Id          int     `json:"id"`
	ProductId   int     `json:"productId"`
	PruductName string  `json:"pruductName"`
	Type        int     `json:"type"`
	Price       float64 `json:"price"`
}

// rpc调用order
type PlaceOrderModel struct {
	users []string
}

type ItemModel struct {
	Id       string
	OrderNum string
}
