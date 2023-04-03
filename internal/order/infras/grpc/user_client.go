package infrasgrpc

import (
	"context"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"platform/cmd/order/config"
	"platform/internal/order/domain"
	"platform/internal/order/domain/entity"
)

type usersGRPCClient struct {
	conn *grpc.ClientConn
}

var _ domain.UserDomainService = (*usersGRPCClient)(nil)

var UsersGRPCClientSet = wire.NewSet(NewGRPCUserClient)

func NewGRPCUserClient(cfg *config.Config) (domain.UserDomainService, error) {
	conn, err := grpc.Dial(cfg.UsersClient.URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &usersGRPCClient{
		conn: conn,
	}, nil
}

func (p *usersGRPCClient) GetUserById(
	ctx context.Context,
	model *domain.OrderModel,
) ([]*entity.Users, error) {
	//c := gen.NewUserServiceClient(p.conn)

	//res, err := c.GetUserInfo(ctx, &gen.GetUserInfoRequest{})
	//if err != nil {
	//	return nil, errors.Wrap(err, "usersGRPCClient-c.GetItemsByType")
	//}

	results := make([]*entity.Users, 0)
	//for _, item := range res.ItemTypes {
	//	results = append(results, &domain.UserModel{
	//		Name: item.Name,
	//		Passwd: item.Passwd,
	//	})
	//}

	return results, nil
}
