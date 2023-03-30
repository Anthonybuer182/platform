package order

import (
	"context"
	"platform/internal/order/domain/entity"
)

type (
	GrpcOrdersRepo interface {
		GetListOrdersDeleted(context.Context) ([]*entity.Order, error)
	}

	UseCase interface {
		DeleteOrder(context.Context, string) error
	}
)
