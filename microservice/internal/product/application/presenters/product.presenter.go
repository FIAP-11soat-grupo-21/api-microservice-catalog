package presenters

import (
	"tech_challenge/internal/product/application/dtos"
	"tech_challenge/internal/product/domain/entities"
	value_objects "tech_challenge/internal/product/domain/value-objects"
)

func ProductFromDomainToResultDTO(product entities.Product) dtos.ProductResultDTO {
	productImages := make([]dtos.ProductImageDTO, len(product.Images))
	for i, img := range product.Images {
		productImages[i] = ProductImageFromDomainToDTO(*img)
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

func ProductImagesFromDomainToResultDTO(images []*value_objects.Image) []dtos.ProductImageDTO {
	imagesResult := make([]dtos.ProductImageDTO, len(images))
	for i, img := range images {
		imagesResult[i] = dtos.ProductImageDTO{
			ID:        img.ID,
			FileName:  img.FileName,
			Url:       img.Url,
			IsDefault: img.IsDefault,
		}
	}
	return imagesResult
}

func ProductImageFromDomainToDTO(img value_objects.Image) dtos.ProductImageDTO {
	return dtos.ProductImageDTO{
		ID:        img.ID,
		FileName:  img.FileName,
		Url:       img.Url,
		IsDefault: img.IsDefault,
	}
}
