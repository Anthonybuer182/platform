//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"google.golang.org/grpc"
	"platform/cmd/product/config"
	"platform/internal/product/app/router"
	"platform/internal/product/infras/repo"
	productsUC "platform/internal/product/usecases/products"
)

func InitApp(
	cfg *config.Config,
	grpcServer *grpc.Server,
) (*App, error) {
	panic(wire.Build(
		New,
		router.ProductGRPCServerSet,
		repo.RepositorySet,
		productsUC.UseCaseSet,
	))
}
