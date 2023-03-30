package entity

import (
	"platform/internal/pkg/event"
	"time"

	shared "platform/internal/pkg/shared_kernel"

	"github.com/google/uuid"
)

type Order struct {
	shared.AggregateRoot
	ID          uuid.UUID
	OrderId     string
	OrderStatus string
	Amount      float32
	Orderdate   time.Time
}

func NewOrder(e event.Ordered) *Order {
	order := &Order{
		ID:          uuid.New(),
		OrderStatus: e.OrderStatus,
		Amount:      e.Amount,
		Orderdate:   time.Now(),
	}

	orderDeletedEvent := event.OrderDeleted{
		OrderId:     e.OrderID,
		OrderStatus: e.OrderStatus,
	}

	order.ApplyDomain(orderDeletedEvent)

	return order
}
