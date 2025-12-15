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

func (uc *DeleteProductUseCase) Execute(id string) error {
	product, err := uc.gateway.FindByID(id)

	if err != nil {
		return &exceptions.ProductNotFoundException{}
	}

	err = uc.gateway.Delete(product.ID)

	if err != nil {
		return &exceptions.InvalidProductDataException{}
	}

	return nil
}
