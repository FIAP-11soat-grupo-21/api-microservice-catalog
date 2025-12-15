package gateways

import (
	"tech_challenge/internal/product/daos"
	"tech_challenge/internal/product/domain/entities"
	"tech_challenge/internal/product/interfaces"
)

type CategoryGateway struct {
	dataSource interfaces.ICategoryDataSource
}

func NewCategoryGateway(dataSource interfaces.ICategoryDataSource) CategoryGateway {
	return CategoryGateway{
		dataSource: dataSource,
	}
}

func (g *CategoryGateway) Insert(category entities.Category) error {
	return g.dataSource.Insert(daos.CategoryDAO{
		ID:     category.ID,
		Name:   category.Name.Value(),
		Active: category.Active,
	})
}

func (g *CategoryGateway) FindAll() ([]*entities.Category, error) {
	categories, err := g.dataSource.FindAll()

	if err != nil {
		return nil, err
	}

	result := make([]*entities.Category, 0, len(categories))
	for _, category := range categories {
		categoryEntity, err := entities.NewCategory(
			category.ID,
			category.Name,
			category.Active,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, categoryEntity)
	}

	return result, nil
}

func (g *CategoryGateway) FindByID(id string) (*entities.Category, error) {
	category, err := g.dataSource.FindByID(id)

	if err != nil {
		return nil, err
	}

	categoryEntity, err := entities.NewCategory(
		category.ID,
		category.Name,
		category.Active,
	)

	if err != nil {
		return nil, err
	}

	return categoryEntity, nil
}

func (g *CategoryGateway) Update(category entities.Category) error {
	return g.dataSource.Update(daos.CategoryDAO{
		ID:     category.ID,
		Name:   category.Name.Value(),
		Active: category.Active,
	})
}

func (g *CategoryGateway) Delete(id string) error {
	return g.dataSource.Delete(id)
}
