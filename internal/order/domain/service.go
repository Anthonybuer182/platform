package domain

import (
	"context"
	"strconv"

	"github.com/google/wire"
	"golang.org/x/exp/slog"
)

var _ AggregateService = (*service)(nil)

var AggregateServiceSet = wire.NewSet(NewService)

type service struct {
	repo             OrderRepo
	userDomainSvc    UserDomainService
	productDomainSvc ProductDomainService
}

func NewService(repo OrderRepo, domainService UserDomainService, productDomainService ProductDomainService) AggregateService {
	return &service{
		repo:             repo,
		userDomainSvc:    domainService,
		productDomainSvc: productDomainService,
	}
}

func (s *service) OrderAggregate(ctx context.Context, order *Order) *Order {

	//通过grpc查询用户服务
	user, err := s.userDomainSvc.GetUserById(ctx, &OrderModel{
		UserId: strconv.Itoa(int(order.Users.UserId)),
	})
	var orderUser *Users
	if err == nil && len(user) > 0 {
		userModel := user[0]
		orderUser = &Users{UserId: userModel.UserId, UserName: userModel.UserName, CreateOn: userModel.CreateOn}
	}
	//处理order的订单明细
	orderParam := &Order{
		OrderId: order.OrderId,
	}
	orderDetails, err := s.repo.FindListOrderDetails(ctx, orderParam)

	details := make([]*OrderDetail, 0)
	for _, detail := range orderDetails {
		orderDetail := OrderDetailAggregate(ctx, detail, s.productDomainSvc)
		details = append(details, orderDetail)
	}

	order.OrderDetail = details
	order.Users = orderUser
	slog.Info("OrderAggregate===========", user, order)

	return order
}
