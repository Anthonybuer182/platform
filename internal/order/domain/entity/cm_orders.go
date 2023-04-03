package entity

import (
	"github.com/google/uuid"
	"platform/internal/pkg/event"
	shared "platform/internal/pkg/shared_kernel"
	"time"
)

type CmdOrders struct {
	shared.AggregateRoot

	OrderId     uuid.UUID
	UserId      int
	OrderDate   time.Time
	Amount      float32
	OrderStatus string
}

func NewOrders(userId int, amount float32, orderStatus string) *CmdOrders {
	return &CmdOrders{
		OrderId:     uuid.New(),
		UserId:      userId,
		OrderDate:   time.Now(),
		Amount:      amount,
		OrderStatus: orderStatus,
	}
}

// 领域订单的删除事件
func DeleteOrder(e event.UserOrderDelete) *CmdOrders {
	order := &CmdOrders{
		OrderId:     uuid.MustParse(e.OrderID),
		OrderStatus: e.OrderStatus,
		OrderDate:   time.Now(),
	}

	orderDeletedEvent := event.UserOrderDeleted{
		OrderId:     e.OrderID,
		OrderStatus: e.OrderStatus,
	}

	order.ApplyDomain(orderDeletedEvent)
	return order
}
