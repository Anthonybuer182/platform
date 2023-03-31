package order

import (
	"context"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"platform/internal/order/domain"
	"platform/internal/order/domain/entity"
)

var _ GrpcOrdersRepo = (*service)(nil)

var OrdersSet = wire.NewSet(NewService)

type service struct {
	repo domain.DomainOrderRepo
}

func NewService(repo domain.DomainOrderRepo) GrpcOrdersRepo {
	return &service{
		repo: repo,
	}
}

func (s *service) GetListOrdersDeleted(ctx context.Context) ([]*entity.Order, error) {
	results, err := s.repo.GetListDeleteOrders(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetListOrdersDeleted")
	}

	return results, nil
}
