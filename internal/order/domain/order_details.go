package domain

import (
	"context"
	"golang.org/x/exp/slog"
	"strconv"
)

type OrderDetail struct {
	ID       int
	OrderID  int       // shadow field
	Product  *Products // shadow field
	Quantity int
	Price    float32
}

func NewOrderDetail(quantity int, OrderId int, prodcut *Products, price float32) *OrderDetail {
	return &OrderDetail{
		OrderID:  OrderId,
		Product:  prodcut,
		Quantity: quantity,
		Price:    price,
	}
}

func OrderDetailAggregate(ctx context.Context, order *OrderDetails, productDomainService ProductDomainService) *OrderDetail {

	//通过grpc查询用户服务
	products, err := productDomainService.GetProductLists(ctx, &OrderModel{
		ProductId: strconv.Itoa(order.ProductID),
	})

	orderDetail := &OrderDetail{
		OrderID:  order.OrderID,
		Quantity: order.Quantity,
		Price:    order.Price,
	}
	if err == nil && len(products) > 0 {
		orderDetail.Product = products[0]
	}
	slog.Info("OrderDetailsAggregate===========", orderDetail, order)
	return orderDetail
}
