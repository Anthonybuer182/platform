package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"platform/cmd/user/config"
	"platform/internal/user/app"
	"platform/pkg/logger"
	"platform/pkg/rabbitmq"

	pkgConsumer "platform/pkg/rabbitmq/consumer"

	pkgPublisher "platform/pkg/rabbitmq/publisher"

	"github.com/sirupsen/logrus"
	"go.uber.org/automaxprocs/maxprocs"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

func main() {
	// set GOMAXPROCS
	_, err := maxprocs.Set()
	if err != nil {
		slog.Error("failed set max procs", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("failed get config", err)
	}

	slog.Info("⚡ init app", "name", cfg.Name, "version", cfg.Version)

	// set up logrus
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logger.ConvertLogLevel(cfg.Log.Level))

	// integrate Logrus with the slog logger
	slog.New(logger.NewLogrusHandler(logrus.StandardLogger()))

	server := grpc.NewServer()

	go func() {
		defer server.GracefulStop()
		<-ctx.Done()
	}()

	cleanup := prepareApp(ctx, cancel, cfg, server)

	// gRPC Server.
	address := fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port)
	network := "tcp"

	l, err := net.Listen(network, address)
	if err != nil {
		slog.Error("failed to listen to address", err, "network", network, "address", address)
		cancel()
	}

	slog.Info("🌏 start server...", "address", address)

	defer func() {
		if err1 := l.Close(); err != nil {
			slog.Error("failed to close", err1, "network", network, "address", address)
		}
	}()

	err = server.Serve(l)
	if err != nil {
		slog.Error("failed start gRPC server", err, "network", network, "address", address)
		cancel()
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		cleanup()
		slog.Info("signal.Notify", v)
	case done := <-ctx.Done():
		cleanup()
		slog.Info("ctx.Done", done)
	}
}
func prepareApp(ctx context.Context, cancel context.CancelFunc, cfg *config.Config, server *grpc.Server) func() {
	a, close, err := app.InitApp(cfg, rabbitmq.RabbitMQConnStr(cfg.RabbitMQ.URL), server)
	if err != nil {
		slog.Error("failed init app", err)
		cancel()
		<-ctx.Done()
	}

	a.OrderPub.Configure(
		pkgPublisher.ExchangeName("order-exchange"),
		pkgPublisher.BindingKey("order-routing-key"),
		pkgPublisher.MessageTypeName("order-delete"),
	)

	a.Consumer.Configure(
		pkgConsumer.ExchangeName("user-order-exchange"),
		pkgConsumer.QueueName("user-order-queue"),
		pkgConsumer.BindingKey("user-order-routing-key"),
		pkgConsumer.ConsumerTag("user-order-consumer"),
	)

	go func() {
		err1 := a.Consumer.StartConsumer(a.Worker)
		if err1 != nil {
			slog.Error("failed to start Consumer", err1)
			cancel()
			<-ctx.Done()
		}
	}()

	return close
}
