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

func mainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "MainPage.html", gin.H{})
}

func shorten(c *gin.Context) {

	var urlRequest _type.URLRequest
	urlRequest.LongUrl = c.PostForm("long_url")
	urlRequest.Expiration = c.PostForm("expiration")

	if urlRequest.LongUrl == "" || urlRequest.Expiration == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Lack of necessary data"})
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

func redirectShorURL(c *gin.Context) {
	db := database.ConnectDB()

	shortURL := c.Param("shortURL")
	longURL, err := db.GetLongURL(shortURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(302, longURL)
}

func generateShortURL(shortURLSuffix string) string {
	var sb strings.Builder
	sb.WriteString("http://localhost:8080/shorten/")
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
