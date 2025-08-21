package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Recipe struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt"`
}

func main() {
	router := gin.Default()

	router.GET("/hello", func(context *gin.Context) {
		clientIP := context.ClientIP()

		context.String(http.StatusOK, "Hello %s", clientIP)
	})

	router.Run(":9999")
}
