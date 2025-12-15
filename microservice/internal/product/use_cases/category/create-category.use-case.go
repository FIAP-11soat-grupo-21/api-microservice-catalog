package use_cases

import (
	"tech_challenge/internal/product/application/gateways"
	"tech_challenge/internal/product/domain/entities"
	identity_manager "tech_challenge/internal/shared/pkg/identity"
)

type CreateCategoryUseCase struct {
	gateway gateways.CategoryGateway
}

func NewCreateCategoryUseCase(gateway gateways.CategoryGateway) *CreateCategoryUseCase {
	return &CreateCategoryUseCase{
		gateway: gateway,
	}
}

func (uc *CreateCategoryUseCase) Execute(name string, active bool) (entities.Category, error) {
	category, err := entities.NewCategory(
		identity_manager.NewUUIDV4(),
		name,
		active,
	)

	if err != nil {
		return entities.Category{}, err
	}

	err = uc.gateway.Insert(*category)

	if err != nil {
		return entities.Category{}, err
	}

	return *category, nil
}
