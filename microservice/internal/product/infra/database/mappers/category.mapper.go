package mappers

import (
	"tech_challenge/internal/product/daos"
	"tech_challenge/internal/product/infra/database/models"
)

func FromCategoryDAOToCategoryModel(category daos.CategoryDAO) models.CategoryModel {
	return models.CategoryModel{
		ID:     category.ID,
		Name:   category.Name,
		Active: category.Active,
	}
}

func FromCategoryModelToCategoryDAO(category *models.CategoryModel) daos.CategoryDAO {
	categoryEntity := daos.CategoryDAO{
		ID:     category.ID,
		Name:   category.Name,
		Active: category.Active,
	}

	return categoryEntity
}

func ArrayFromCategoryModelToCategoryDAO(categories []*models.CategoryModel) []daos.CategoryDAO {
	categoryEntities := make([]daos.CategoryDAO, 0, len(categories))

	for _, category := range categories {
		categoryEntity := FromCategoryModelToCategoryDAO(category)

		categoryEntities = append(categoryEntities, categoryEntity)
	}

	return categoryEntities
}
