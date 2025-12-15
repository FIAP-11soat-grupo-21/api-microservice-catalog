package value_objects

import (
	"tech_challenge/internal/product/domain/exceptions"
)

type Price struct {
	value float64
}

func NewPrice(value float64) (Price, error) {
	if value <= 0 {
		return Price{}, &exceptions.InvalidProductDataException{
			Message: "price must be greater than 0",
		}
	}

	return Price{value: value}, nil
}

func (p *Price) Value() float64 {
	return p.value
}
