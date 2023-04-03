package router

import (
	"context"
	"platform/internal/order/usecases/order"
	"platform/proto/gen"
	"strconv"

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
	uc order.UseCase
}

func NewOrderGRPCServer(
	grpcServer *grpc.Server,
	uc order.UseCase,
) gen.OrderServiceServer {
	svc := OrderGRPCServer{
		uc: uc,
	}

	gen.RegisterOrderServiceServer(grpcServer, &svc)

	reflection.Register(grpcServer)

	return &svc
}

func (g *OrderGRPCServer) GetListDeleteOrders(
	ctx context.Context,
	request *gen.GetListDeleteOrdersRequest,
) (*gen.GetListOrderDeleteResponse, error) {
	slog.Info("gRPC client", "http_method", "GET", "http_name", "GetListDeleteOrders")

	res := gen.GetListOrderDeleteResponse{}

	results, err := g.uc.GetListOrdersDeleted(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "OrderGRPCServer-GetListDeleteOrders")
	}

	for _, item := range results {
		users := make([]*gen.UserDto, 1)
		user := &gen.UserDto{
			Name:      item.Users.UserName,
			Telephone: "xxxx",
		}
		users = append(users, user)
		res.Orders = append(res.Orders, &gen.OrderDto{
			OrderNum:    strconv.Itoa(item.OrderId),
			OrderStatus: item.OrderStatus,
			Users:       users,
		})
	}

	return &res, nil
}
