// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"platform/cmd/user/config"
	"platform/internal/user/app/router"
	"platform/internal/user/infras/repo"
	"platform/internal/user/usecases/products"
	"google.golang.org/grpc"
)

// Injectors from wire.go:

func InitApp(cfg *config.Config, grpcServer *grpc.Server) (*App, error) {
	productRepo := repo.NewOrderRepo()
	useCase := products.NewService(productRepo)
	productServiceServer := router.NewProductGRPCServer(grpcServer, useCase)
	app := New(cfg, useCase, productServiceServer)
	return app, nil
}
