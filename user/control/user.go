package control

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"seckill_system/user/dao"
	"seckill_system/user/user_model"
	"seckill_system/util"
)

func Login(c *gin.Context) {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := c.ShouldBindJSON(&loginRequest) //获取用户名和密码
	if err != nil {
		util.ParamError(c) //获取失败返回错误
		return
	}

	// 获取用户信息
	users, err := dao.GetUser()
	if err != nil {
		// 处理错误
		util.InternetError(c)
		return
	}

	// 遍历用户切片，查找匹配的用户名
	var foundUser user_model.User
	for _, user := range users {
		if user.Username == loginRequest.Username {
			foundUser = user
			break
		}
	}

	// 如果没有找到对应的用户名，返回错误
	if foundUser.Username == "" {
		util.UserFoundError(c)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(loginRequest.Password))
	if err != nil {
		// 密码不匹配，返回错误
		util.UnauthorizedError(c)
		return
	}

	util.OK(c)
}

func Register(c *gin.Context) {
	var User user_model.User
	err := c.ShouldBindJSON(&User) //获取注册的用户信息
	if err != nil {
		util.ParamError(c)
		return
	}

	users, err := dao.GetUser() //检查用户是否被注册
	if err != nil {
		util.InternetError(c)
		return
	}

	for _, user := range users {
		if user.Username == User.Username {
			util.UserHasExist(c)
			return
		}
	}

	err = dao.InsertUser(User)
	if err != nil {
		util.InternetError(c)
		return
	}
	util.OK(c)
}
