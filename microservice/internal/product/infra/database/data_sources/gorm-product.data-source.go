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
	err := r.db.Create(&productModel).Error
	if err != nil {
		return err
	}

	if len(productDAO.Images) > 0 {
		img := productDAO.Images[0]
		img.ProductID = productModel.ID
		if err := r.AddProductImage(img); err != nil {
			return err
		}
	}
	return nil
}

func (r *GormProductDataSource) FindAll() ([]daos.ProductDAO, error) {
	var products []*models.ProductModel
	err := r.db.Preload("Images", func(db *gorm.DB) *gorm.DB {
		return db.Where("is_default = ?", true).Order("created_at desc")
	}).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return mappers.ArrayFromProductModelToProductDAO(products)
}

func (r *GormProductDataSource) FindAllByCategoryID(categoryID string) ([]daos.ProductDAO, error) {
	var products []*models.ProductModel
	err := r.db.Preload("Images", func(db *gorm.DB) *gorm.DB {
		return db.Where("is_default = ?", true).Order("created_at desc")
	}).Where("category_id = ?", categoryID).Find(&products).Error
	if err != nil {
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

func (r *GormProductDataSource) AddProductImage(productImage daos.ProductImageDAO) error {
	return r.db.Create(&productImage).Error
}

func (r *GormProductDataSource) UpdateProductImage(productImage daos.ProductImageDAO) error {
	return r.db.Model(&daos.ProductImageDAO{}).
		Where("id = ? AND product_id = ?", productImage.ID, productImage.ProductID).
		Updates(map[string]interface{}{
			"is_default": productImage.IsDefault,
		}).Error
}

func (r *GormProductDataSource) SetPreviousImagesAsNotDefault(productID, exceptImageID string) error {
	return r.db.Model(&models.ProductImageModel{}).
		Where("product_id = ? AND id <> ?", productID, exceptImageID).
		Update("is_default", false).Error
}
