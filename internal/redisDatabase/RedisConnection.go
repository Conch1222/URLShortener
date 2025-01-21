package redisDatabase

import (
	"github.com/redis/go-redis/v9"
	"os"
	"sync"
)

type RedisConnection struct {
	Rdb *redis.Client
}

var RedisConn *RedisConnection
var onceRedisConn sync.Once

func ConnectRedis() *RedisConnection {
	onceRedisConn.Do(func() {
		rdb := initRedis()
		RedisConn = rdb
	})
	return RedisConn
}

func initRedis() *RedisConnection {
	if RedisConn == nil {
		redisAddr := os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")
		if redisAddr == ":" {
			redisAddr = "localhost:6379"
		}

		rdb := redis.NewClient(&redis.Options{
			Addr:     redisAddr,
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		})
		return &RedisConnection{Rdb: rdb}
	}
	return RedisConn
}
