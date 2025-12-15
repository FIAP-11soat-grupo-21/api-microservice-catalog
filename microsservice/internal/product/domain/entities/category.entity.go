package entities

import value_objects "tech_challenge/internal/product/domain/value-objects"

type Category struct {
	ID     string
	Name   value_objects.CategoryName
	Active bool
}

func NewCategory(id, name string, active bool) (*Category, error) {
	categoryName, err := value_objects.NewCategoryName(name)

	if err != nil {
		return nil, err
	}

	return &Category{
		ID:     id,
		Name:   categoryName,
		Active: active,
	}, nil
}
func (c *Category) SetName(name string) error {
	newName, err := value_objects.NewCategoryName(name)
	if err != nil {
		return err
	}

	c.Name = newName
	return nil
}
