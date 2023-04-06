package order

import (
	"context"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"golang.org/x/exp/slog"
	"platform/internal/order/domain"
	"platform/internal/order/domain/svc"
)

var _ UseCase = (*service)(nil)

var UseCaseSet = wire.NewSet(NewService)

type service struct {
	repo             OrdersRepo
	AggregateService svc.AggregateService
	UserEventPub     UserEventPublisher
}

func NewService(repo OrdersRepo,
	aggregateService svc.AggregateService,
	publisher UserEventPublisher) UseCase {
	return &service{
		repo:             repo,
		AggregateService: aggregateService,
		UserEventPub:     publisher,
	}
}

func (s *service) GetListOrdersDeleted(ctx context.Context) ([]*domain.Order, error) {
	results, err := s.repo.UFindListDeleteOrder(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetListOrdersDeleted")
	}

	for _, order := range results {
		s.AggregateService.OrderAggregate(ctx, order)
	}
	slog.Info("GetListOrdersDeleted=======", results)
	return results, nil
}

func (s *service) DeleteOrder(cxt context.Context, en *domain.CmdOrders) error {
	err := s.repo.UDeleteOrder(cxt, en)
	if err != nil {
		return errors.Wrap(err, "service.U_DeleteOrder")
	}
	return err
}
