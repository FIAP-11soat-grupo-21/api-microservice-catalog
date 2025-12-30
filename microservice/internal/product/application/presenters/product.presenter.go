package presenters

import (
	"tech_challenge/internal/product/application/dtos"
	"tech_challenge/internal/product/domain/entities"
	value_objects "tech_challenge/internal/product/domain/value-objects"
)

func ProductFromDomainToResultDTO(product entities.Product) dtos.ProductResultDTO {
	productImages := make([]dtos.ProductImageDTO, len(product.Images))
	for i, img := range product.Images {
		productImages[i] = productImageFromDomainToDTO(*img)
	}
	return dtos.ProductResultDTO{
		ID:          product.ID,
		Name:        product.Name.Value(),
		Description: product.Description,
		Price:       product.Price.Value(),
		Active:      product.Active,
		CategoryID:  product.CategoryID,
		Images:      productImages,
	}
}

func ListProductDomainToResultDTO(products []entities.Product) []dtos.ProductResultDTO {
	result := make([]dtos.ProductResultDTO, len(products))
	for i, p := range products {
		result[i] = ProductFromDomainToResultDTO(p)
	}
	return result
}

func productImageFromDomainToDTO(image value_objects.Image) dtos.ProductImageDTO {
	return dtos.ProductImageDTO{
		ID:        image.ID,
		FileName:  image.FileName,
		Url:       image.Url,
		IsDefault: image.IsDefault,
	}
}
