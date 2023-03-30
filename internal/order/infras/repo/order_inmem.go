package repo

import (
	"context"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"platform/internal/order/domain/entity"
	"platform/internal/order/infras/postgresql"
	"platform/internal/order/usecases/order"
	"platform/pkg/postgres"
)

type OrderRepo struct {
	pg postgres.DBEngine
}

var _ order.OrdersRepo = (*OrderRepo)(nil)

var RepositorySet = wire.NewSet(NewOrderRepo)

func NewOrderRepo(pg postgres.DBEngine) order.OrdersRepo {
	return &OrderRepo{
		pg: pg,
	}
}

func (o *OrderRepo) GetListDeleteOrders(ctx context.Context) ([]*entity.Order, error) {
	entities := make([]*entity.Order, 0)
	//todo
	querier := postgresql.New(o.pg.GetDB())

	results, err := querier.GetDeleteOrderList(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "querier.GetAll")
	}

	for _, item := range results {
		order := &entity.Order{
			OrderId:     string(item.OrderID),
			OrderStatus: item.OrderState,
		}
		entities = append(entities, order)
	}

	return entities, nil
}
