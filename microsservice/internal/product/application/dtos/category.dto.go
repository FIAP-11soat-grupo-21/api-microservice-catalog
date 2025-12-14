package dtos

type CreateCategoryDTO struct {
	Name   string
	Active bool
}

type UpdateCategoryDTO struct {
	ID     string
	Name   string
	Active bool
}

type CategoryResultDTO struct {
	ID     string
	Name   string
	Active bool
}
