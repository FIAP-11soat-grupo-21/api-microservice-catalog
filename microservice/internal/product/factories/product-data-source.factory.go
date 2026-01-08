package factories

import (
	"tech_challenge/internal/product/infra/database/data_sources"
	"tech_challenge/internal/product/interfaces"
	"tech_challenge/internal/shared/infra/database"
)

func NewProductDataSource() interfaces.IProductDataSource {
	return data_sources.NewProductDataSource(database.GetDB())
}
