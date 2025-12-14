package mappers

import (
	"tech_challenge/internal/product/daos"
	"tech_challenge/internal/product/infra/database/models"
)

func FromProductDAOToProductModel(product daos.ProductDAO) *models.ProductModel {
	images := make([]models.ProductImageModel, 0, len(product.Images))
	for _, image := range product.Images {
		images = append(images, models.ProductImageModel{
			FileName: image.FileName,
			Url:      image.Url,
		})
	}

	return &models.ProductModel{
		ID:          product.ID,
		CategoryID:  product.CategoryID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Active:      product.Active,
		Images:      images,
	}
}

func FromProductModelToProductDAO(product *models.ProductModel) (daos.ProductDAO, error) {
	images := make([]daos.ProductImageDAO, len(product.Images))
	for i, img := range product.Images {
		images[i] = daos.ProductImageDAO{
			FileName: img.FileName,
			Url:      img.Url,
		}
	}

	productDAO := daos.ProductDAO{
		ID:          product.ID,
		CategoryID:  product.CategoryID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Images:      images,
		Active:      product.Active,
	}
	return productDAO, nil
}

func ArrayFromProductModelToProductDAO(products []*models.ProductModel) ([]daos.ProductDAO, error) {
	productsEntities := make([]daos.ProductDAO, 0, len(products))

	for _, product := range products {
		productEntity, err := FromProductModelToProductDAO(product)

		if err != nil {
			return []daos.ProductDAO{}, err
		}

		productsEntities = append(productsEntities, productEntity)
	}

	return productsEntities, nil
}
