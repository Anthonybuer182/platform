package products

import (
	"context"

	"platform/internal/user/domain"
)

type UseCase interface {
	GetDeletedOrder(context.Context) ([]*domain.OrderDto, error)
	DeleteOrder(context.Context) ([]*bool, error)
	GetItemTypes(context.Context) ([]*domain.ItemTypeDto, error)
	GetItemsByType(context.Context, string) ([]*domain.ItemDto, error)
}
