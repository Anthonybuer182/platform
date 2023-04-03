package users

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	events "platform/internal/pkg/event"
	shared "platform/internal/pkg/shared_kernel"
	"platform/internal/user/domain"

	"github.com/google/uuid"
	"github.com/google/wire"
	"github.com/pkg/errors"
)

var _ UseCase = (*service)(nil)

var UseCaseSet = wire.NewSet(NewUseCase)

type service struct {
	repo           domain.ProductRepo
	orderDomainSvc domain.OrderDomainService
	orderEventPub  OrderEventPublisher
}

func NewUseCase(repo domain.ProductRepo, orderDomainSvc domain.OrderDomainService, orderEventPub OrderEventPublisher) UseCase {
	return &service{
		repo:           repo,
		orderDomainSvc: orderDomainSvc,
		orderEventPub:  orderEventPub,
	}
}

func (s *service) GetDeletedOrder(ctx context.Context) ([]*domain.OrderDto, error) {
	// 基于rpc调用order服务获取已删除订单列表
	model := domain.PlaceOrderModel{}
	deletedOrders, err := s.orderDomainSvc.GetDeletedOreders(ctx, &model, true)
	fmt.Println(deletedOrders, err)

	//mock数据
	order1 := domain.OrderDto{
		Id:          121,
		ProductId:   212,
		PruductName: "iphone",
		Type:        1,
		Price:       6000,
	}
	orders := []*domain.OrderDto{&order1}
	return orders, nil
}
func (s *service) DeleteOrder(ctx context.Context) (bool, error) {
	// 基于mq 发布订阅删除订单
	event := events.OrderDelete{
		OrderId:     "1", //uuid.New(),
		OrderStatus: "1",
	}
	order := NewOrder()
	order.ApplyDomain(event)
	orderEvent := order.DomainEvents()[0]
	eventBytes, err := json.Marshal(orderEvent)
	if err != nil {
		return false, errors.Wrap(err, "json.Marshal[event]")
	}
	if err := s.orderEventPub.Publish(ctx, eventBytes, "text/plain"); err != nil {
		return false, nil
	}
	return true, nil
}

type Order struct {
	shared.AggregateRoot
	ID              uuid.UUID
	OrderSource     shared.OrderSource
	LoyaltyMemberID uuid.UUID
	OrderStatus     shared.Status
	Location        shared.Location
	// LineItems       []*LineItem
}

func NewOrder() *Order {
	return &Order{
		ID: uuid.New(),
	}
}

func (s *service) GetItemTypes(ctx context.Context) ([]*domain.ItemTypeDto, error) {
	results, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetItemTypes")
	}

	return results, nil
}

func (s *service) GetItemsByType(ctx context.Context, itemTypes string) ([]*domain.ItemDto, error) {
	types := strings.Split(itemTypes, ",")

	results, err := s.repo.GetByTypes(ctx, types)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetItemsByType")
	}

	return results, nil
}
