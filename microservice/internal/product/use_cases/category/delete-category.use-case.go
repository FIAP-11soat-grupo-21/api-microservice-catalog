package use_cases

import (
	"tech_challenge/internal/product/application/gateways"
	"tech_challenge/internal/product/domain/exceptions"
)

type DeleteCategoryUseCase struct {
	gateway gateways.CategoryGateway
}

func NewDeleteCategoryUseCase(gateway gateways.CategoryGateway) *DeleteCategoryUseCase {
	return &DeleteCategoryUseCase{
		gateway: gateway,
	}
}

func (uc *DeleteCategoryUseCase) Execute(id string) error {
	category, err := uc.gateway.FindByID(id)

	if err != nil {
		return &exceptions.CategoryNotFoundException{}
	}

	err = uc.gateway.Delete(category.ID)

	if err != nil {
		return err
	}

	return nil
}
