package router

import (
	"context"

	"platform/internal/user/domain"
	"platform/internal/user/usecases/users"
	"platform/proto/gen"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/samber/lo"
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
func (g *userGRPCServer) GetUsers(
	ctx context.Context,
	request *gen.GetUsersRequest,
) (*gen.GetUsersResponse, error) {
	slog.Info("gRPC client", "http_method", "POST", "http_name", "DeleteOrders")
	mockUsers := []string{"张三", "李四", "王五", "赵六"}
	res := gen.GetUsersResponse{}
	for i, v := range request.Id {
		res.Users = append(res.Users, &gen.UsersDto{UserName: mockUsers[i] + v})
	}
	return &res, nil
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
		res.Orders = append(res.Orders, &gen.OrderDtos{
			Id:          item.Id,
			OrderNum:    item.OrderNum,
			OrderStatus: item.OrderStatus,
			Details: lo.Map(item.DetailsDto, func(item *domain.DetailsDto, _ int) *gen.DetailsDtos {
				return &gen.DetailsDtos{
					Id: item.Id,
					Products: &gen.ProductDtos{
						ProductName: item.ProductDto.ProductName,
						Category:    item.ProductDto.Category,
						Price:       item.ProductDto.Price,
					},
				}
			}),
			Users: &gen.UserDtos{
				Id:   item.UserDto.Id,
				Name: item.UserDto.Name,
			},
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
