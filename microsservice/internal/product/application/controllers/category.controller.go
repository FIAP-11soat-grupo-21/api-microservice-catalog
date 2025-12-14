package controllers

import (
	"tech_challenge/internal/product/application/dtos"
	"tech_challenge/internal/product/application/gateways"
	"tech_challenge/internal/product/application/presenters"
	"tech_challenge/internal/product/interfaces"
	use_cases "tech_challenge/internal/product/use_cases/category"
)

type CategoryController struct {
	gateway gateways.CategoryGateway
}

func NewCategoryController(dataSource interfaces.ICategoryDataSource) *CategoryController {
	return &CategoryController{
		gateway: gateways.NewCategoryGateway(dataSource),
	}
}

func (c *CategoryController) Create(categoryDTO dtos.CreateCategoryDTO) (dtos.CategoryResultDTO, error) {
	createCategoryUseCase := use_cases.NewCreateCategoryUseCase(c.gateway)

	category, err := createCategoryUseCase.Execute(categoryDTO.Name, categoryDTO.Active)

	if err != nil {
		return dtos.CategoryResultDTO{}, err
	}

	return presenters.CategoryFromDomainToResultDTO(category), nil
}

func (c *CategoryController) FindByID(id string) (dtos.CategoryResultDTO, error) {
	findCategoryByIDUseCase := use_cases.NewFindCategoryByIDUseCase(c.gateway)

	category, err := findCategoryByIDUseCase.Execute(id)

	if err != nil {
		return dtos.CategoryResultDTO{}, err
	}

	return presenters.CategoryFromDomainToResultDTO(category), nil
}

func (c *CategoryController) FindAll() ([]dtos.CategoryResultDTO, error) {
	findAllCategoryUseCase := use_cases.NewFindAllCategoryUseCase(c.gateway)

	categories, err := findAllCategoryUseCase.Execute()

	if err != nil {
		return []dtos.CategoryResultDTO{}, err
	}

	return presenters.CategoriesFromDomainToResultDTO(categories), nil
}

func (c *CategoryController) Update(categoryDTO dtos.UpdateCategoryDTO) (dtos.CategoryResultDTO, error) {
	updateCategoryUseCase := use_cases.NewUpdateCategoryUseCase(c.gateway)

	category, err := updateCategoryUseCase.Execute(categoryDTO)

	if err != nil {
		return dtos.CategoryResultDTO{}, err
	}

	return presenters.CategoryFromDomainToResultDTO(category), nil
}

func (c *CategoryController) Delete(id string) error {
	deleteCategoryUseCase := use_cases.NewDeleteCategoryUseCase(c.gateway)

	err := deleteCategoryUseCase.Execute(id)

	if err != nil {
		return err
	}

	return nil
}
