package grpc

import (
	"context"

	"platform/cmd/user/config"

	"platform/internal/user/domain"
	gen "platform/proto/gen"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type orderGRPCClient struct {
	conn *grpc.ClientConn
}

var _ domain.OrderDomainService = (*orderGRPCClient)(nil)

var ProductGRPCClientSet = wire.NewSet(NewGRPCOrderClient)

func NewGRPCOrderClient(cfg *config.Config) (domain.OrderDomainService, error) {
	conn, err := grpc.Dial(cfg.OrderClient.URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &orderGRPCClient{
		conn: conn,
	}, nil
}

func (p *orderGRPCClient) GetDeletedOreders(
	ctx context.Context,
	model *domain.PlaceOrderModel,
	isBarista bool,
) ([]*domain.OrderDto, error) {
	c := gen.NewOrderServiceClient(p.conn)
	// res, err := c.GetListDeleteOrders(ctx, &gen.GetListDeleteOrdersRequest{ItemTypes: strings.TrimLeft(itemTypes, ",")})
	res, err := c.GetListDeleteOrders(ctx, &gen.GetListDeleteOrdersRequest{})
	if err != nil {
		return nil, errors.Wrap(err, "orderGRPCClient-c.GetListDeleteOrders")
	}

	results := make([]*domain.OrderDto, 0)
	for _, item := range res.Orders {
		results = append(results, &domain.OrderDto{
			Id:          item.Id,
			OrderNum:    item.OrderNum,
			OrderStatus: item.OrderStatus,
			DetailsDto: lo.Map(item.Details, func(item *gen.DetailsDto, _ int) *domain.DetailsDto {
				return &domain.DetailsDto{
					Id: item.Id,
					ProductDto: domain.ProductDto{
						ProductName: item.Products.ProductName,
						Category:    item.Products.Category,
						Price:       item.Products.Price,
					},
				}
			}),
			UserDto: &domain.UserDto{
				Id:   item.Users.Id,
				Name: item.Users.Name,
			},
		})
	}

	return results, nil
}
