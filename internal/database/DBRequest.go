package database

import (
	"URLShortener/internal/redisDatabase"
	"context"
	"fmt"
	"github.com/deatil/go-encoding/encoding"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

var ctx = context.Background()

func (dbConn *DBConnection) SaveShortURLRecord(longURL string, expiration int64) (string, error) {
	tx, err := dbConn.db.Begin() //begin transaction
	if err != nil {
		return "", err
	}

	res, err := tx.Exec("insert into URL_conversion(long_url, expiration, create_at) values(?, ?, ?);", longURL, expiration, time.Now().UTC())
	if err != nil {
		tx.Rollback()
		return "", err
	}

	id, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return "", err
	}

	shortURL := encoding.FromString(strconv.FormatInt(id, 10)).Base62Encode().ToString()

	_, err = tx.Exec("update URL_conversion set short_url = ? where id = ?", shortURL, id)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	err = tx.Commit()
	if err != nil {
		return "", err
	}

	rdb := redisDatabase.ConnectRedis()
	err = rdb.Rdb.Set(ctx, shortURL, longURL, time.Duration(expiration)*time.Second).Err()
	if err != nil {
		return "", err
	}

	fmt.Println("Short URL stored in Redis with expiration:", expiration)

	return shortURL, nil
}

func (dbConn *DBConnection) DeleteShortURLRecord(shortURL string) error {
	_, err := dbConn.db.Exec("delete from URL_conversion where short_url=?", shortURL)
	if err != nil {
		fmt.Println("Error deleting from DB:", err)
		return err
	} else {
		fmt.Println("Deleted expired URL from DB:", shortURL)
		return nil
	}
}

func (dbConn *DBConnection) HandleExpirationURL() error {
	rdb := redisDatabase.ConnectRedis()
	pubSub := rdb.Rdb.PSubscribe(ctx, "__keyevent@0__:expired")

	fmt.Println("Subscribed to Redis expired events channel")
	for msg := range pubSub.Channel() {
		shortURL := msg.Payload
		fmt.Println("Key expired:", shortURL)

		err := dbConn.DeleteShortURLRecord(shortURL)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dbConn *DBConnection) GetLongURL(shortURL string) (string, error) {
	rdb := redisDatabase.ConnectRedis()
	longURL, err := rdb.Rdb.Get(ctx, shortURL).Result()
	if err == redis.Nil {
		fmt.Println("Key not found in Redis, checking database...")

		var expiration int64
		var createAt time.Time
		err := dbConn.db.QueryRow("SELECT long_url, expiration, CAST(create_at AS datetime) FROM URL_conversion WHERE short_url = ?", shortURL).Scan(&longURL, &expiration, &createAt)
		if err != nil {
			return "", err
		}

		// calculate if expired
		expirationTime := createAt.Add(time.Duration(expiration) * time.Second)
		currentTime := time.Now().UTC()

		if currentTime.After(expirationTime) {
			return "", fmt.Errorf("short URL has expired")
		}

		ttl := time.Until(expirationTime)
		err = rdb.Rdb.Set(ctx, shortURL, longURL, ttl).Err()
		if err != nil {
			return "", err
		}
	} else if err != nil {
		return "", err
	}
	return longURL, nil
}
