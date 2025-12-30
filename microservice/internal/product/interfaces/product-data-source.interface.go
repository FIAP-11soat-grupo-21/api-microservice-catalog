package interfaces

import "tech_challenge/internal/product/daos"

type IProductDataSource interface {
	Insert(product daos.ProductDAO) error
	FindByID(id string) (daos.ProductDAO, error)
	FindAll() ([]daos.ProductDAO, error)
	FindAllByCategoryID(categoryID string) ([]daos.ProductDAO, error)
	Update(product daos.ProductDAO) error
	Delete(id string) error
	AddProductImage(productImage daos.ProductImageDAO) error
	UpdateProductImage(productImage daos.ProductImageDAO) error
	SetPreviousImagesAsNotDefault(productID, exceptImageID string) error
}
