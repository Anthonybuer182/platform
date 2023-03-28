package eventhandlers

import (
	"context"

	"platform/internal/pkg/event"
)

type KitchenOrderedEventHandler interface {
	Handle(context.Context, event.KitchenOrdered) error
}
