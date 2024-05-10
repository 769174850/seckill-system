package Init

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"seckill_system/srv-activity/activity_model"
	"seckill_system/srv-order/order_model"
	"seckill_system/srv-product/pdc_model"
	"seckill_system/user/user_model"
)

var DB *gorm.DB

func init() {
	err := Mysql()
	if err != nil {
		return
	}
	err = Redis()
	if err != nil {
		return
	}
}

func Mysql() error {
	var dns = "root:123456@tcp(127.0.0.1:3306)/seckill?charset=utf8mb4&parseTime=True&loc=Local" //链接数据库
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Println("mysql call error", err)
	}
	DB = db

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.SetConnMaxIdleTime(100000)
	sqlDB.SetMaxOpenConns(100000)

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("database connect error:", err)
	}

	err = db.AutoMigrate(&user_model.User{})
	if err != nil {
		log.Println("create table user error:", err)
		return err
	}

	err = db.AutoMigrate(&pdc_model.Product{})
	if err != nil {
		log.Println("create table product error:", err)
		return err
	}

	err = db.AutoMigrate(&activity_model.Activity{})
	if err != nil {
		log.Println("create table activity error:", err)
		return err
	}

	err = db.AutoMigrate(&order_model.Order{})
	if err != nil {
		log.Println("create table order error:", err)
		return err
	}

	return nil
}
