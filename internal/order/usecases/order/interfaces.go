package order

import (
	"context"
	"platform/internal/order/domain/entity"
)

type (
	OrdersRepo interface {
		GetListDeleteOrders(context.Context) ([]*entity.Order, error)
	}

	UseCase interface {
		DeleteOrder(context.Context, string) error
	}
)
