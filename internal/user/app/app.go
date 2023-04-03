package app

import (
	"context"
	"encoding/json"
	"platform/cmd/user/config"
	shared "platform/internal/pkg/event"
	ordersUC "platform/internal/user/usecases/users"
	pkgConsumer "platform/pkg/rabbitmq/consumer"
	pkgPublisher "platform/pkg/rabbitmq/publisher"
	"platform/proto/gen"

	"platform/internal/user/domain/events/subscribe"

	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/exp/slog"
)

type App struct {
	Cfg            *config.Config
	UC             ordersUC.UseCase
	UserGRPCServer gen.UserServiceServer
	AMQPConn       *amqp.Connection
	Publisher      pkgPublisher.EventPublisher
	Consumer       pkgConsumer.EventConsumer
	OrderPub       ordersUC.OrderEventPublisher
	orderHandler   subscribe.OrderDeletedEventHandler
}

func New(
	cfg *config.Config,
	uc ordersUC.UseCase,
	userGRPCServer gen.UserServiceServer,
	publisher pkgPublisher.EventPublisher,
	consumer pkgConsumer.EventConsumer,
	orderPub ordersUC.OrderEventPublisher,
	amqpConn *amqp.Connection,
	orderHandler subscribe.OrderDeletedEventHandler,
) *App {
	return &App{
		Cfg:            cfg,
		UC:             uc,
		UserGRPCServer: userGRPCServer,
		AMQPConn:       amqpConn,
		Publisher:      publisher,
		Consumer:       consumer,
		orderHandler:   orderHandler,
		OrderPub:       orderPub,
	}
}
func (a *App) Worker(ctx context.Context, messages <-chan amqp.Delivery) {
	for delivery := range messages {
		slog.Info("processDeliveries", "delivery_tag", delivery.DeliveryTag)
		slog.Info("received", "delivery_type", delivery.Type)

		switch delivery.Type {
		case "user-order-deleted":
			var payload shared.UserOrderDeleted

			err := json.Unmarshal(delivery.Body, &payload)
			if err != nil {
				slog.Error("failed to Unmarshal message", err)
			}

			err = a.orderHandler.Handle(ctx, &payload)

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

	slog.Info("Deliveries channel closed")
}
