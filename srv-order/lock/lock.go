package lock

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type Lock struct {
	Client     *redis.Client
	Key        string
	ExpireTime time.Duration
}

func NewLock(rdb *redis.Client, key string, expireTime time.Duration) *Lock {
	return &Lock{
		Client:     rdb,
		Key:        key,
		ExpireTime: expireTime,
	}
}

func (l *Lock) Lock(id int64) (ok bool, err error) {
	ok, err = l.Client.SetNX(context.Background(), l.Key, id, l.ExpireTime).Result()
	if err != nil {
		log.Println("lock err:", err)
		return
	}

	return
}

func (l *Lock) UnLock() (err error) {
	_, err = l.Client.Del(context.Background(), l.Key).Result()
	if err != nil {
		log.Println("unlock err:", err)
		return
	}
	return
}
