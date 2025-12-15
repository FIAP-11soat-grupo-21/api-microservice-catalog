package use_cases

import (
	"tech_challenge/internal/product/application/gateways"
	"tech_challenge/internal/product/domain/entities"
	"tech_challenge/internal/product/domain/exceptions"
)

type FindCategoryByIDUseCase struct {
	gateway gateways.CategoryGateway
}

func NewFindCategoryByIDUseCase(gateway gateways.CategoryGateway) *FindCategoryByIDUseCase {
	return &FindCategoryByIDUseCase{
		gateway: gateway,
	}
}

func (uc *FindCategoryByIDUseCase) Execute(id string) (entities.Category, error) {
	category, err := uc.gateway.FindByID(id)

	if err != nil {
		return entities.Category{}, &exceptions.CategoryNotFoundException{}
	}

	return *category, nil
}
