package models

import (
	"time"
)

type ProductModel struct {
	ID          string              `gorm:"primaryKey; size:36"`
	CategoryID  string              `gorm:"not null;size:100;"`
	Name        string              `gorm:"not null;size:100;"`
	Description string              `gorm:"not null;"`
	Price       float64             `gorm:"not null; decimal(10,4);"`
	Active      bool                `gorm:"not null;"`
	CreatedAt   time.Time           `gorm:"autoCreateTime"`
	UpdatedAt   time.Time           `gorm:"autoUpdateTime"`
	Images      []ProductImageModel `gorm:"foreignKey:ProductID"`
}

func (ProductModel) TableName() string {
	return "products"
}
