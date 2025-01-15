package worker

import (
	"URLShortener/internal/redis"
	"fmt"
)

func main() {
	rdb := redis.ConnectRedis()
	err := rdb.HandleExpirationURL()
	if err != nil {
		fmt.Println(err)
	}
}
