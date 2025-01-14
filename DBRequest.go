package main

import (
	"github.com/deatil/go-encoding/encoding"
	"strconv"
)

func (dbConn *DBConnection) saveShortURLRecord(longURL string, expiration int64) (string, error) {
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

	return generateShortURL(shortURL), nil
}
