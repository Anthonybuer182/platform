package domain

import (
	"context"
	"platform/internal/order/domain/entity"
)

type (
	////获取订单商品的信息
	//ProductDomainService interface {
	//	GetProductLists(context.Context, *OrderModel) ([]*entity.Products, error)
	//}

	//获取删除订单的信息
	UserDomainService interface {
		GetUserById(context.Context, *OrderModel) ([]*entity.Users, error)
	}
)
