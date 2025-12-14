package routes

import (
	"tech_challenge/internal/product/infra/api/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(router *gin.RouterGroup) {
	productHandler := handlers.NewProductHandler()

	router.POST("/", productHandler.CreateProduct)
	router.GET("/", productHandler.FindAllProducts)
	router.GET("/:id", productHandler.FindProductByID)
	router.PUT("/:id", productHandler.UpdateProduct)
	router.PATCH("/:id/images", productHandler.UploadProductImage)
	router.DELETE("/:id/images/:image_file_name", productHandler.DeleteProductImage)
	router.DELETE("/:id", productHandler.DeleteProduct)
}
