package schemas

import "tech_challenge/internal/product/application/dtos"

type CreateCategorySchema struct {
	Name   string `json:"name" binding:"required"`
	Active bool   `json:"active"`
}

func (s *CreateCategorySchema) ToDTO() dtos.CreateCategoryDTO {
	return dtos.CreateCategoryDTO{
		Name:   s.Name,
		Active: s.Active,
	}
}

type UpdateCategoryRequestBodySchema struct {
	Name   string `json:"name" binding:"required"`
	Active *bool  `json:"active"`
}

func (s *UpdateCategoryRequestBodySchema) ToDTO(categoryID string) dtos.UpdateCategoryDTO {
	return dtos.UpdateCategoryDTO{
		ID:     categoryID,
		Name:   s.Name,
		Active: *s.Active,
	}
}

type CategoryResponseSchema struct {
	ID     string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Name   string `json:"name" example:"Bebidas"`
	Active bool   `json:"active" example:"true"`
}

func ToCategoryResponseSchema(dto dtos.CategoryResultDTO) CategoryResponseSchema {
	return CategoryResponseSchema{
		ID:     dto.ID,
		Name:   dto.Name,
		Active: dto.Active,
	}
}

func ListToCategoryResponseSchema(dtos []dtos.CategoryResultDTO) []CategoryResponseSchema {
	schemas := make([]CategoryResponseSchema, len(dtos))
	for i, dto := range dtos {
		schemas[i] = ToCategoryResponseSchema(dto)
	}
	return schemas
}

type InvalidCategoryDataErrorSchema struct {
	Error string `json:"error" example:"Invalid category data"`
}

type CategoryNotFoundErrorSchema struct {
	Error string `json:"error" example:"Category not found"`
}
