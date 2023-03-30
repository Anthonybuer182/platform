package handlers

import (
	"context"

	"platform/internal/counter/events"
	"platform/internal/counter/usecases/orders"
	"platform/internal/pkg/event"

	"github.com/google/wire"
	"github.com/pkg/errors"
)

type kitchenOrderUpdatedEventHandler struct {
	orderRepo orders.OrderRepo
}

var _ events.KitchenOrderUpdatedEventHandler = (*kitchenOrderUpdatedEventHandler)(nil)

var KitchenOrderUpdatedEventHandlerSet = wire.NewSet(NewKitchenOrderUpdatedEventHandler)

func NewKitchenOrderUpdatedEventHandler(orderRepo orders.OrderRepo) events.KitchenOrderUpdatedEventHandler {
	return &kitchenOrderUpdatedEventHandler{
		orderRepo: orderRepo,
	}
}

func (h *kitchenOrderUpdatedEventHandler) Handle(ctx context.Context, e *event.OrderUpdated) error {
	order, err := h.orderRepo.GetByID(ctx, e.OrderID)
	if err != nil {
		return errors.Wrap(err, "orderRepo.GetOrderByID")
	}

	orderUp := event.OrderUp{
		OrderID:    e.OrderID,
		ItemLineID: e.ItemLineID,
		Name:       e.Name,
		ItemType:   e.ItemType,
		TimeUp:     e.TimeUp,
		MadeBy:     e.MadeBy,
	}

	if err = order.Apply(&orderUp); err != nil {
		return errors.Wrap(err, "order.Apply")
	}

	_, err = h.orderRepo.Update(ctx, order)
	if err != nil {
		return errors.Wrap(err, "orderRepo.Update")
	}

	return nil
}
