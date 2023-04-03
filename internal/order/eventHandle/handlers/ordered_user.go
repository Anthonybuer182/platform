package handlers

import (
	"context"
	"encoding/json"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"golang.org/x/exp/slog"
	"platform/internal/order/domain"
	event2 "platform/internal/order/eventHandle"
	"platform/internal/order/usecases/order"
	"platform/internal/pkg/event"
)

type OrderedEventHandlerImpl struct {
	uc      order.UseCase
	userPub order.UserEventPublisher
}

var _ event2.OrderedDeletedEventHandler = (*OrderedEventHandlerImpl)(nil)

var OrderedEventHandlerSet = wire.NewSet(NewOrderedEventHandlerImpl)

func NewOrderedEventHandlerImpl(
	uc order.UseCase,
	userPub order.UserEventPublisher,
) event2.OrderedDeletedEventHandler {
	return &OrderedEventHandlerImpl{
		uc:      uc,
		userPub: userPub,
	}
}

func (h *OrderedEventHandlerImpl) Handle(ctx context.Context, e event.UserOrderDelete) error {
	slog.Info("OrderedEventHandlerImpl-Handle", "UserOrderDelete", e)

	order := domain.DeleteOrder(e)
	//删除订单
	err := h.uc.DeleteOrder(ctx, order)

	if err != nil {
		slog.Info("failed to call to repo", "error", err)

		return errors.Wrap(err, "OrderedEventHandlerImpl-querier.deleteOrder")
	}

	//订单删除操作完成后需要触发的事件
	for _, event := range order.DomainEvents() {
		eventBytes, err := json.Marshal(event)
		if err != nil {
			return errors.Wrap(err, "json.Marshal[event]")
		}

		if err := h.userPub.Publish(ctx, eventBytes, "text/plain"); err != nil {
			return errors.Wrap(err, "orderPub.Publish")
		}
	}

	return err
}
