package models

import (
	"time"
)

type ProductModel struct {
	ID          string              `gorm:"primaryKey; size:36"`
	CategoryID  string              `gorm:"not null;size:100;index"`
	Category    CategoryModel       `gorm:"foreignKey:CategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Name        string              `gorm:"not null;size:100;"`
	Description string              `gorm:"not null;"`
	Price       float64             `gorm:"not null; decimal(10,4);"`
	Active      bool                `gorm:"not null;"`
	CreatedAt   time.Time           `gorm:"autoCreateTime"`
	Images      []ProductImageModel `gorm:"foreignKey:ProductID"`
}

func (ProductModel) TableName() string {
	return "products"
}
