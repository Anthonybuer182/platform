package router

import (
	"context"
	"platform/internal/order/usecases/order"
	"platform/proto/gen"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var _ gen.OrderServiceServer = (*OrderGRPCServer)(nil)

var OrderGRPCServerSet = wire.NewSet(NewOrderGRPCServer)

type OrderGRPCServer struct {
	gen.UnimplementedOrderServiceServer
	orepo order.OrdersRepo
}

func NewOrderGRPCServer(
	grpcServer *grpc.Server,
	uc order.OrdersRepo,
) gen.OrderServiceServer {
	svc := OrderGRPCServer{
		orepo: uc,
	}

	gen.RegisterOrderServiceServer(grpcServer, &svc)

	reflection.Register(grpcServer)

	return &svc
}

func (g *OrderGRPCServer) GetListDeleteOrders(
	ctx context.Context,
	request *gen.GetListDeleteOrdersRequest,
) (*gen.GetListOrderDeleteResponse, error) {
	slog.Info("gRPC client", "http_method", "GET", "http_name", "GetItemTypes")

	res := gen.GetListOrderDeleteResponse{}

	results, err := g.orepo.GetListDeleteOrders(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "OrderGRPCServer-GetItemTypes")
	}

	for _, item := range results {
		res.Orders = append(res.Orders, &gen.OrderDto{
			OrderNum:    item.OrderId,
			OrderStatus: item.OrderStatus,
			Products:    nil,
			Users:       nil,
		})
	}

	return &res, nil
}
