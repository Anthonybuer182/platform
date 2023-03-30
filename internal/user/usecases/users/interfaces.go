package users

import (
	"context"

	"platform/internal/user/domain"
	"platform/pkg/rabbitmq/publisher"
)

type (
	OrderEventPublisher interface {
		Configure(...publisher.Option)
		Publish(context.Context, []byte, string) error
	}
)

type UseCase interface {
	GetDeletedOrder(context.Context) ([]*domain.OrderDto, error)
	DeleteOrder(context.Context) ([]*bool, error)
	GetItemTypes(context.Context) ([]*domain.ItemTypeDto, error)
	GetItemsByType(context.Context, string) ([]*domain.ItemDto, error)
}
