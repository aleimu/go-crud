package cache

import (
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

// RedisClient Redis缓存客户端单例
var RedisClient *redis.Client
var RedisNil = redis.Nil
// Redis 在中间件中初始化redis链接
func Redis() {
	db, _ := strconv.ParseUint(os.Getenv("REDIS_DB"), 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:         os.Getenv("REDIS_ADDR"),
		Password:     os.Getenv("REDIS_PW"),
		DB:           int(db),
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     5,
		PoolTimeout:  30 * time.Second,
	})

	_, err := client.Ping().Result()

	if err != nil {
		panic(err)
	}

	RedisClient = client
}

// redis的操作需要看看example中的例子
