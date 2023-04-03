package infrasgrpc

import (
	"context"
	"github.com/google/wire"
	"platform/internal/order/domain"
)

type productGRPCClient struct {
	//conn *grpc.ClientConn
}

var _ domain.ProductDomainService = (*productGRPCClient)(nil)

var ProductsGRPCClientSet = wire.NewSet(NewGRPCProductClient)

func NewGRPCProductClient() (domain.ProductDomainService, error) {
	//conn, err := grpc.Dial(cfg.UsersClient.URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	return nil, err
	//}

	return &productGRPCClient{}, nil
}

func (p *productGRPCClient) GetProductLists(
	ctx context.Context,
	model *domain.OrderModel,
) ([]*domain.Products, error) {

	results := make([]*domain.Products, 0)
	results = append(results, &domain.Products{
		ProductName: "大象",
		Price:       2000,
	})

	results = append(results, &domain.Products{
		ProductName: "猴子",
		Price:       4000,
	})
	return results, nil
}
