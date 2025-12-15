package use_cases

import (
	"tech_challenge/internal/product/application/gateways"
	"tech_challenge/internal/product/domain/entities"
)

type FindAllProductsUseCase struct {
	gateway gateways.ProductGateway
}

func NewFindAllProductsUseCase(gateway gateways.ProductGateway) *FindAllProductsUseCase {
	return &FindAllProductsUseCase{
		gateway: gateway,
	}
}

func (uc *FindAllProductsUseCase) Execute(categoryID *string) ([]entities.Product, error) {
	if categoryID != nil {
		return uc.gateway.FindAllByCategoryID(*categoryID)
	}

	return uc.gateway.FindAll()
}
