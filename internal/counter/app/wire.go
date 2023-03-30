//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"platform/cmd/counter/config"
	"platform/internal/counter/app/router"
	"platform/internal/counter/events/handlers"
	"platform/internal/counter/infras"
	infrasGRPC "platform/internal/counter/infras/grpc"
	"platform/internal/counter/infras/repo"
	ordersUC "platform/internal/counter/usecases/orders"
	"platform/pkg/postgres"
	"platform/pkg/rabbitmq"
	pkgConsumer "platform/pkg/rabbitmq/consumer"
	pkgPublisher "platform/pkg/rabbitmq/publisher"
)

func InitApp(
	cfg *config.Config,
	dbConnStr postgres.DBConnString,
	rabbitMQConnStr rabbitmq.RabbitMQConnStr,
	grpcServer *grpc.Server,
) (*App, func(), error) {
	panic(wire.Build(
		New,
		dbEngineFunc,
		rabbitMQFunc,
		pkgPublisher.EventPublisherSet,
		pkgConsumer.EventConsumerSet,

		infras.BaristaEventPublisherSet,
		infras.KitchenEventPublisherSet,
		infrasGRPC.ProductGRPCClientSet,
		router.CounterGRPCServerSet,
		repo.RepositorySet,
		ordersUC.UseCaseSet,
		handlers.BaristaOrderUpdatedEventHandlerSet,
		handlers.KitchenOrderUpdatedEventHandlerSet,
	))
}

func dbEngineFunc(url postgres.DBConnString) (postgres.DBEngine, func(), error) {
	db, err := postgres.NewPostgresDB(url)
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
