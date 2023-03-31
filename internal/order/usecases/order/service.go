package order

import (
	"context"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"platform/internal/order/domain/entity"
)

var _ UseCase = (*service)(nil)

var UseCaseSet = wire.NewSet(NewService)

type service struct {
	repo         OrdersRepo
	UserEventPub UserEventPublisher
}

func NewService(repo OrdersRepo,
	publisher UserEventPublisher) UseCase {
	return &service{
		repo:         repo,
		UserEventPub: publisher,
	}
}

func (s *service) GetListOrdersDeleted(ctx context.Context) ([]*entity.Order, error) {
	results, err := s.repo.FindListDeleteOrder(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetListOrdersDeleted")
	}

	return results, nil
}

func (s *service) DeleteOrder(cxt context.Context, en entity.Order) error {
	err := s.repo.DeleteOrder(cxt, en)
	if err != nil {
		return errors.Wrap(err, "service.DeleteOrder")
	}
	return nil
}
