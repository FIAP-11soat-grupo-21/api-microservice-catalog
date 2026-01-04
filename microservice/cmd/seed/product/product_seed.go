package product_seed

import (
	"log"
	"tech_challenge/internal/product/daos"
	"tech_challenge/internal/product/interfaces"
	"tech_challenge/internal/shared/infra/database"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

func SeedProducts(dataSource interfaces.IProductDataSource) {
	for i := 0; i < 10; i++ {
		id := gofakeit.UUID()
		categoryID := gofakeit.UUID()
		name := gofakeit.ProductName()
		description := gofakeit.ProductDescription()
		price := float64(gofakeit.Number(1, 100))
		active := gofakeit.Bool()

		product := daos.ProductDAO{
			ID:          id,
			CategoryID:  categoryID,
			Name:        name,
			Description: description,
			Price:       price,
			Active:      active,
		}

		if err := dataSource.Insert(product); err != nil {
			log.Printf("Failed to persist product (index %d): %v", i, err)
			continue
		}

		imageID := uuid.NewString()
		fileName := "default_product_image.webp"
		url := "http://minio:9000/product-photo-fiap-tech-challenge-catalog/default_product_image.webp"
		createdAt := time.Now()

		img := map[string]interface{}{
			"id":         imageID,
			"product_id": id,
			"file_name":  fileName,
			"url":        url,
			"is_default": true,
			"created_at": createdAt,
		}
		db := database.GetDB()
		if err := db.Table("product_images").Create(img).Error; err != nil {
			log.Printf("Failed to persist product image (index %d): %v", i, err)
		}
	}

	log.Println("Products and images seeded successfully.")
}
