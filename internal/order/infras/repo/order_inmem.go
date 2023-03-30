package repo

import (
	"context"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"platform/internal/order/domain"
	"platform/internal/order/domain/entity"
	"platform/internal/order/infras/postgresql"
	"platform/pkg/postgres"
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
