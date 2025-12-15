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

type ProductDAO struct {
	ID          string
	CategoryID  string
	Name        string
	Description string
	Price       float64
	Images      []ProductImageDAO
	Active      bool
}
