package http

import (
	"URLShortener/internal/database"
	"URLShortener/type"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func shorten(c *gin.Context) {

	var urlRequest _type.URLRequest
	if err := c.BindJSON(&urlRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.ConnectDB()
	expiration, err := generateExpiration(urlRequest.Expiration)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortURL, err := db.SaveShortURLRecord(urlRequest.LongUrl, expiration)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortURL = generateShortURL(shortURL)

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
