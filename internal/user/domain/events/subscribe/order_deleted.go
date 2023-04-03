package subscribe

import (
	"context"

	"platform/internal/pkg/event"
	"platform/internal/user/usecases/users"

	"github.com/google/wire"
)

type orderDeletedEventHandler struct {
	uc users.UseCase
}

var _ OrderDeletedEventHandler = (*orderDeletedEventHandler)(nil)

var OrderDeletedEventHandlerSet = wire.NewSet(NewOrderDeletedEventHandler)

func NewOrderDeletedEventHandler(uc users.UseCase) OrderDeletedEventHandler {
	return &orderDeletedEventHandler{
		uc: uc,
	}
}

func (h *orderDeletedEventHandler) Handle(ctx context.Context, e *event.OrderDeleted) error {
	// 收到order服务删除订单的消息

	// order, err := h.orderRepo.GetByID(ctx, e.OrderID)
	// if err != nil {
	// 	return errors.Wrap(err, "orderRepo.GetOrderByID")
	// }

	// orderUp := event.OrderUp{
	// 	OrderID:    e.OrderID,
	// 	ItemLineID: e.ItemLineID,
	// 	Name:       e.Name,
	// 	ItemType:   e.ItemType,
	// 	TimeUp:     e.TimeUp,
	// 	MadeBy:     e.MadeBy,
	// }

	// if err = order.Apply(&orderUp); err != nil {
	// 	return errors.Wrap(err, "order.Apply")
	// }

	// _, err = h.orderRepo.Update(ctx, order)
	// if err != nil {
	// 	return errors.Wrap(err, "orderRepo.Update")
	// }

	return nil
}
