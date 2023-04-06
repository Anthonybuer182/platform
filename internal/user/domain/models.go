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

// type OrderDto struct {
// 	Id          int     `json:"id"`
// 	ProductId   int     `json:"productId"`
// 	PruductName string  `json:"pruductName"`
// 	Type        int     `json:"type"`
// 	Price       float64 `json:"price"`
// }

// rpc调用order
type PlaceOrderModel struct {
	users []string
}

type OrderDto struct {
	Id          string        `json:"id"`
	OrderNum    string        `json:"orderNum"`
	OrderStatus string        `json:"orderStatus"`
	DetailsDto  []*DetailsDto `json:"detailsDto"`
	UserDto     *UserDto      `json:"userDto"`
}
type DetailsDto struct {
	Id         string
	ProductDto ProductDto
	Quantity   int32
	Amount     float64
}
type ProductDto struct {
	ProductName string
	Category    string
	Price       float32
}
type UserDto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
