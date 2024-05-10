package cache

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"seckill_system/Init"
	"seckill_system/prepare"
	"seckill_system/srv-order/lock"
	"seckill_system/srv-order/order_model"
	"strconv"
	"sync"
	"time"
)

var Mutex sync.Mutex

func CreateOrder(userID string, productID string) (message string, err error) {
	Mutex.Lock()
	defer Mutex.Unlock()

	orderID, err := prepare.GenerateOrderID()
	if err != nil {
		log.Println("CreateOrder GenerateOrderID call error :", err)
		return
	}

	fmt.Println("CreateOrder ID:", orderID)
	log.Println("CreateOrder UserID:", userID)
	log.Println("CreateOrder ProductID:", productID)

	orderLockKey := "order"
	productStock := "product:stock"

	// 加锁
	orderLock := lock.NewLock(Init.Rdb, orderLockKey+strconv.FormatInt(orderID, 10), 10*time.Second)
	ok, err := orderLock.Lock(orderID)
	if err != nil {
		log.Println("Add Lock call error :", err)
		return
	}
	if !ok {
		err = errors.New("lock error")
		return
	}
	defer func(orderLock *lock.Lock) {
		err := orderLock.UnLock()
		if err != nil {
			log.Println("UnLock Lock call error :", err)
			return
		}
	}(orderLock)

	//使用redis事务
	_, err = Init.Rdb.TxPipelined(context.Background(), func(pipe redis.Pipeliner) error {
		// 获取库存
		stockCmd := pipe.HGet(context.Background(), productStock, fmt.Sprint(productID))

		// 检验库存是否足够
		if stockCmd.Err() != nil && !errors.Is(stockCmd.Err(), redis.Nil) {
			return stockCmd.Err()
		}
		stock, err := stockCmd.Int64()
		if err != nil {
			return err
		}
		if stock < 1 {
			return errors.New("stock is not enough")
		}

		// 减库存
		pipe.HIncrBy(context.Background(), productStock, fmt.Sprint(productID), -1)

		// 创建订单
		order := order_model.Order{
			ID:        strconv.FormatInt(orderID, 10),
			UserID:    userID,
			ProductID: productID,
			Status:    "1",
		}
		err = Init.DB.Create(&order).Error
		if err != nil {
			log.Println("Create data Order call error :", err)
		}

		_, err = pipe.Exec(context.Background())
		if err != nil && errors.Is(err, redis.TxFailedErr) {
			log.Println("Redis Transaction call error :", err)
			return err
		}
		return nil
	})

	message = "success"

	return
}
