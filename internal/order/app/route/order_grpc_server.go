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
	slog.Info("grpc", res)
	results, err := g.uc.GetListOrdersDeleted(ctx)
	slog.Info("uc=======", results, err)
	if err != nil {
		return nil, errors.Wrap(err, "OrderGRPCServer-GetListDeleteOrders")
	}

	for _, item := range results {
		itemUser := item.Users
		username := itemUser.UserName
		slog.Info("item===========,", username)
		user := &gen.UserDto{
			Name:      username,
			Telephone: "xxxx",
		}

		detailsDtos := make([]*gen.DetailsDto, 0)
		for _, detail := range item.OrderDetail {
			detailDto := &gen.DetailsDto{
				Quantity: int32(detail.Quantity),
				Amount:   float64(detail.Price),
				Products: &gen.ProductDto{
					Price:       detail.Product.Price,
					ProductName: detail.Product.ProductName,
					Category:    detail.Product.Category,
				},
			}
			detailsDtos = append(detailsDtos, detailDto)
		}
		res.Orders = append(res.Orders, &gen.OrderDto{
			OrderNum:    strconv.Itoa(item.OrderId),
			OrderStatus: item.OrderStatus,
			Users:       user,
			Details:     detailsDtos,
		})
	}

	return &res, nil
}
