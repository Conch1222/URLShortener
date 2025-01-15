package redis

import (
	"URLShortener/internal/database"
	"context"
	"fmt"
)

var ctx = context.Background()

func (rdb *RedisConnection) HandleExpirationURL() error {
	db := database.ConnectDB()

	pubsub := rdb.Rdb.PSubscribe(ctx, "__keyevent@0__:expired")

	for msg := range pubsub.Channel() {
		shortURL := msg.Payload
		fmt.Println("Key expired:", shortURL)

		err := db.DeleteShortURLRecord(shortURL)
		if err != nil {
			return err
		}
	}
	return nil
}
