package presenters

import (
	"tech_challenge/internal/product/application/dtos"
	"tech_challenge/internal/product/domain/entities"
)

func CategoryFromDomainToResultDTO(category entities.Category) dtos.CategoryResultDTO {
	return dtos.CategoryResultDTO{
		ID:     category.ID,
		Name:   category.Name.Value(),
		Active: category.Active,
	}
}

func CategoriesFromDomainToResultDTO(categories []*entities.Category) []dtos.CategoryResultDTO {
	var result []dtos.CategoryResultDTO

	for _, category := range categories {
		result = append(result, CategoryFromDomainToResultDTO(*category))
	}

	return result
}
