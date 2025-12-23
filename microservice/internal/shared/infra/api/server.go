package api

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	product_router "tech_challenge/internal/product/infra/api/routes"
	"tech_challenge/internal/shared/config/env"
	"tech_challenge/internal/shared/infra/api/handlers"
	"tech_challenge/internal/shared/infra/api/middlewares"
	file_router "tech_challenge/internal/shared/infra/api/routes"
	_ "tech_challenge/internal/shared/infra/api/swagger"
	"tech_challenge/internal/shared/infra/database"
)

func Init() {
	config := env.GetConfig()

	if config.IsProduction() {
		log.Printf("Running in production mode on [%s]", config.APIUrl)
		gin.SetMode(gin.ReleaseMode)
	}

	database.Connect()

	if config.Database.RunMigrations {
		database.RunMigrations()
	}

	ginRouter := gin.Default()

	ginRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ginRouter.Use(gin.Logger())
	ginRouter.Use(gin.Recovery())
	ginRouter.Use(middlewares.ErrorHandlerMiddleware())

	// ginRouter.GET("/health", healthHandler.Health)
	// ginRouter.GET("/v1/health", healthHandler.Health)

	v1Routes := ginRouter.Group("/v1")
	healthHandler := handlers.NewHealthHandler()
	v1Routes.GET("/health", healthHandler.Health)

	file_router.RegisterFileRoutes(v1Routes.Group("/uploads"))

	product_router.RegisterProductRoutes(v1Routes.Group("/products"))
	product_router.RegisterCategoryRoutes(v1Routes.Group("/categories"))

	if err := ginRouter.Run(config.APIUrl); err != nil {
		log.Fatalf("failed to start gin server: %v", err)
	}
}
