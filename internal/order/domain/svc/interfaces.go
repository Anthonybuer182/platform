package svc

import (
	"context"
	"platform/internal/order/domain"
)

type (
	////获取订单商品的信息
	ProductDomainService interface {
		GetProductLists(context.Context, *domain.OrderModel) ([]*domain.Products, error)
	}

	//获取用户的信息
	UserDomainService interface {
		GetUserById(context.Context, *domain.OrderModel) ([]*domain.Users, error)
	}

	//基础设施的数据库操作
	OrderRepo interface {
		FindListOrderDetails(context.Context, *domain.Order) ([]*domain.OrderDetails, error)
	}

	AggregateService interface {
		OrderAggregate(ctx context.Context, order *domain.Order) *domain.Order
	}
)
