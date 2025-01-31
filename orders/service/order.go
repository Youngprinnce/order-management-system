package service

import (
	"context"

	"github.com/youngprinnce/order-management-system/common/genproto/orders"
)

var ordersDb = make([]*orders.Order, 0)

type OrderService struct {
	// store
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *orders.Order) error {
	ordersDb = append(ordersDb, order)
	return nil
}

func (s *OrderService) GetOrders(ctx context.Context) []*orders.Order {
	return ordersDb
}

func (s *OrderService) GetOrder(ctx context.Context, orderID int32) (*orders.Order, error) {
	for _, o := range ordersDb {
		if o.OrderID == orderID {
			return o, nil
		}
	}
	return nil, nil
}
