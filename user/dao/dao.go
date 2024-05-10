package dao

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"seckill_system/Init"
	"seckill_system/user/user_model"
)

func GetUser() ([]user_model.User, error) { //获取用户信息

	var users []user_model.User //定义切片存储用户

	err := Init.DB.Select("ID", "Username", "Password").Find(&users).Error //查询用户
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return users, nil
}

func InsertUser(u user_model.User) (err error) {

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost) //哈希加密
	if err != nil {
		log.Println("HashPassword generate has error :", err)
		return
	}

	if err = Init.DB.Create(&user_model.User{ //存储用户
		Username: u.Username,
		Password: string(hashPassword),
	}).Error; err != nil {
		log.Println("Create user error:", err)
	}

	return
}
