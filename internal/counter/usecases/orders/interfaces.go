package orders

import (
	"context"

	"platform/internal/counter/domain"
	"platform/pkg/rabbitmq/publisher"

	"github.com/google/uuid"
)

type (
	OrderRepo interface {
		GetAll(context.Context) ([]*domain.Order, error)
		GetByID(context.Context, uuid.UUID) (*domain.Order, error)
		Create(context.Context, *domain.Order) error
		Update(context.Context, *domain.Order) (*domain.Order, error)
	}

	BaristaEventPublisher interface {
		Configure(...publisher.Option)
		Publish(context.Context, []byte, string) error
	}

	KitchenEventPublisher interface {
		Configure(...publisher.Option)
		Publish(context.Context, []byte, string) error
	}

	UseCase interface {
		GetListOrderFulfillment(context.Context) ([]*domain.Order, error)
		PlaceOrder(context.Context, *domain.PlaceOrderModel) error
	}
)
