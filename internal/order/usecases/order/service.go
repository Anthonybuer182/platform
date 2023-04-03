package order

import (
	"context"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"golang.org/x/exp/slog"
	"platform/internal/order/domain"
)

var _ UseCase = (*service)(nil)

var UseCaseSet = wire.NewSet(NewService)

type service struct {
	repo             OrdersRepo
	userDomainSvc    domain.UserDomainService
	productDomainSvc domain.ProductDomainService
	UserEventPub     UserEventPublisher
}

func NewService(repo OrdersRepo,
	userDomainSvc domain.UserDomainService,
	productDomainSvc domain.ProductDomainService,
	publisher UserEventPublisher) UseCase {
	return &service{
		repo:             repo,
		userDomainSvc:    userDomainSvc,
		productDomainSvc: productDomainSvc,
		UserEventPub:     publisher,
	}
}

func (s *service) GetListOrdersDeleted(ctx context.Context) ([]*domain.Order, error) {
	results, err := s.repo.UFindListDeleteOrder(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetListOrdersDeleted")
	}

	for _, order := range results {
		//处理order的订单明细
		orderDetails, err := s.repo.UFindListOrderDetails(ctx, &domain.Order{
			OrderId: order.OrderId,
		})
		details := make([]*domain.OrderDetail, 0)
		for _, detail := range orderDetails {
			orderDetail := domain.OrderDetailAggregate(ctx, detail, s.productDomainSvc)
			details = append(details, orderDetail)
		}

		if err != nil {
			return nil, errors.Wrap(err, "service.GetListOrdersDeleted")
		}
		domain.OrderAggregate(ctx, order, s.userDomainSvc, details)
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
