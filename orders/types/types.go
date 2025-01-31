package types

import (
	"context"

	"github.com/youngprinnce/order-management-system/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
	GetOrders(context.Context) []*orders.Order
	GetOrder(context.Context, int32) (*orders.Order, error)
}
