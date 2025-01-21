package http

import (
	"github.com/gin-gonic/gin"
	"log"
)

func SetRouter() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.POST("/api/shorten", shorten)
	r.GET("/shorten", mainPage)
	r.GET("/shorten/:shortURL", redirectShorURL)
	r.GET("/result", result)

	err := r.Run(":8080") //set port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
