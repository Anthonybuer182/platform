package domain

import (
	"context"
)

type (
	////获取订单商品的信息
	ProductDomainService interface {
		GetProductLists(context.Context, *OrderModel) ([]*Products, error)
	}

	//获取用户的信息
	UserDomainService interface {
		GetUserById(context.Context, *OrderModel) ([]*Users, error)
	}

	//基础设施的数据库操作
	OrderRepo interface {
		FindListOrderDetails(context.Context, *Order) ([]*OrderDetails, error)
	}

	AggregateService interface {
		OrderAggregate(ctx context.Context, order *Order) *Order
	}
)
