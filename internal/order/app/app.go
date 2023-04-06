package app

import (
	"context"
	"encoding/json"
	"platform/internal/order/domain"
	"platform/internal/order/eventHandle"
	"platform/internal/order/usecases/order"
	"platform/proto/gen"

	"platform/cmd/order/config"
	"platform/internal/pkg/event"
	"platform/pkg/postgres"
	pkgConsumer "platform/pkg/rabbitmq/consumer"

	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/exp/slog"
)

type App struct {
	Cfg                 *config.Config
	PG                  postgres.DBEngine
	AMQPConn            *amqp.Connection
	Consumer            pkgConsumer.EventConsumer
	OrderPub            order.UserEventPublisher
	Repo                order.OrdersRepo
	DomainRepo          domain.OrderRepo
	UserDomainServer    domain.UserDomainService
	ProductDomainServer domain.ProductDomainService
	AggregateService    domain.AggregateService
	UC                  order.UseCase
	OrderGRPCServer     gen.OrderServiceServer
	Handler             eventHandle.OrderedDeletedEventHandler
}

func New(
	cfg *config.Config,
	pg postgres.DBEngine,
	amqpConn *amqp.Connection,
	consumer pkgConsumer.EventConsumer,
	orderPub order.UserEventPublisher,
	repo order.OrdersRepo,
	domainRepo domain.OrderRepo,
	userDomainSVC domain.UserDomainService,
	productDomainSvc domain.ProductDomainService,
	aggregateService domain.AggregateService,
	uc order.UseCase,
	orderGRPCServer gen.OrderServiceServer,
	handler eventHandle.OrderedDeletedEventHandler,

) *App {
	return &App{
		Cfg:                 cfg,
		PG:                  pg,
		AMQPConn:            amqpConn,
		Consumer:            consumer,
		OrderPub:            orderPub,
		Repo:                repo,
		DomainRepo:          domainRepo,
		UserDomainServer:    userDomainSVC,
		ProductDomainServer: productDomainSvc,
		AggregateService:    aggregateService,
		UC:                  uc,
		OrderGRPCServer:     orderGRPCServer,
		Handler:             handler,
	}
}

func (c *App) Worker(ctx context.Context, messages <-chan amqp.Delivery) {
	for delivery := range messages {
		slog.Info("processDeliveries", "delivery_tag", delivery.DeliveryTag)
		slog.Info("received", "delivery_type", delivery.Type)

		switch delivery.Type {
		case "kitchen-order-created":
			var payload event.UserOrderDelete
			err := json.Unmarshal(delivery.Body, &payload)

			if err != nil {
				slog.Error("failed to Unmarshal message", err)
			}

			err = c.Handler.Handle(ctx, payload)

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
