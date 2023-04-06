package infrasgrpc

import (
	"context"
	"platform/cmd/order/config"
	"platform/internal/order/domain"
	"platform/proto/gen"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
) ([]*domain.Users, error) {
	c := gen.NewUserServiceClient(p.conn)
	ids := make([]string, 0)
	ids = append(ids, model.UserId)
	slog.Info("usergrpc-getuser:", ids, ctx)
	res, err := c.GetUsers(ctx, &gen.GetUsersRequest{Id: ids})
	slog.Info("usergrpc-getuser:", res, err)
	if err != nil {
		return nil, errors.Wrap(err, "usersGRPCClient-c.GetItemsByType")
	}

	results := make([]*domain.Users, 0)
	for _, item := range res.GetUsers() {
		results = append(results, &domain.Users{
			UserName: item.UserName,
			UserId:   item.Id,
		})
	}

	return results, nil
}
