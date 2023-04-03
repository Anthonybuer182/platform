package mq

import (
	"context"

	"github.com/google/wire"
	"platform/internal/order/usecases/order"
	"platform/pkg/rabbitmq/publisher"
)

var (
	UserEventPublisherSet = wire.NewSet(NewUserEventPublisher)
)

type (
	UserEventPublisher struct {
		pub publisher.EventPublisher
	}
)

func NewUserEventPublisher(pub publisher.EventPublisher) order.UserEventPublisher {
	return &UserEventPublisher{
		pub: pub,
	}
}

func (p *UserEventPublisher) Configure(opts ...publisher.Option) {
	p.pub.Configure(opts...)
}

func (p *UserEventPublisher) Publish(ctx context.Context, body []byte, contentType string) error {
	return p.pub.Publish(ctx, body, contentType)
}
