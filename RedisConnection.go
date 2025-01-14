package main

import (
	"github.com/redis/go-redis/v9"
	"sync"
)

type RedisConnection struct {
	rdb *redis.Client
}

var RedisConn *RedisConnection
var onceRedisConn sync.Once

func connectRedis() *RedisConnection {
	onceRedisConn.Do(func() {
		rdb := initRedis()
		RedisConn = rdb
	})
	return RedisConn
}

func initRedis() *RedisConnection {
	if RedisConn == nil {
		rdb := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
		return &RedisConnection{rdb: rdb}
	}
	return RedisConn
}
