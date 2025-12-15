package presenters

import (
	"tech_challenge/internal/product/application/dtos"
	"tech_challenge/internal/product/domain/entities"
	value_objects "tech_challenge/internal/product/domain/value-objects"
)

func ProductFromDomainToResultDTO(product entities.Product) dtos.ProductResultDTO {
	productImages := make([]dtos.ProductImageDTO, len(product.Images))

	for i, img := range product.Images {
		productImages[i] = productImageFromDomainToDTO(img)
	}

	return dtos.ProductResultDTO{
		ID:          product.ID,
		Name:        product.Name.Value(),
		Description: product.Description,
		Price:       product.Price.Value(),
		Images:      productImages,
		CategoryID:  product.CategoryID,
	}
}

func ListProductDomainToResultDTO(products []entities.Product) []dtos.ProductResultDTO {
	productDTOs := make([]dtos.ProductResultDTO, 0, len(products))

	for _, p := range products {
		productDTOs = append(productDTOs, ProductFromDomainToResultDTO(p))
	}

	return productDTOs
}

func productImageFromDomainToDTO(image value_objects.Image) dtos.ProductImageDTO {
	return dtos.ProductImageDTO{
		FileName: image.FileName,
		Url:      image.Url,
	}
}
