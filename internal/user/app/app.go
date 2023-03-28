package app

import (
	"platform/cmd/product/config"
	productUC "platform/internal/product/usecases/products"
	"platform/proto/gen"
)

type App struct {
	Cfg               *config.Config
	UC                productUC.UseCase
	ProductGRPCServer gen.ProductServiceServer
}

func New(
	cfg *config.Config,
	uc productUC.UseCase,
	productGRPCServer gen.ProductServiceServer,
) *App {
	return &App{
		Cfg:               cfg,
		UC:                uc,
		ProductGRPCServer: productGRPCServer,
	}
}
