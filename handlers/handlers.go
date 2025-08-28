package handlers

import (
	"fmt"
	"gintest/models"
	"gintest/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RecipeHandler struct {
	storage storage.RecipeStorage
}

func InitRecipeHandler(s storage.RecipeStorage) *RecipeHandler {
	return &RecipeHandler{
		storage: s,
	}
}

func (s *RecipeHandler) GetRecipes(c *gin.Context) {
	fmt.Println("hello")
	c.JSON(200, gin.H{
		"recipes": s.storage.GetAll(),
	})
}

func (s *RecipeHandler) GetRecipeById(c *gin.Context) {
	id, iderr := strconv.Atoi(c.Param("id"))
	if iderr != nil {
		id = -1
	}
	rec, err := s.storage.GetById(id)
	if err != nil {
		c.JSON(404, gin.H{})
		return
	}

	c.JSON(200, gin.H{
		"recipe": rec,
	})

	return
}

func (s *RecipeHandler) CreateNewRecipe(c *gin.Context) {
	var recipe models.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err := s.storage.Create(&recipe)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"recipe": recipe,
	})
}

func (s *RecipeHandler) HelloHandler(context *gin.Context) {
	clientIP := context.ClientIP()

	context.String(http.StatusOK, "Hello %s", clientIP)
}
