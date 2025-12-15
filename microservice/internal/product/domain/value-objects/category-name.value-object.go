package value_objects

import (
	"strings"

	"tech_challenge/internal/product/domain/exceptions"
)

type CategoryName struct {
	value string
}

func NewCategoryName(value string) (CategoryName, error) {
	value = strings.TrimSpace(value)

	if len(value) < 3 {
		return CategoryName{}, &exceptions.InvalidCategoryDataException{
			Message: "category name must have at least 3 characters",
		}
	}

	if len(value) > 100 {
		return CategoryName{}, &exceptions.InvalidCategoryDataException{
			Message: "category name must have at most 100 characters",
		}
	}

	return CategoryName{value: value}, nil
}

func (n CategoryName) Value() string {
	return n.value
}
