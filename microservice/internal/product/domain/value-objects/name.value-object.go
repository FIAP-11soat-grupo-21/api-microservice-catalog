package value_objects

import (
	"strings"
	"tech_challenge/internal/product/domain/exceptions"
)

type Name struct {
	value string
}

func NewName(value string) (Name, error) {
	if len(strings.TrimSpace(value)) < 3 {
		return Name{}, &exceptions.InvalidProductDataException{
			Message: "name must be at least 3 characters long",
		}
	}

	if len(strings.TrimSpace(value)) > 100 {
		return Name{}, &exceptions.InvalidProductDataException{
			Message: "name must be at most 100 characters long",
		}
	}

	return Name{value: value}, nil
}

func (n *Name) Value() string {
	return n.value
}
