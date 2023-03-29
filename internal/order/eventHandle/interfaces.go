package eventHandle

import (
	"context"

	"platform/internal/pkg/event"
)

type OrderedEventHandler interface {
	Handle(context.Context, event.Ordered) error
}
