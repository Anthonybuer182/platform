package repo

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"platform/internal/order/domain"
	"platform/internal/order/infras/postgresql"
	"platform/internal/order/usecases/order"
	"platform/pkg/postgres"
	"strconv"
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

func (o *OrderRepo) UFindListDeleteOrder(ctx context.Context) ([]*domain.Order, error) {
	entities := make([]*domain.Order, 0)
	//todo
	querier := postgresql.New(o.pg.GetDB())

	results, err := querier.GetDeleteOrderList(ctx)
	fmt.Printf("查询订单删除列表结果 %v\n", results)
	if err != nil {
		return nil, errors.Wrap(err, "querier.GetAll")
	}

	for _, item := range results {
		amount, _ := strconv.ParseFloat(item.Amount, 64)
		userId, _ := strconv.Atoi(item.UserID)
		order := &domain.Order{
			OrderId:     int(item.OrderID),
			OrderStatus: item.OrderState,
			Amount:      float32(amount),
			Users:       &domain.Users{UserId: int32(userId)},
		}
		entities = append(entities, order)
	}

	return entities, nil
}

func (o *OrderRepo) UFindListOrderDetails(ctx context.Context, model *domain.Order) ([]*domain.OrderDetails, error) {
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

func (o *OrderRepo) UDeleteOrder(ctx context.Context, entity *domain.CmdOrders) error {
	db := o.pg.GetDB()
	querier := postgresql.New(db)

	tx, err := db.Begin()
	if err != nil {
		return errors.Wrap(err, "OrderedEventHandlerImpl.Handle")
	}

	qtx := querier.WithTx(tx)

	orderId, _ := strconv.Atoi(entity.OrderId.String())
	orderids := int32(orderId)
	err = qtx.UpdateOrder(ctx, postgresql.UpdateOrderParams{
		OrderID:    orderids,
		OrderState: entity.OrderStatus,
	})

	return tx.Commit()
}
