package events

import (
	"context"

	"platform/internal/pkg/event"
)

type (
	BaristaOrderUpdatedEventHandler interface {
		Handle(context.Context, *event.BaristaOrderUpdated) error
	}

	KitchenOrderUpdatedEventHandler interface {
		Handle(context.Context, *event.OrderUpdated) error
	}
)
