package main

import (
	"log"
	"seckill_system/prepare"
	"seckill_system/router"
	"seckill_system/srv-activity/activity_idl"
	"seckill_system/srv-order/order_idl"
	"seckill_system/srv-product/product_idl"
)

func main() {
	//初始化Product服务
	product_idl.InitProductClient()
	//初始化activity服务
	activity_idl.InitActivityClient()
	//初始化order服务
	order_idl.InitOrderClient()

	//进行活动准备
	err := prepare.Prepare()
	if err != nil {
		log.Println("prepare error :", err)
		return
	}

	//创建gin的路由
	r := router.InitRouter()

	if err := r.Run(":8080"); err != nil {
		log.Println("gin call error :", err)
		return
	}

}
