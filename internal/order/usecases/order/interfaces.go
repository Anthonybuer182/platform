package order

import (
	"context"
	"platform/internal/order/domain"
	"platform/internal/order/domain/entity"
	"platform/pkg/rabbitmq/publisher"
)

type (
	//基础设施的数据库操作
	OrdersRepo interface {
		UFindListDeleteOrder(context.Context) ([]*entity.Order, error)
		UFindListOrderDetails(context.Context, *domain.OrderModel) ([]*entity.OrderDetail, error)
		UDeleteOrder(context.Context, *entity.CmdOrders) error
	}

	//基础设施的mq消息操作
	UserEventPublisher interface {
		Configure(...publisher.Option)
		Publish(context.Context, []byte, string) error
	}

	//用例
	UseCase interface {
		GetListOrdersDeleted(ctx context.Context) ([]*entity.Order, error)
		DeleteOrder(context.Context, *entity.CmdOrders) error
	}
)
