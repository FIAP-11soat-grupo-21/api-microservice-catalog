package main

import (
	"log"

	product_seed "tech_challenge/cmd/seed/product"
	product_factory "tech_challenge/internal/product/factories"
	"tech_challenge/internal/shared/config/env"
	"tech_challenge/internal/shared/infra/database"
)

func main() {
	config := env.GetConfig()

	config.Database.Host = "localhost"

	database.Connect()

	productDataSource := product_factory.NewProductDataSource()
	categoryDataSource := product_factory.NewCategoryDataSource()

	product_seed.SeedCategories(categoryDataSource)
	product_seed.SeedProducts(productDataSource)

	log.Println("Seed completed successfully!")
}
