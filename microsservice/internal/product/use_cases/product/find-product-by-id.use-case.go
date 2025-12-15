package use_cases

import (
	"tech_challenge/internal/product/application/gateways"
	"tech_challenge/internal/product/domain/entities"
	"tech_challenge/internal/product/domain/exceptions"
)

type FindProductByIDUseCase struct {
	gateway gateways.ProductGateway
}

func NewFindProductByIDUseCase(gateway gateways.ProductGateway) *FindProductByIDUseCase {
	return &FindProductByIDUseCase{
		gateway: gateway,
	}
}

func (uc *FindProductByIDUseCase) Execute(id string) (entities.Product, error) {
	product, err := uc.gateway.FindByID(id)

	if err != nil || product.IsEmpty() {
		return entities.Product{}, &exceptions.ProductNotFoundException{}
	}

	return product, nil
}
