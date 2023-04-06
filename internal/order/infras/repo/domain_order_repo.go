package repo

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"platform/internal/order/domain"
	"platform/internal/order/domain/svc"
	"platform/internal/order/infras/postgresql"
	"platform/pkg/postgres"
	"strconv"
)

type DomainOrderRepo struct {
	pg postgres.DBEngine
}

var _ svc.OrderRepo = (*DomainOrderRepo)(nil)

var DomainRepositorySet = wire.NewSet(NewDomainOrderRepo)

func NewDomainOrderRepo(pg postgres.DBEngine) svc.OrderRepo {
	return &DomainOrderRepo{
		pg: pg,
	}
}

func (o *DomainOrderRepo) FindListOrderDetails(ctx context.Context, model *domain.Order) ([]*domain.OrderDetails, error) {
	entities := make([]*domain.OrderDetails, 0)
	//todo
	querier := postgresql.New(o.pg.GetDB())
	results, err := querier.GetOrderDetails(ctx, int32(model.OrderId))
	fmt.Printf("查询订单详情结果 %v\n", results)
	if err != nil {
		return nil, errors.Wrap(err, "querier.GetAll")
	}

	for _, item := range results {
		price, _ := strconv.ParseFloat(item.Price, 64)

		orderDetails := &domain.OrderDetails{
			OrderID:   int(item.OrderID),
			ProductID: int(item.ProductID),
			Quantity:  int(item.Quantity),
			Price:     float32(price),
		}
		entities = append(entities, orderDetails)
	}

	return entities, nil
}
