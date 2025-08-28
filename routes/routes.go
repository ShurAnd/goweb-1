package routes

import (
	"gintest/handlers"
	"gintest/storage"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	recipeStorage := storage.InitInMemoryRecipeStorage()
	recipeHandler := handlers.InitRecipeHandler(recipeStorage)

	router.GET("/recipes", recipeHandler.GetRecipes)

	router.GET("/recipes/:id", recipeHandler.GetRecipeById)

	router.POST("/recipes", recipeHandler.CreateNewRecipe)

	router.GET("/hello", recipeHandler.HelloHandler)

	return router
}
