//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"google.golang.org/grpc"
	"platform/cmd/user/config"
	"platform/internal/user/app/router"
	"platform/internal/user/infras/repo"
	productsUC "platform/internal/user/usecases/products"
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
