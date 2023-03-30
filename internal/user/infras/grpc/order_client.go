package grpc

import (
	"context"

	"platform/cmd/user/config"

	"platform/internal/user/domain"
	gen "platform/proto/gen"

	"github.com/google/wire"
	"github.com/pkg/errors"
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
) ([]*domain.ItemModel, error) {
	c := gen.NewOrderServiceClient(p.conn)

	// itemTypes := ""
	// if isBarista {
	// 	itemTypes = lo.Reduce(model.BaristaItems, func(agg string, item *domain.OrderItemModel, _ int) string {
	// 		return fmt.Sprintf("%s,%s", agg, item.ItemType.String())
	// 	}, "")
	// } else {
	// 	itemTypes = lo.Reduce(model.KitchenItems, func(agg string, item *domain.OrderItemModel, _ int) string {
	// 		return fmt.Sprintf("%s,%s", agg, item.ItemType.String())
	// 	}, "")
	// }

	// res, err := c.GetListDeleteOrders(ctx, &gen.GetListDeleteOrdersRequest{ItemTypes: strings.TrimLeft(itemTypes, ",")})
	res, err := c.GetListDeleteOrders(ctx, &gen.GetListDeleteOrdersRequest{})
	if err != nil {
		return nil, errors.Wrap(err, "orderGRPCClient-c.GetListDeleteOrders")
	}

	results := make([]*domain.ItemModel, 0)
	for _, item := range res.Orders {
		results = append(results, &domain.ItemModel{
			Id:       item.Id,
			OrderNum: item.OrderNum,
		})
	}

	return results, nil
}
