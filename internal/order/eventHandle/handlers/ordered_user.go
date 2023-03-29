package handlers

import (
	"context"
	"encoding/json"
	"platform/internal/order/domain/entity"
	event2 "platform/internal/order/eventHandle"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"golang.org/x/exp/slog"
	"platform/internal/order/infras/postgresql"
	"platform/internal/pkg/event"
	"platform/pkg/postgres"
	pkgPublisher "platform/pkg/rabbitmq/publisher"
)

type OrderedEventHandlerImpl struct {
	pg      postgres.DBEngine
	userPub pkgPublisher.EventPublisher
}

var _ event2.OrderedEventHandler = (*OrderedEventHandlerImpl)(nil)

var OrderedEventHandlerSet = wire.NewSet(NewOrderedEventHandlerImpl)

func NewOrderedEventHandlerImpl(
	pg postgres.DBEngine,
	userPub pkgPublisher.EventPublisher,
) event2.OrderedEventHandler {
	return &OrderedEventHandlerImpl{
		pg:      pg,
		userPub: userPub,
	}
}

func (h *OrderedEventHandlerImpl) Handle(ctx context.Context, e event.Ordered) error {
	slog.Info("OrderedEventHandlerImpl-Handle", "Ordered", e)

	order := entity.NewOrder(e)

	db := h.pg.GetDB()
	querier := postgresql.New(db)

	tx, err := db.Begin()
	if err != nil {
		return errors.Wrap(err, "OrderedEventHandlerImpl.Handle")
	}

	qtx := querier.WithTx(tx)

	err = qtx.UpdateOrder(ctx, postgresql.UpdateOrderParams{
		OrderID:    e.OrderID,
		OrderState: e.OrderStatus,
	})
	if err != nil {
		slog.Info("failed to call to repo", "error", err)

		return errors.Wrap(err, "OrderedEventHandlerImpl-querier.CreateOrder")
	}

	// todo: it might cause dual-write problem, but we accept it temporary
	for _, event := range order.DomainEvents() {
		eventBytes, err := json.Marshal(event)
		if err != nil {
			return errors.Wrap(err, "json.Marshal[event]")
		}

		if err := h.userPub.Publish(ctx, eventBytes, "text/plain"); err != nil {
			return errors.Wrap(err, "counterPub.Publish")
		}
	}

	return tx.Commit()
}
