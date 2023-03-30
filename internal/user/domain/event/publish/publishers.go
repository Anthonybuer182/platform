package publish

import (
	"context"

	"platform/internal/user/usecases/users"

	"platform/pkg/rabbitmq/publisher"

	"github.com/google/wire"
)

var (
	OrderEventPublisherSet = wire.NewSet(NewOrderEventPublisher)
)

type (
	baristaEventPublisher struct {
		pub publisher.EventPublisher
	}
	orderEventPublisher struct {
		pub publisher.EventPublisher
	}
)

func NewOrderEventPublisher(pub publisher.EventPublisher) users.OrderEventPublisher {
	return &orderEventPublisher{
		pub: pub,
	}
}

func (p *orderEventPublisher) Configure(opts ...publisher.Option) {
	p.pub.Configure(opts...)
}

func (p *orderEventPublisher) Publish(ctx context.Context, body []byte, contentType string) error {
	return p.pub.Publish(ctx, body, contentType)
}
