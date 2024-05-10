package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Status  int
	Message string
}

var (
	Ok = Response{
		Status:  200,
		Message: "success",
	} //成功的返回值

	BadRequest = Response{
		Status:  400,
		Message: "bad request",
	} //错误的请求的返回值

	internalServerError = Response{
		Status:  500,
		Message: "internal server error",
	} //服务器错误的返回值

	parameterMissing = Response{
		Status:  400,
		Message: "parameter missing",
	} //参数获取失败的返回值

	UserNotFound = Response{
		Status:  40000,
		Message: "user not found",
	} //用户不存在的返回值

	unauthorized = Response{
		Status:  40001,
		Message: "unauthorized",
	} //未授权的返回值

	userExist = Response{
		Status:  30000,
		Message: "user has exist!",
	} //用户已存在的返回值

	TimeHasNotStart = Response{
		Status:  40002,
		Message: "Seckill Activity Not Start!",
	} //秒杀活动未开始的返回值

	TimeOut = Response{
		Status:  40002,
		Message: "Seckill Activity Time Out!",
	} //秒杀活动已结束的返回值
)

func OK(c *gin.Context) {
	c.JSON(http.StatusOK, Ok)
}

func OKWithData(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Success",
		"data":    data,
	})
}

func ParamError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, parameterMissing)
}

func InternetError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, internalServerError)
}

func UserFoundError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, UserNotFound)
}

func UnauthorizedError(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, unauthorized)
}

func UserHasExist(c *gin.Context) {
	c.JSON(http.StatusBadRequest, userExist)
}

func TimeNotStart(c *gin.Context) {
	c.JSON(http.StatusBadRequest, TimeHasNotStart)
}

func TimeHasOUT(c *gin.Context) {
	c.JSON(http.StatusBadRequest, TimeOut)
}
