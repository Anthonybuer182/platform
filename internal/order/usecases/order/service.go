package order

import (
	"context"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"platform/internal/order/domain/entity"
)

var _ OrdersRepo = (*service)(nil)

var OrdersSet = wire.NewSet(NewService)

type service struct {
	repo OrdersRepo
}

func NewService(repo OrdersRepo) OrdersRepo {
	return &service{
		repo: repo,
	}
}

func (s *service) GetListDeleteOrders(ctx context.Context) ([]*entity.Order, error) {
	results, err := s.repo.GetListDeleteOrders(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetItemTypes")
	}

	return results, nil
}
