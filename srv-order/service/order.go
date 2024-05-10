package service

import (
	"context"
	"fmt"
	"log"
	"seckill_system/srv-order/cache"
	"seckill_system/srv-order/kitex_gen/order"
	"seckill_system/srv-order/order_idl"
	"seckill_system/srv-order/order_model"
)

func CreateOrderServiceCall(o order_model.Order) (resp *order.OrderResponse, err error) {
	cache.Mutex.Lock()
	defer cache.Mutex.Unlock()

	fmt.Println("CreateOrderServiceCall get order info :", o)
	fmt.Println("CreateOrderServiceCall get order id :", o.ID)

	//调用服务
	resp, err = order_idl.OrderClient.CreateOrder(context.Background(), &order.Order{
		Id:        o.ID,
		UserId:    o.UserID,
		ProductId: o.ProductID,
	})
	if err != nil {
		log.Println("CreateOrderServiceCall ID :", o.ID, "userID:", o.UserID, "productID:", o.ProductID, "status:", o.Status)
		log.Println("CreateOrderServiceCall call error :", err)
		return
	}
	return
}

func CreateOrderService(o *order.Order) (info order.Order, err error) {
	cache.Mutex.Lock()
	defer cache.Mutex.Unlock()

	fmt.Println("CreateOrderService get order info :", o)
	_, err = cache.CreateOrder(o.UserId, o.ProductId)
	if err != nil {
		log.Println("CreateOrderService call error :", err)
		return
	}
	return
}
