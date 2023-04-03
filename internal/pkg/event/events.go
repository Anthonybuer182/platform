package event

import (
	"github.com/google/uuid"

	shared "platform/internal/pkg/shared_kernel"
)

type UserOrderDelete struct {
	shared.DomainEvent
	OrderID     string    `json:"orderId"`
	ItemLineID  uuid.UUID `json:"itemLineId"`
	OrderStatus string    `json:"orderStatus"`
}

func (e UserOrderDelete) Identity() string {
	return "UserOrderDelete"
}

type UserOrderDeleted struct {
	shared.DomainEvent
	OrderId     string `json:"orderId"`
	OrderStatus string `json:"orderStatus"`
}

func (e UserOrderDeleted) Identity() string {
	return "UserOrderDeleted"
}
