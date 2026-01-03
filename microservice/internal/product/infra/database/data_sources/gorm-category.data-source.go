package data_sources

import (
	"gorm.io/gorm"

	"tech_challenge/internal/product/daos"
	database_errors "tech_challenge/internal/product/infra/database/database_errors"
	"tech_challenge/internal/product/infra/database/mappers"
	"tech_challenge/internal/product/infra/database/models"
	"tech_challenge/internal/shared/infra/database"
)

type GormCategoryDataSource struct {
	db *gorm.DB
}

func NewGormCategoryDataSource() *GormCategoryDataSource {
	return &GormCategoryDataSource{
		db: database.GetDB(),
	}
}

func (r *GormCategoryDataSource) Insert(category daos.CategoryDAO) error {
	categoryModel := mappers.FromCategoryDAOToCategoryModel(category)

	return r.db.Model(&models.CategoryModel{}).Create(&categoryModel).Error
}

func (r *GormCategoryDataSource) FindAll() ([]daos.CategoryDAO, error) {
	var categories []*models.CategoryModel

	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}

	return mappers.ArrayFromCategoryModelToCategoryDAO(categories), nil
}

func (r *GormCategoryDataSource) FindByID(id string) (daos.CategoryDAO, error) {
	var category *models.CategoryModel

	if err := r.db.First(&category, "id = ?", id).Error; err != nil {
		return daos.CategoryDAO{}, err
	}

	return mappers.FromCategoryModelToCategoryDAO(category), nil
}

func (r *GormCategoryDataSource) Update(category daos.CategoryDAO) error {
	return r.db.Save(mappers.FromCategoryDAOToCategoryModel(category)).Error
}

func (r *GormCategoryDataSource) Delete(id string) error {
	result := r.db.Delete(&models.CategoryModel{}, "id = ?", id)
	if result.Error != nil {
		return database_errors.HandleDatabaseErrors(result.Error)
	}
	return nil
}
