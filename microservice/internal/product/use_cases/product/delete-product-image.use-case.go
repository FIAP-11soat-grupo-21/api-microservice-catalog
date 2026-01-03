package use_cases

import (
	"fmt"
	"tech_challenge/internal/product/application/gateways"
	"tech_challenge/internal/product/domain/exceptions"
	value_objects "tech_challenge/internal/product/domain/value-objects"
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
	_, err := uc.gateway.FindByID(productID)
	if err != nil {
		return &exceptions.ProductNotFoundException{}
	}
	productImages, err := uc.gateway.FindAllImagesProductById(productID)
	if err != nil || len(productImages.Images) == 0 {
		return &exceptions.ProductImagesNotFoundException{}
	}

	if len(productImages.Images) == 1 {
		return &exceptions.ProductImageCannotBeEmptyException{}
	}

	isDefault := productImages.ImageIsDefault(imageFileName)

	if isDefault {
		err = uc.gateway.SetLastImageAsDefault(productID, imageFileName)
		if err != nil {
			return err
		}
	}

	err = uc.gateway.DeleteProductImage(imageFileName)
	if err != nil {
		return &exceptions.InvalidProductImageException{Message: "Failed to delete image from database"}
	}

	if imageFileName != value_objects.DEFAULT_IMAGE_FILE_NAME {
		err = uc.gateway.DeleteImage(imageFileName)
		if err != nil {

			return fmt.Errorf("failed to delete file from bucket: %w", err)
		}
	}
	return nil
}
