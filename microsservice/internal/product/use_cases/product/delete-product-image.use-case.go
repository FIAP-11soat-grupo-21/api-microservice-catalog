package use_cases

import (
	"tech_challenge/internal/product/application/gateways"
	"tech_challenge/internal/product/domain/exceptions"
)

type DeleteProductImageUseCase struct {
	gateway gateways.ProductGateway
}

func NewDeleteProductImageUseCase(gateway gateways.ProductGateway) *DeleteProductImageUseCase {
	return &DeleteProductImageUseCase{
		gateway: gateway,
	}
}

func (uc *DeleteProductImageUseCase) Execute(productID string, imageFileName string) error {
	product, err := uc.gateway.FindByID(productID)

	if err != nil {
		return &exceptions.ProductNotFoundException{}
	}

	imageToBeDeletedIsTheDefault := product.ImageIsDefault(imageFileName)

	err = product.RemoveImage(imageFileName)

	if err != nil {
		return err
	}

	err = uc.gateway.Update(product)

	if err != nil {
		return &exceptions.InvalidProductDataException{}
	}

	if !imageToBeDeletedIsTheDefault {
		err = uc.gateway.DeleteImage(imageFileName)

		if err != nil {
			return &exceptions.InvalidProductImageException{
				Message: "Failed to delete image file from storage",
			}
		}
	}

	return nil
}
