package daos

import "time"

type ProductImageDAO struct {
	ID        string
	ProductID string
	FileName  string
	Url       string
	IsDefault bool
	CreatedAt time.Time
}

func (ProductImageDAO) TableName() string {
	return "product_images"
}
