package domain

import (
	"context"
	"golang.org/x/exp/slog"
	shared "platform/internal/pkg/shared_kernel"
	"strconv"
	"time"
)

type Order struct {
	shared.AggregateRoot
	OrderId     int
	Users       *Users
	OrderDate   time.Time
	Amount      float32
	OrderStatus string
	OrderDetail []*OrderDetail
}

func OrderAggregate(ctx context.Context, order *Order, userDomainService UserDomainService, details []*OrderDetail) *Order {

	//通过grpc查询用户服务
	user, err := userDomainService.GetUserById(ctx, &OrderModel{
		UserId: strconv.Itoa(int(order.Users.UserId)),
	})
	var orderUser *Users
	if err != nil && len(user) > 0 {
		userModel := user[0]
		orderUser = &Users{UserId: userModel.UserId, UserName: userModel.UserName, CreateOn: userModel.CreateOn}
	}
	order.OrderDetail = details
	order.Users = orderUser
	slog.Info("OrderAggregate===========", user, order)

	return order
}
