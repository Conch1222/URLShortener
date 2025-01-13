package main

import (
	"github.com/deatil/go-encoding/encoding"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func shorten(c *gin.Context) {

	var urlRequest URLRequest
	if err := c.BindJSON(&urlRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortURL := generateShortURL(urlRequest.LongUrl)

	c.JSON(http.StatusOK, gin.H{"shortURL": shortURL})
}

func generateShortURL(longURL string) string {
	var sb strings.Builder
	suffix := encoding.FromString(longURL).Base62Encode().ToString()

	sb.WriteString("http://shortURL/")
	sb.WriteString(suffix)

	return sb.String()
}
