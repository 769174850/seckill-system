package order_control

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"seckill_system/prepare"
	"seckill_system/srv-order/order_model"
	"seckill_system/srv-order/service"
	"seckill_system/util"
	"strconv"
	"time"
)

func CreateOrderEngine(c *gin.Context) {
	o := order_model.Order{}

	//获取现在时间
	nowTime := time.Now().String()
	log.Println("nowTime is :", nowTime)
	//判断是否在活动时间内
	if nowTime < prepare.StartAt {
		util.TimeNotStart(c)
		return
	}

	if nowTime > prepare.EndAt {
		util.TimeHasOUT(c)
		return
	}

	//获取参数
	if err := c.ShouldBindJSON(&o); err != nil {
		log.Println("order parma error:", err)
		util.ParamError(c)
		return
	}
	fmt.Println("order get parma is :", o)

	id, err := prepare.GenerateOrderID()
	if err != nil {
		log.Println("CreateOrderEngine GenerateOrderID call error:", err)
		return
	}
	o.ID = strconv.FormatInt(id, 10)
	fmt.Println("CreateOrderEngine get orderID is :", o.ID)

	//调用服务端
	if _, err := service.CreateOrderServiceCall(o); err != nil {
		fmt.Println("order service error:", err)
		util.InternetError(c)
		return
	}

	util.OK(c)
}
