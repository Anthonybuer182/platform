package router

import (
	"context"

	"platform/internal/user/usecases/users"
	"platform/proto/gen"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var _ gen.UserServiceServer = (*userGRPCServer)(nil)

var UserGRPCServerSet = wire.NewSet(NewUserGRPCServer)

type userGRPCServer struct {
	gen.UnimplementedUserServiceServer
	uc users.UseCase
}

func NewUserGRPCServer(
	grpcServer *grpc.Server,
	uc users.UseCase,
) gen.UserServiceServer {
	svc := userGRPCServer{
		uc: uc,
	}

	gen.RegisterUserServiceServer(grpcServer, &svc)

	reflection.Register(grpcServer)

	return &svc
}

func (g *userGRPCServer) DeleteOrders(
	ctx context.Context,
	request *gen.DeleteOrdersRequest,
) (*gen.DeleteOrdersResponse, error) {
	slog.Info("gRPC client", "http_method", "POST", "http_name", "DeleteOrders")
	res := gen.DeleteOrdersResponse{}
	// 基于mq的发布订阅删除订单
	result, _ := g.uc.DeleteOrder(ctx)
	res.Sunccess = result
	return &res, nil
}
func (g *userGRPCServer) GetDeletedOrders(
	ctx context.Context,
	request *gen.GetDeletedOrdersRequest,
) (*gen.GetDeletedOrdersResponse, error) {
	slog.Info("gRPC client", "http_method", "GET", "http_name", "GetDeletedOrders")
	res := gen.GetDeletedOrdersResponse{}

	results, err := g.uc.GetDeletedOrder(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "userGRPCServer-GetDeletedOrder")
	}

	for _, item := range results {
		res.Orders = append(res.Orders, &gen.OrdersDto{
			Id:          int32(item.Id),
			PruductId:   int32(item.ProductId),
			PruductName: item.PruductName,
			Type:        int32(item.Type),
			Price:       item.Price,
		})
	}
	// 基于rpc 调用order接口查询已删除的订单
	return &res, nil
}

func (g *userGRPCServer) GetItemTypes(
	ctx context.Context,
	request *gen.GetItemTypesRequest,
) (*gen.GetItemTypesResponse, error) {
	slog.Info("gRPC client", "http_method", "GET", "http_name", "GetItemTypes")

	res := gen.GetItemTypesResponse{}

	results, err := g.uc.GetItemTypes(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "userGRPCServer-GetItemTypes")
	}

	for _, item := range results {
		res.ItemTypes = append(res.ItemTypes, &gen.ItemTypeDto{
			Name:  item.Name,
			Type:  int32(item.Type),
			Price: item.Price,
			Image: item.Image,
		})
	}

	return &res, nil
}

func (g *userGRPCServer) GetItemsByType(
	ctx context.Context,
	request *gen.GetItemsByTypeRequest,
) (*gen.GetItemsByTypeResponse, error) {
	slog.Info("gRPC client", "http_method", "GET", "http_name", "GetItemsByType", "item_types", request.ItemTypes)

	res := gen.GetItemsByTypeResponse{}

	results, err := g.uc.GetItemsByType(ctx, request.ItemTypes)
	if err != nil {
		return nil, errors.Wrap(err, "userGRPCServer-GetItemsByType")
	}

	for _, item := range results {
		res.Items = append(res.Items, &gen.ItemDto{
			Type:  int32(item.Type),
			Price: item.Price,
		})
	}

	return &res, nil
}
