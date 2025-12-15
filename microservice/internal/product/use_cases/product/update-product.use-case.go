package use_cases

import (
	"tech_challenge/internal/product/application/dtos"
	"tech_challenge/internal/product/application/gateways"
	"tech_challenge/internal/product/domain/entities"
	"tech_challenge/internal/product/domain/exceptions"
)

type UpdateProductUseCase struct {
	gateway gateways.ProductGateway
}

func NewUpdateProductUseCase(gateway gateways.ProductGateway) *UpdateProductUseCase {
	return &UpdateProductUseCase{
		gateway: gateway,
	}
}

func (uc *UpdateProductUseCase) Execute(productDTO dtos.UpdateProductDTO) (entities.Product, error) {
	product, err := uc.gateway.FindByID(productDTO.ID)

	if err != nil {
		return entities.Product{}, &exceptions.ProductNotFoundException{}
	}

	if err = product.SetName(productDTO.Name); err != nil {
		return entities.Product{}, err
	}

	if err = product.SetDescription(productDTO.Description); err != nil {
		return entities.Product{}, err
	}

	if err = product.SetPrice(productDTO.Price); err != nil {
		return entities.Product{}, err
	}

	if err = product.SetCategory(productDTO.CategoryID); err != nil {
		return entities.Product{}, err
	}

	if productDTO.Active {
		if err := product.Activate(); err != nil {
			return entities.Product{}, err
		}
	} else {
		if err := product.Deactivate(); err != nil {
			return entities.Product{}, err
		}
	}

	err = uc.gateway.Update(product)

	if err != nil {
		return entities.Product{}, &exceptions.InvalidProductDataException{}
	}

	return product, nil
}
