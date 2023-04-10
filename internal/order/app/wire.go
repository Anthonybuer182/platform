//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"platform/cmd/order/config"
	router "platform/internal/order/app/route"
	"platform/internal/order/domain"
	"platform/internal/order/eventHandle/handlers"
	infrasgrpc "platform/internal/order/infras/grpc"
	"platform/internal/order/infras/mq"
	"platform/internal/order/infras/repo"
	"platform/internal/order/usecases/order"
	"platform/pkg"
	"platform/pkg/db"
	"platform/pkg/rabbitmq"
	pkgConsumer "platform/pkg/rabbitmq/consumer"
	pkgPublisher "platform/pkg/rabbitmq/publisher"
)

func InitApp(
	cfg *config.Config,
	ds *config.DataSource,
	rabbitMQConnStr rabbitmq.RabbitMQConnStr,
	grpcServer *grpc.Server,
) (*App, func(), error) {
	panic(wire.Build(
		New,
		dbEngineFunc,
		rabbitMQFunc,
		pkgConsumer.EventConsumerSet,
		pkgPublisher.EventPublisherSet,
		mq.UserEventPublisherSet,
		repo.RepositorySet,
		repo.DomainRepositorySet,
		infrasgrpc.UsersGRPCClientSet,
		infrasgrpc.ProductsGRPCClientSet,
		domain.AggregateServiceSet,
		order.UseCaseSet,
		router.OrderGRPCServerSet,
		handlers.OrderedEventHandlerSet,
	))
}

func dbEngineFunc(dts *config.DataSource) (pkg.DB, func(), error) {
	db, err := db.GetDb(*dts)
	if err != nil {
		return nil, nil, err
	}
	return db, func() { db.Close() }, nil
}

func rabbitMQFunc(url rabbitmq.RabbitMQConnStr) (*amqp.Connection, func(), error) {
	conn, err := rabbitmq.NewRabbitMQConn(url)
	if err != nil {
		return nil, nil, err
	}
	return conn, func() { conn.Close() }, nil
}
