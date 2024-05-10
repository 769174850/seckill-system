package router

import (
	"github.com/gin-gonic/gin"
	"seckill_system/srv-activity/activity_control"
	"seckill_system/srv-order/order_control"
	"seckill_system/srv-product/product_control"
	"seckill_system/user/control"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	//用户接口
	u := r.Group("/user")
	{
		u.POST("/login", control.Login)       //用户登录
		u.POST("/register", control.Register) //用户注册
	}

	//产品接口
	r.GET("/product/info", product_control.GetProductInfo) //获取商品信息

	//服务接口
	a := r.Group("/activity")
	{
		a.GET("/product", activity_control.GetAllProducts)  //获取参加了活动的商品
		a.GET("/info", activity_control.GetActivityMessage) //获取活动详细信息
	}

	//订单接口
	o := r.Group("/order")
	{
		o.POST("/create", order_control.CreateOrderEngine)
	}

	return r
}
