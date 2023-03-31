package order

import (
	"context"
	"platform/internal/order/domain/entity"
	"platform/pkg/rabbitmq/publisher"
)

type (
	OrdersRepo interface {
		FindListDeleteOrder(context.Context) ([]*entity.Order, error)
		DeleteOrder(context.Context, entity.Order) error
	}

	UseCase interface {
		GetListOrdersDeleted(ctx context.Context) ([]*entity.Order, error)
		DeleteOrder(context.Context, entity.Order) error
	}

	UserEventPublisher interface {
		Configure(...publisher.Option)
		Publish(context.Context, []byte, string) error
	}
)
