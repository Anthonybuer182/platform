package event

import (
	"github.com/google/uuid"

	shared "platform/internal/pkg/shared_kernel"
)

type Ordered struct {
	shared.DomainEvent
	OrderID     int32     `json:"orderId"`
	ItemLineID  uuid.UUID `json:"itemLineId"`
	OrderStatus string    `json:"orderStatus"`
	Amount      float32   `json:"amount"`
}

func (e Ordered) Identity() string {
	return "Ordered"
}

type OrderDeleted struct {
	shared.DomainEvent
	OrderId     int32  `json:"orderId"`
	OrderStatus string `json:"orderStatus"`
}
