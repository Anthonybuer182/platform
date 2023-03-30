package domain

import (
	"context"
)

type (
	ProductRepo interface {
		GetAll(context.Context) ([]*ItemTypeDto, error)
		GetByTypes(context.Context, []string) ([]*ItemDto, error)
	}
)
type (
	OrderDomainService interface {
		GetDeletedOreders(context.Context, *PlaceOrderModel, bool) ([]*ItemModel, error)
	}
)
