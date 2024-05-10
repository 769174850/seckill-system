package dao

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"seckill_system/Init"
	"seckill_system/srv-activity/activity_model"
)

func GetProducts(id int64) (products []activity_model.Product, err error) {
	log.Println("GetProducts:", id)
	err = Init.DB.Where("activity_id = ?", id).Find(&products).Error
	log.Println("GetProducts:", products, len(products))
	if err != nil {
		log.Println("GetAllProduct err:", err)
		return
	}
	return
}

func SearchActivityByID(id int64) (a activity_model.Activity, err error) {
	fmt.Println("SearchActivityByID:", id)
	err = Init.DB.Where("id = ?", id).Find(&a).Error
	fmt.Println("info:", a)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("search record error : ", err)
	}
	return
}
