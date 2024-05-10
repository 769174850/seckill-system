package dao

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"seckill_system/Init"
	"seckill_system/srv-product/pdc_model"
)

func SearchProductByID(id int64) (p pdc_model.Product, err error) {

	err = Init.DB.Where("id = ?", id).Find(&p).Error
	log.Println("Search test data:", p)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("search record error : ", err)
	}
	return
}
