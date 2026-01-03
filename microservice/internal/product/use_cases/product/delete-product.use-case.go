package use_cases

import (
	"tech_challenge/internal/product/application/gateways"
	"tech_challenge/internal/product/domain/exceptions"
)

type DeleteProductUseCase struct {
	gateway gateways.ProductGateway
}

func NewDeleteProductUseCase(gateway gateways.ProductGateway) *DeleteProductUseCase {
	return &DeleteProductUseCase{
		gateway: gateway,
	}
}

func (uc *DeleteProductUseCase) Execute(productID string) error {
	_, err := uc.gateway.FindByID(productID)
	if err != nil {
		return &exceptions.ProductNotFoundException{}
	}

	product_images, err := uc.gateway.FindAllImagesProductById(productID)
	if err != nil {
		return &exceptions.ProductImagesNotFoundException{}
	}
	err = uc.gateway.DeleteFiles(product_images.Images)
	if err != nil {
		return &exceptions.DeleteImagesStorageException{}
	}
	err = uc.gateway.Delete(productID)
	if err != nil {
		return err
	}

	return nil
}
