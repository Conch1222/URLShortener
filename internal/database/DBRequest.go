package database

import (
	"URLShortener/internal/http"
	"URLShortener/internal/redis"
	"context"
	"fmt"
	"github.com/deatil/go-encoding/encoding"
	"strconv"
	"time"
)

var ctx = context.Background()

func (dbConn *DBConnection) SaveShortURLRecord(longURL string, expiration int64) (string, error) {
	tx, err := dbConn.db.Begin() //begin transaction
	if err != nil {
		return "", err
	}

	res, err := tx.Exec("insert into URL_conversion(long_url, expiration, create_at) values(?, ?, NOW());", longURL, expiration)
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

	rdb := redis.ConnectRedis()
	err = rdb.Rdb.Set(ctx, shortURL, longURL, time.Duration(expiration)*time.Second).Err()
	if err != nil {
		return "", err
	}

	fmt.Println("Short URL stored in Redis with expiration:", expiration)

	return http.GenerateShortURL(shortURL), nil
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
