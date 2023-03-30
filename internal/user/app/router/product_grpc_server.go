package router

import (
	"context"

	"platform/internal/user/usecases/products"
	"platform/proto/gen"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var _ gen.UserServiceServer = (*productGRPCServer)(nil)

var ProductGRPCServerSet = wire.NewSet(NewProductGRPCServer)

type productGRPCServer struct {
	gen.UnimplementedUserServiceServer
	uc products.UseCase
}

func NewProductGRPCServer(
	grpcServer *grpc.Server,
	uc products.UseCase,
) gen.UserServiceServer {
	svc := productGRPCServer{
		uc: uc,
	}

	gen.RegisterUserServiceServer(grpcServer, &svc)

	reflection.Register(grpcServer)

	return &svc
}

func (g *productGRPCServer) DeleteOrders(
	ctx context.Context,
	request *gen.DeleteOrdersRequest,
) (*gen.DeleteOrdersResponse, error) {
	slog.Info("gRPC client", "http_method", "GET", "http_name", "GetItemTypes")
	res := gen.DeleteOrdersResponse{}
	// 基于mq的发布订阅删除订单

	res.Sunccess = true
	return &res, nil
}
func (g *productGRPCServer) GetDeletedOrders(
	ctx context.Context,
	request *gen.GetDeletedOrdersRequest,
) (*gen.GetDeletedOrdersResponse, error) {
	slog.Info("gRPC client", "http_method", "GET", "http_name", "GetItemTypes")
	res := gen.GetDeletedOrdersResponse{}

	results, err := g.uc.GetDeletedOrder(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "productGRPCServer-GetItemTypes")
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

func (g *productGRPCServer) GetItemTypes(
	ctx context.Context,
	request *gen.GetItemTypesRequest,
) (*gen.GetItemTypesResponse, error) {
	slog.Info("gRPC client", "http_method", "GET", "http_name", "GetItemTypes")

	res := gen.GetItemTypesResponse{}

	results, err := g.uc.GetItemTypes(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "productGRPCServer-GetItemTypes")
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

func (g *productGRPCServer) GetItemsByType(
	ctx context.Context,
	request *gen.GetItemsByTypeRequest,
) (*gen.GetItemsByTypeResponse, error) {
	slog.Info("gRPC client", "http_method", "GET", "http_name", "GetItemsByType", "item_types", request.ItemTypes)

	res := gen.GetItemsByTypeResponse{}

	results, err := g.uc.GetItemsByType(ctx, request.ItemTypes)
	if err != nil {
		return nil, errors.Wrap(err, "productGRPCServer-GetItemsByType")
	}

	for _, item := range results {
		res.Items = append(res.Items, &gen.ItemDto{
			Type:  int32(item.Type),
			Price: item.Price,
		})
	}

	return &res, nil
}
