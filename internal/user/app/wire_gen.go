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
	grpc2 "platform/internal/user/infras/grpc"
)

// Injectors from wire.go:

func InitApp(cfg *config.Config, grpcServer *grpc.Server) (*App, error) {
	// rpc 调用注入 start
	productDomainService, err := grpc2.NewGRPCOrderClient(cfg)
	if err != nil {
		return nil, err
	}


	productRepo := repo.NewOrderRepo()
	// rpc 调用注入到用例中
	useCase := products.NewUseCase(productRepo,productDomainService)
	productServiceServer := router.NewProductGRPCServer(grpcServer, useCase)
	app := New(cfg, useCase, productServiceServer)
	return app, nil
}
