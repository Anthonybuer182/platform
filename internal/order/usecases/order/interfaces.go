package order

import (
	"context"
	"platform/internal/order/domain"
	"platform/pkg/rabbitmq/publisher"
)

type (
	//基础设施的数据库操作
	OrdersRepo interface {
		UFindListDeleteOrder(context.Context) ([]*domain.Order, error)
		UDeleteOrder(context.Context, *domain.CmdOrders) error
	}

	//基础设施的mq消息操作
	UserEventPublisher interface {
		Configure(...publisher.Option)
		Publish(context.Context, []byte, string) error
	}

	//用例
	UseCase interface {
		GetListOrdersDeleted(ctx context.Context) ([]*domain.Order, error)
		DeleteOrder(context.Context, *domain.CmdOrders) error
	}
)
