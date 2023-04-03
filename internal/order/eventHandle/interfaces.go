package eventHandle

import (
	"context"

	"platform/internal/pkg/event"
)

type OrderedDeletedEventHandler interface {
	Handle(context.Context, event.UserOrderDelete) error
}
