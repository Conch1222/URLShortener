package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

const NETWORK = "tcp"

type DBConnection struct {
	db *sql.DB
}

var DBConn *DBConnection
var onceDBConn sync.Once

func ConnectDB() *DBConnection {
	onceDBConn.Do(func() {
		db := initDB()
		DBConn = db
	})
	return DBConn
}

func initDB() *DBConnection {
	if DBConn == nil {
		username := "admin"
		password := "admin"
		server := "localhost"
		port := "3306"
		database := "web_URL_Shortener"

		fmt.Println("Connecting to " + database)
		conn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", username, password, NETWORK, server, port, database)
		db, err := sql.Open("mysql", conn)
		if err != nil {
			fmt.Println("Open Mysql error: ", err)
			return nil
		}
		if err := db.Ping(); err != nil {
			fmt.Println("database connect error: ", err.Error())
			return nil
		}

		return &DBConnection{db: db}
	}
	return DBConn
}
