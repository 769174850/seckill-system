package activity_control

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"seckill_system/srv-activity/activity_model"
	"seckill_system/srv-activity/service"
	"seckill_system/util"
)

func GetAllProducts(c *gin.Context) {
	a := activity_model.Activity{}

	// 绑定参数
	if err := c.ShouldBindJSON(&a); err != nil {
		util.ParamError(c)
		return
	}

	log.Printf("GetAllProducts: a.ID = %d", a.ID) // 添加日志输出 a.ID

	//获取信息
	resp, err := service.GetTheProducts(a)
	fmt.Println("GetAllProducts", resp)
	if err != nil {
		util.InternetError(c)
		return
	}

	util.OKWithData(c, resp)
}

func GetActivityMessage(c *gin.Context) {
	a := activity_model.Activity{}

	//获取参数
	if err := c.ShouldBindJSON(&a); err != nil {
		util.ParamError(c)
		return
	}
	fmt.Println("param:", a.ID)

	//获取信息
	resp, err := service.GetActInfo(a)
	if err != nil {
		util.InternetError(c)
		return
	}

	util.OKWithData(c, resp)

}
