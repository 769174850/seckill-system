package order_idl

import (
	"github.com/cloudwego/kitex/client"
	"log"
	order "seckill_system/srv-order/kitex_gen/order/orderservice"
)

var OrderClient order.Client

func InitOrderClient() {
	o, err := order.NewClient("Order", client.WithHostPorts("localhost:8890"))
	if err != nil {
		log.Println("order client init err:", err)
		return
	}
	OrderClient = o
}
