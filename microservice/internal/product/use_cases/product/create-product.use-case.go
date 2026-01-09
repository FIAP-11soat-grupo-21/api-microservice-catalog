package use_cases

import (
	"tech_challenge/internal/product/application/dtos"
	"tech_challenge/internal/product/application/gateways"
	"tech_challenge/internal/product/domain/entities"
	"tech_challenge/internal/product/domain/exceptions"
	identity_manager "tech_challenge/internal/shared/pkg/identity"
)

type CreateProductUseCase struct {
	productGateway  gateways.ProductGateway
	categoryGateway gateways.CategoryGateway
}

func NewCreateProductUseCase(productGateway gateways.ProductGateway, categoryGateway gateways.CategoryGateway) *CreateProductUseCase {
	return &CreateProductUseCase{
		productGateway:  productGateway,
		categoryGateway: categoryGateway,
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

	_, err = uc.categoryGateway.FindByID(product.CategoryID)
	if err != nil {
		return entities.Product{}, &exceptions.CategoryNotFoundException{}
	}

	err = uc.productGateway.Insert(*product)
	if err != nil {
		return entities.Product{}, err
	}

	return *product, nil
}
