package main

import (
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
	order "seckill_system/srv-order/kitex_gen/order/orderservice"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8890")

	svr := order.NewServer(new(OrderServiceImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
