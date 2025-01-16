package main

import (
	"URLShortener/internal/database"
	"fmt"
)

func main() {
	db := database.ConnectDB()
	err := db.HandleExpirationURL()
	if err != nil {
		fmt.Println(err)
	}
}
