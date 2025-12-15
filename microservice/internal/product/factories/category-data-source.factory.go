package factories

import (
	"tech_challenge/internal/product/infra/database/data_sources"
	"tech_challenge/internal/product/interfaces"
)

func NewCategoryDataSource() interfaces.ICategoryDataSource {
	return data_sources.NewGormCategoryDataSource()
}
