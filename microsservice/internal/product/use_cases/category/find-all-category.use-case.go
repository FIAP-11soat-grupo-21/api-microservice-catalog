package use_cases

import (
	"tech_challenge/internal/product/application/gateways"
	"tech_challenge/internal/product/domain/entities"
)

type FindAllCategoryUseCase struct {
	gateway gateways.CategoryGateway
}

func NewFindAllCategoryUseCase(gateway gateways.CategoryGateway) *FindAllCategoryUseCase {
	return &FindAllCategoryUseCase{
		gateway: gateway,
	}
}

func (uc *FindAllCategoryUseCase) Execute() ([]*entities.Category, error) {
	categories, err := uc.gateway.FindAll()

	if err != nil {
		return []*entities.Category{}, err
	}

	return categories, nil
}
