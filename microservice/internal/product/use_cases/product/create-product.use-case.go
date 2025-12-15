package use_cases

import (
	"tech_challenge/internal/product/application/dtos"
	"tech_challenge/internal/product/application/gateways"
	"tech_challenge/internal/product/domain/entities"
	identity_manager "tech_challenge/internal/shared/pkg/identity"
)

type CreateProductUseCase struct {
	gateway gateways.ProductGateway
}

func NewCreateProductUseCase(gateway gateways.ProductGateway) *CreateProductUseCase {
	return &CreateProductUseCase{
		gateway: gateway,
	}
}

func (uc *CreateProductUseCase) Execute(productDTO dtos.CreateProductDTO) (entities.Product, error) {
	product, err := entities.NewProduct(
		identity_manager.NewUUIDV4(),
		productDTO.CategoryID,
		productDTO.Name,
		productDTO.Description,
		productDTO.Price,
		productDTO.Active,
	)

	if err != nil {
		return entities.Product{}, err
	}

	err = uc.gateway.Insert(*product)

	if err != nil {
		return entities.Product{}, err
	}

	return *product, nil
}
