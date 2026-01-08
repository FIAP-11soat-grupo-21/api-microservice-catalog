package data_sources

import (
	"gorm.io/gorm"

	"tech_challenge/internal/product/daos"
	"tech_challenge/internal/product/infra/database/mappers"
	"tech_challenge/internal/product/infra/database/models"
)

type GormProductDataSource struct {
	db *gorm.DB
}

//	func NewProductDataSource(db *gorm.DB) *GormProductDataSource {
//		return &GormProductDataSource{
//			db: database.GetDB(),
//		}
//	}
func NewProductDataSource(db *gorm.DB) *GormProductDataSource {
	return &GormProductDataSource{db: db}
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

	if err := r.db.Preload("Images", "is_default = ?", true).First(&product, "id = ?", id).Error; err != nil {
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

func (r *GormProductDataSource) SetAllPreviousImagesAsNotDefault(productID, exceptImageID string) error {
	return r.db.Model(&models.ProductImageModel{}).
		Where("product_id = ? AND id <> ?", productID, exceptImageID).
		Update("is_default", false).Error
}

func (r *GormProductDataSource) FindAllImagesProductById(productID string) ([]daos.ProductImageDAO, error) {
	var images []models.ProductImageModel
	err := r.db.Where("product_id = ?", productID).Order("created_at desc").Find(&images).Error
	if err != nil {
		return nil, err
	}
	var result []daos.ProductImageDAO
	for _, img := range images {
		result = append(result, daos.ProductImageDAO{
			ID:        img.ID,
			ProductID: img.ProductID,
			FileName:  img.FileName,
			Url:       img.Url,
			IsDefault: img.IsDefault,
			CreatedAt: img.CreatedAt,
		})
	}
	return result, nil
}

func (r *GormProductDataSource) SetImageAsDefault(productID, imageID string) error {
	// Primeiro, seta todas as imagens do produto como não default
	err := r.db.Model(&models.ProductImageModel{}).
		Where("product_id = ?", productID).
		Update("is_default", false).Error
	if err != nil {
		return err
	}
	// Agora, seta a imagem escolhida como default
	return r.db.Model(&models.ProductImageModel{}).
		Where("product_id = ? AND id = ?", productID, imageID).
		Update("is_default", true).Error
}

func (r *GormProductDataSource) DeleteImage(imageFileName string) error {
	return r.db.Where("file_name = ?", imageFileName).Delete(&models.ProductImageModel{}).Error
}
