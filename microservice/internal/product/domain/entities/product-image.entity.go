package entities

import (
	"tech_challenge/internal/product/daos"
	value_objects "tech_challenge/internal/product/domain/value-objects"
	"time"
)

type ProductImage struct {
	ID        string
	ProductID string
	FileName  string
	Url       string
	CreatedAt time.Time
	IsDefault bool
}

func NewProductImage(id, productID string, img value_objects.Image, createdAt time.Time, isDefault bool) ProductImage {
	return ProductImage{
		ID:        id,
		ProductID: productID,
		FileName:  img.FileName,
		Url:       img.Url,
		CreatedAt: createdAt,
		IsDefault: isDefault,
	}
}

func (pi ProductImage) ToDAO() daos.ProductImageDAO {
	return daos.ProductImageDAO{
		ID:        pi.ID,
		ProductID: pi.ProductID,
		FileName:  pi.FileName,
		Url:       pi.Url,
		IsDefault: pi.IsDefault,
		CreatedAt: pi.CreatedAt,
	}
}
