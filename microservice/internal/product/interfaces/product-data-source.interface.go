package interfaces

import (
	"tech_challenge/internal/product/daos"
)

type IProductDataSource interface {
	Insert(product daos.ProductDAO) error
	Update(product daos.ProductDAO) error
	Delete(id string) error
	FindAll() ([]daos.ProductDAO, error)
	FindByID(id string) (daos.ProductDAO, error)
	FindAllByCategoryID(categoryID string) ([]daos.ProductDAO, error)
	FindAllImagesProductById(productID string) ([]daos.ProductImageDAO, error)
	AddProductImage(productImage daos.ProductImageDAO) error
	SetAllPreviousImagesAsNotDefault(productID, exceptImageID string) error
	SetImageAsDefault(productID, imageID string) error
	DeleteImage(imageFileName string) error
}
