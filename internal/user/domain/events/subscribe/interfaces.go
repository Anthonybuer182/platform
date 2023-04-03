package subscribe

import (
	"context"

	"platform/internal/pkg/event"
)

type (
	OrderDeletedEventHandler interface {
		Handle(context.Context, *event.UserOrderDeleted) error
	}
)
