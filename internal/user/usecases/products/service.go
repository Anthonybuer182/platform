package products

import (
	"context"
	"strings"

	"platform/internal/user/domain"

	"github.com/google/wire"
	"github.com/pkg/errors"
)

var _ UseCase = (*service)(nil)

var UseCaseSet = wire.NewSet(NewService)

type service struct {
	repo domain.ProductRepo
}

func NewService(repo domain.ProductRepo) UseCase {
	return &service{
		repo: repo,
	}
}

func (s *service) GetDeletedOrder(ctx context.Context) ([]*domain.OrderDto, error) {
	// 基于rpc调用order服务获取已删除订单列表

	order1 := domain.OrderDto{
		Id:          121,
		ProductId:   212,
		PruductName: "iphone",
		Type:        1,
		Price:       6000,
	}
	orders := []*domain.OrderDto{&order1}

	// results, err := s.repo.GetAll(ctx)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "service.GetItemTypes")
	// }

	return orders, nil
}
func (s *service) DeleteOrder(ctx context.Context) ([]*bool, error) {
	// 基于mq 发布订阅删除订单
	// results, err := s.repo.GetAll(ctx)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "service.GetItemTypes")
	// }

	return nil, nil
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
