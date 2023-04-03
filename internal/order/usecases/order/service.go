package order

import (
	"context"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"platform/internal/order/domain"
	"platform/internal/order/domain/entity"
)

var _ UseCase = (*service)(nil)

var UseCaseSet = wire.NewSet(NewService)

type service struct {
	repo          OrdersRepo
	userDomainSvc domain.UserDomainService
	UserEventPub  UserEventPublisher
}

func NewService(repo OrdersRepo,
	userDomainSvc domain.UserDomainService,
	publisher UserEventPublisher) UseCase {
	return &service{
		repo:          repo,
		userDomainSvc: userDomainSvc,
		UserEventPub:  publisher,
	}
}

func (s *service) GetListOrdersDeleted(ctx context.Context) ([]*entity.Order, error) {
	results, err := s.repo.UFindListDeleteOrder(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetListOrdersDeleted")
	}

	for _, order := range results {
		//处理order的订单明细
		orderDetails, err := s.repo.UFindListOrderDetails(ctx, &domain.OrderModel{
			OrderId: order.OrderId.String(),
		})
		if err != nil {
			return nil, errors.Wrap(err, "service.GetListOrdersDeleted")
		}
		entity.OrderAggregate(ctx, order, s.userDomainSvc, orderDetails)
	}
	return results, nil
}

func (s *service) DeleteOrder(cxt context.Context, en *entity.CmdOrders) error {
	err := s.repo.UDeleteOrder(cxt, en)
	if err != nil {
		return errors.Wrap(err, "service.U_DeleteOrder")
	}
	return err
}
