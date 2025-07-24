package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	router.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": "Hello " + name + "!",
		})
	})

	router.Run()
}
