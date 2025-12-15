package product_seed

import (
	"log"
	"tech_challenge/internal/product/daos"
	"tech_challenge/internal/product/interfaces"

	"github.com/brianvoe/gofakeit/v6"
)

func SeedCategories(dataSource interfaces.ICategoryDataSource) {
	categories := []struct {
		Name string
	}{
		{"Lanche"},
		{"Acompanhamento"},
		{"Bebida"},
		{"Sobremesa"},
	}

	for _, c := range categories {
		id := gofakeit.UUID()

		category := daos.CategoryDAO{
			ID:   id,
			Name: c.Name,
		}

		if err := dataSource.Insert(category); err != nil {
			log.Printf("Failed to persist category '%s': %v", c.Name, err)
		}
	}

	log.Println("Categories seeded successfully.")
}
