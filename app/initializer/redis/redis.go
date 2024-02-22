package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var rdb *redis.Client

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "",
		DB:           0,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
	ping, err := rdb.Ping().Result()
	if err == redis.Nil {
		panic("Redis异常")
	} else if err != nil {
		panic("Redis连接失败：" + err.Error())
	} else {
		fmt.Print(ping)
	}
}
