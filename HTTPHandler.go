package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func shorten(c *gin.Context) {

	var urlRequest URLRequest
	if err := c.BindJSON(&urlRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortURL := generateShortURL(urlRequest.LongUrl)
	expiration, err := generateExpiration(urlRequest.Expiration)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"shortURL":   shortURL,
		"expiration": expiration,
	})
}

func generateShortURL(shortURLSuffix string) string {
	var sb strings.Builder
	sb.WriteString("http://shortURL/")
	sb.WriteString(shortURLSuffix)

	return sb.String()
}

func generateExpiration(expiration string) (int64, error) {
	duration, err := time.ParseDuration(expiration)
	if err != nil {
		return -1, errors.New("invalid expiration format")
	}
	ttl := int64(duration.Seconds())
	return ttl, nil
}
