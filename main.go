package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

var recipes []Recipe

func init() {
	recipes = []Recipe{Recipe{
		ID:           1,
		Name:         "Soup",
		Tags:         []string{"soup", "pervoe"},
		Ingredients:  []string{"water", "chicken"},
		Instructions: []string{"get water", "put chicken", "get fire"},
		PublishedAt:  time.Now(),
	}, Recipe{
		ID:           2,
		Name:         "Salad",
		Tags:         []string{"salad", "zakuska"},
		Ingredients:  []string{"cucumber", "tomato"},
		Instructions: []string{"cut cucumber", "cut tomato", "mess"},
		PublishedAt:  time.Now(),
	}}
}

func main() {
	router := gin.Default()
	router.GET("/recipes", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"recipes": recipes,
		})
	})

	router.GET("/recipes/:id", func(c *gin.Context) {
		id, iderr := strconv.Atoi(c.Param("id"))
		if iderr != nil {
			id = -1
		}
		for _, rec := range recipes {
			if rec.ID == id {
				c.JSON(200, gin.H{
					"recipe": rec,
				})
				return
			}
		}

		c.JSON(404, gin.H{})

	})

	router.POST("/recipes", func(c *gin.Context) {
		var recipe Recipe
		if err := c.ShouldBindJSON(&recipe); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		recipe.ID = getNextID()
		recipe.PublishedAt = time.Now()
		recipes = append(recipes, recipe)
		c.JSON(http.StatusCreated, gin.H{
			"recipe": recipe,
		})
	})

	router.GET("/hello", func(context *gin.Context) {
		clientIP := context.ClientIP()

		context.String(http.StatusOK, "Hello %s", clientIP)
	})

	router.Run(":9999")
}

func getNextID() int {
	result := 0
	if len(recipes) == 0 {
		return result
	}
	result = recipes[0].ID
	for _, item := range recipes {
		if item.ID > result {
			result = item.ID
		}
	}
	return result + 1
}
