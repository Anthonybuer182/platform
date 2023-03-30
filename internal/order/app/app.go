package app

import (
	"context"
	"encoding/json"
	"platform/internal/order/eventHandle"
	"platform/proto/gen"

	"platform/cmd/order/config"
	"platform/internal/pkg/event"
	"platform/pkg/postgres"
	pkgConsumer "platform/pkg/rabbitmq/consumer"
	pkgPublisher "platform/pkg/rabbitmq/publisher"

	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/exp/slog"
)

type App struct {
	Cfg *config.Config

	PG       postgres.DBEngine
	AMQPConn *amqp.Connection

	OrderPub        pkgPublisher.EventPublisher
	Consumer        pkgConsumer.EventConsumer
	OrderGRPCServer gen.OrderServiceServer
	handler         eventHandle.OrderedEventHandler
}

func New(
	cfg *config.Config,
	pg postgres.DBEngine,
	amqpConn *amqp.Connection,
	counterOrderPub pkgPublisher.EventPublisher,
	consumer pkgConsumer.EventConsumer,
	orderGRPCServer gen.OrderServiceServer,
	handler eventHandle.OrderedEventHandler,
) *App {
	return &App{
		Cfg:             cfg,
		PG:              pg,
		AMQPConn:        amqpConn,
		OrderPub:        counterOrderPub,
		Consumer:        consumer,
		OrderGRPCServer: orderGRPCServer,
		handler:         handler,
	}
}

func (c *App) Worker(ctx context.Context, messages <-chan amqp.Delivery) {
	for delivery := range messages {
		slog.Info("processDeliveries", "delivery_tag", delivery.DeliveryTag)
		slog.Info("received", "delivery_type", delivery.Type)

		switch delivery.Type {
		case "kitchen-order-created":
			var payload event.Ordered
			err := json.Unmarshal(delivery.Body, &payload)

			if err != nil {
				slog.Error("failed to Unmarshal message", err)
			}

			err = c.handler.Handle(ctx, payload)

			if err != nil {
				if err = delivery.Reject(false); err != nil {
					slog.Error("failed to delivery.Reject", err)
				}

				slog.Error("failed to process delivery", err)
			} else {
				err = delivery.Ack(false)
				if err != nil {
					slog.Error("failed to acknowledge delivery", err)
				}
			}
		default:
			slog.Info("default")
		}
	}

	slog.Info("deliveries channel closed")
}
