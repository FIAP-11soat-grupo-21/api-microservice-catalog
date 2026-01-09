package data_sources

import (
	"fmt"

	"gorm.io/gorm"

	"tech_challenge/internal/product/daos"
	database_errors "tech_challenge/internal/product/infra/database/database_errors"
	"tech_challenge/internal/product/infra/database/mappers"
	"tech_challenge/internal/product/infra/database/models"
)

type GormCategoryDataSource struct {
	db *gorm.DB
}

func NewGormCategoryDataSource(db *gorm.DB) *GormCategoryDataSource {
	return &GormCategoryDataSource{db: db}
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
	fmt.Printf("[GormCategoryDataSource] Chamando Delete para categoria id: %s\n", id)
	result := r.db.Delete(&models.CategoryModel{}, "id = ?", id)
	if result.Error != nil {
		fmt.Printf("[GormCategoryDataSource] Erro ao deletar categoria: %v\n", result.Error)
		errTratado := database_errors.HandleDatabaseErrors(result.Error)
		fmt.Printf("[GormCategoryDataSource] Resultado do HandleDatabaseErrors: %v\n", errTratado)
		return errTratado
	}
	fmt.Println("[GormCategoryDataSource] Categoria deletada com sucesso")
	return nil
}
