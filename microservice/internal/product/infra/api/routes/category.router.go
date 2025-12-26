package routes

import (
	"tech_challenge/internal/product/infra/api/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterCategoryRoutes(router *gin.RouterGroup) {
	categoryHandler := handlers.NewCategoryHandler()

	router.GET("", categoryHandler.FindAllCategories)
	router.GET("/:id", categoryHandler.FindCategoryByID)
	router.POST("", categoryHandler.CreateCategory)
	router.PUT("/:id", categoryHandler.UpdateCategory)
	router.DELETE("/:id", categoryHandler.DeleteCategory)
}
