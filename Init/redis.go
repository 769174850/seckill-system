package Init

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

var Rdb *redis.Client

func Redis() error {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := Rdb.Ping(context.Background()).Result() // 检验是否链接成功
	if err != nil {
		log.Println(err) // 输出错误日志
		return err
	}

	val, err := Rdb.Exists(context.Background(), "order_id").Result()
	if err != nil {
		log.Println("Redis Exists call error :", err)
		return err
	}
	if val == 0 {
		err = Rdb.Set(context.Background(), "order_id", 0, 0).Err()
		if err != nil {
			// 处理错误
			log.Println("Redis Set call error :", err)
			return err
		}
	}

	return nil
}
