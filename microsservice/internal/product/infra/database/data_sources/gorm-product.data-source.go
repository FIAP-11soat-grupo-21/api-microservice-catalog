package data_sources

import (
	"gorm.io/gorm"

	"tech_challenge/internal/product/daos"
	"tech_challenge/internal/product/infra/database/mappers"
	"tech_challenge/internal/product/infra/database/models"
	"tech_challenge/internal/shared/infra/database"
)

type GormProductDataSource struct {
	db *gorm.DB
}

func NewProductDataSource() *GormProductDataSource {
	return &GormProductDataSource{
		db: database.GetDB(),
	}
}

func (r *GormProductDataSource) Insert(productDAO daos.ProductDAO) error {
	productModel := mappers.FromProductDAOToProductModel(productDAO)

	return r.db.Model(&models.ProductModel{}).Create(&productModel).Error
}

func (r *GormProductDataSource) FindAll() ([]daos.ProductDAO, error) {
	var products []*models.ProductModel

	if err := r.db.Preload("Images").Find(&products).Error; err != nil {
		return nil, err
	}

	return mappers.ArrayFromProductModelToProductDAO(products)
}

func (r *GormProductDataSource) FindAllByCategoryID(categoryID string) ([]daos.ProductDAO, error) {
	var products []*models.ProductModel

	if err := r.db.Preload("Images").Where("category_id = ?", categoryID).Find(&products).Error; err != nil {
		return nil, err
	}

	return mappers.ArrayFromProductModelToProductDAO(products)
}

func (r *GormProductDataSource) FindByID(id string) (daos.ProductDAO, error) {
	var product *models.ProductModel

	if err := r.db.Preload("Images").First(&product, "id = ?", id).Error; err != nil {
		return daos.ProductDAO{}, err
	}

	return mappers.FromProductModelToProductDAO(product)
}

func (r *GormProductDataSource) Update(product daos.ProductDAO) error {
	return r.db.Save(mappers.FromProductDAOToProductModel(product)).Error
}

func (r *GormProductDataSource) Delete(id string) error {
	return r.db.Delete(&models.ProductModel{}, "id = ?", id).Error
}
