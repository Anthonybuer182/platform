package events

import (
	"context"

	"platform/internal/pkg/event"
)

type (
	OrderDeletedEventHandler interface {
		Handle(context.Context, *event.OrderDeleted) error
	}
)
