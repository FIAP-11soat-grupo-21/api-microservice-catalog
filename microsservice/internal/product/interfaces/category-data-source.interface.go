package interfaces

import (
	"tech_challenge/internal/product/daos"
)

type ICategoryDataSource interface {
	Insert(category daos.CategoryDAO) error
	FindByID(id string) (daos.CategoryDAO, error)
	FindAll() ([]daos.CategoryDAO, error)
	Update(category daos.CategoryDAO) error
	Delete(id string) error
}
