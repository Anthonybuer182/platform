package repo

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"platform/internal/order/domain"
	"platform/internal/order/domain/entity"
	"platform/internal/order/infras/postgresql"
	"platform/pkg/postgres"
	"strconv"
)

type DomainOrderRepo struct {
	pg postgres.DBEngine
}

var _ domain.DomainOrderRepo = (*DomainOrderRepo)(nil)

var RepositorySet = wire.NewSet(NewOrderRepo)

func NewOrderRepo(pg postgres.DBEngine) domain.DomainOrderRepo {
	return &DomainOrderRepo{
		pg: pg,
	}
}

func (o *DomainOrderRepo) GetListDeleteOrders(ctx context.Context) ([]*entity.Order, error) {
	entities := make([]*entity.Order, 0)
	//todo
	querier := postgresql.New(o.pg.GetDB())

	results, err := querier.GetDeleteOrderList(ctx)
	fmt.Printf("查询订单删除列表结果 %v\n", results)
	if err != nil {
		return nil, errors.Wrap(err, "querier.GetAll")
	}

	for _, item := range results {
		amount, _ := strconv.ParseFloat(item.Amount, 64)
		orderId := strconv.Itoa(int(item.OrderID))
		order := &entity.Order{
			OrderId:     orderId,
			OrderStatus: item.OrderState,
			Amount:      float32(amount),
		}
		entities = append(entities, order)
	}

	return entities, nil
}
