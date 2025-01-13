package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func setRouter() {
	r := gin.Default()

	r.POST("/api/shorten", shorten)

	err := r.Run(":8080") //set port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
