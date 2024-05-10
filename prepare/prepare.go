package prepare

import (
	"context"
	"errors"
	"fmt"
	"log"
	"seckill_system/Init"
	"seckill_system/srv-order/lock"
	"time"
)

var StartAt string
var EndAt string

func Prepare() (err error) {
	//缓存商品库存
	err = GetActivityStock()
	if err != nil {
		log.Println("get activity stock error :", err)
		return
	}

	//获取活动时间
	startAt, endAt, err := GetActivityTime("1")
	if err != nil {
		log.Println("get activity time error :", err)
		return err
	}
	StartAt = startAt
	EndAt = endAt

	return
}

type ActivityTime struct {
	StartAt string `json:"start_at"`
	EndAt   string `json:"end_at"`
}

func GetActivityTime(id string) (startAt, endAt string, err error) {
	log.Println("Activity ID is:", id)
	var activityTime ActivityTime

	// 查询活动时间
	err = Init.DB.Table("activities").Where("id = ?", id).Select("start_at, end_at").Scan(&activityTime).Error

	// 获取活动时间
	startAt = activityTime.StartAt
	endAt = activityTime.EndAt
	fmt.Println("time is :", startAt, endAt)

	if err != nil {
		log.Println("Get activity time call error")
		return
	}
	return
}

func GetActivityStock() (err error) {
	type ProductStock struct {
		ID    uint `json:"id"`
		Stock uint `json:"stock"`
	}

	var products []ProductStock

	productStock := "product_stock"
	// 查询活动库存
	err = Init.DB.Table("products").Select("id, stock").Find(&products).Error
	if err != nil {
		log.Println("Get activity stock call error")
		return
	}
	log.Println("Get activity stock :", products)

	for _, product := range products {
		err := Init.Rdb.HSet(context.Background(), productStock, fmt.Sprintf("%d", product.ID), product.Stock).Err()
		if err != nil {
			log.Println("Redis HSet 错误:", err)
			return err
		}
	}

	return
}
func GenerateOrderID() (int64, error) {
	// 获取分布式锁
	lockKey := "order_id_lock"
	l := lock.NewLock(Init.Rdb, lockKey, 10*time.Second)

	lockValue := int64(123)
	ok, err := l.Lock(lockValue)
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, errors.New("获取锁失败")
	}
	defer func(l *lock.Lock) {
		err := l.UnLock()
		if err != nil {
			log.Println("prepare unlock call error:", err)
		}
	}(l)

	// 生成订单 ID
	orderID, err := Init.Rdb.Incr(context.Background(), "order_id").Result()
	if err != nil {
		return 0, err
	}
	return orderID, nil
}
