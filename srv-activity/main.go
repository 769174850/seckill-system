package main

import (
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
	activity "seckill_system/srv-activity/kitex_gen/activity/activityservice"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8889")
	svr := activity.NewServer(new(ActivityServiceImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
