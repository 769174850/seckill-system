package main

import (
	"context"
	order "seckill_system/srv-order/kitex_gen/order"
	"seckill_system/srv-order/service"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// CreateOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) CreateOrder(ctx context.Context, o *order.Order) (resp *order.OrderResponse, err error) {
	// TODO: Your code here...
	_, err = service.CreateOrderService(o)
	return
}

// DeleteOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) DeleteOrder(ctx context.Context, o *order.Order) (resp *order.OrderResponse, err error) {
	// TODO: Your code here...
	return
}

// GetOrderStatus implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) GetOrderStatus(ctx context.Context, o *order.Order) (resp int64, err error) {
	// TODO: Your code here...
	return
}
