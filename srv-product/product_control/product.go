package product_control

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"seckill_system/srv-product/pdc_model"
	"seckill_system/srv-product/service"
	"seckill_system/util"
)

func GetProductInfo(c *gin.Context) {
	p := pdc_model.Product{}

	//获取参数
	if err := c.ShouldBindJSON(&p); err != nil {
		util.ParamError(c)
		return
	}
	fmt.Println("param:", p.ID)

	//获取信息
	resp, err := service.GetInfo(p)
	if err != nil {
		util.InternetError(c)
		return
	}

	util.OKWithData(c, resp)

}
