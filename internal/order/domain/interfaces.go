package domain

import (
	"context"
	"platform/internal/order/domain/entity"
)

type (
	//获取删除订单的信息
	DomainOrder interface {
		GetListDeleteOrders(context.Context) ([]*entity.Order, error)
	}
)
