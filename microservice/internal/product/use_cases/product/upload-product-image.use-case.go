package use_cases

import (
	"tech_challenge/internal/product/application/dtos"
	"tech_challenge/internal/product/application/gateways"
	"tech_challenge/internal/product/domain/exceptions"
)

type UploadProductImageUseCase struct {
	gateway gateways.ProductGateway
}

func NewUploadProductImageUseCase(gateway gateways.ProductGateway) *UploadProductImageUseCase {
	return &UploadProductImageUseCase{
		gateway: gateway,
	}
}

func (uc *UploadProductImageUseCase) Execute(productDTO dtos.UploadProductImageDTO) error {
	product, err := uc.gateway.FindByID(productDTO.ProductID)
	if err != nil {
		return &exceptions.ProductNotFoundException{}
	}

	newFileName, err := product.AddImage(productDTO.FileName)
	if err != nil {
		return &exceptions.InvalidProductImageException{}
	}

	url, err := uc.gateway.UploadImage(*newFileName, productDTO.FileContent)
	if err != nil {
		return err
	}

	if err := uc.gateway.AddAndSetDefaultImage(product, url); err != nil {
		return &exceptions.InvalidProductDataException{}
	}

	return nil
}
