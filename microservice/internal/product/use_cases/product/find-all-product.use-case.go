package use_cases

import (
	"tech_challenge/internal/product/application/gateways"
	"tech_challenge/internal/product/domain/entities"
	"tech_challenge/internal/product/domain/exceptions"
)

type FindAllProductsUseCase struct {
	gateway         gateways.ProductGateway
	categoryGateway gateways.CategoryGateway
}

func NewFindAllProductsUseCase(gateway gateways.ProductGateway, categoryGateway gateways.CategoryGateway) *FindAllProductsUseCase {
	return &FindAllProductsUseCase{
		gateway:         gateway,
		categoryGateway: categoryGateway,
	}
}

func (uc *FindAllProductsUseCase) Execute(categoryID *string) ([]entities.Product, error) {
	if categoryID != nil {
		_, err := uc.categoryGateway.FindByID(*categoryID)
		if err != nil {
			return nil, &exceptions.CategoryNotFoundException{}
		}
		return uc.gateway.FindAllByCategoryID(*categoryID)
	}
	return uc.gateway.FindAll()
}
