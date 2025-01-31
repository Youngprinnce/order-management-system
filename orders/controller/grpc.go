package controller

import (
	"context"
	"log"

	"github.com/youngprinnce/order-management-system/common/genproto/orders"
	"github.com/youngprinnce/order-management-system/orders/types"
	"google.golang.org/grpc"
)

type OrdersGrpcController struct {
	ordersService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewOrdersGrpcController(grpc *grpc.Server, ordersService types.OrderService) {
	gRPCHandler := &OrdersGrpcController{
		ordersService: ordersService,
	}
	orders.RegisterOrderServiceServer(grpc, gRPCHandler)
}

func (h *OrdersGrpcController) GetOrders(ctx context.Context, req *orders.GetOrdersRequest) (*orders.GetOrderResponse, error) {
	o := h.ordersService.GetOrders(ctx)
	log.Println("Orders: ", o)
	res := &orders.GetOrderResponse{
		Orders: o,
	}

	return res, nil
}

func (h *OrdersGrpcController) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID:    int32(len(h.ordersService.GetOrders(ctx)) + 1),
		CustomerID: req.CustomerID,
		ProductID:  req.ProductID,
		Quantity:   req.Quantity,
	}

	err := h.ordersService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}

	return res, nil
}

func (h *OrdersGrpcController) GetOrder(ctx context.Context, req *orders.GetOrderRequest) (*orders.Order, error) {
	o, err := h.ordersService.GetOrder(ctx, req.OrderID)
	if err != nil {
		return nil, err
	}

	return o, nil
}
