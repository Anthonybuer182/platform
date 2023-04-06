package domain

import (
	shared "platform/internal/pkg/shared_kernel"
	"time"
)

type Order struct {
	shared.AggregateRoot
	OrderId     int
	Users       *Users
	OrderDate   time.Time
	Amount      float32
	OrderStatus string
	OrderDetail []*OrderDetail
}

func (o Order) Deadline() (deadline time.Time, ok bool) {
	panic("implement me")
}

func (o Order) Done() <-chan struct{} {
	panic("implement me")
}

func (o Order) Err() error {
	panic("implement me")
}

func (o Order) Value(key any) any {
	panic("implement me")
}
