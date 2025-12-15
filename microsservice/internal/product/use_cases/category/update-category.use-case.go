package use_cases

import (
	"tech_challenge/internal/product/application/dtos"
	"tech_challenge/internal/product/application/gateways"
	"tech_challenge/internal/product/domain/entities"
	"tech_challenge/internal/product/domain/exceptions"
)

type UpdateCategoryUseCase struct {
	gateway gateways.CategoryGateway
}

func NewUpdateCategoryUseCase(gateway gateways.CategoryGateway) *UpdateCategoryUseCase {
	return &UpdateCategoryUseCase{
		gateway: gateway,
	}
}

func (uc *UpdateCategoryUseCase) Execute(categoryDTO dtos.UpdateCategoryDTO) (entities.Category, error) {
	category, err := uc.gateway.FindByID(categoryDTO.ID)

	if err != nil {
		return entities.Category{}, &exceptions.CategoryNotFoundException{}
	}

	if err = category.SetName(categoryDTO.Name); err != nil {
		return entities.Category{}, err
	}

	category.Active = categoryDTO.Active

	err = uc.gateway.Update(*category)

	if err != nil {
		return entities.Category{}, &exceptions.InvalidCategoryDataException{}
	}

	return *category, nil
}
