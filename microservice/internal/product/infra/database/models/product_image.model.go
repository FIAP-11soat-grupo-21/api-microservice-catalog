package models

import "time"

// ProductImageModel representa uma imagem de produto no banco de dados
// Cada imagem tem um ProductID como chave estrangeira
//
type ProductImageModel struct {
	ID        string    `gorm:"primaryKey;size:36"`
	ProductID string    `gorm:"not null;index;constraint:OnDelete:CASCADE"`
	FileName  string    `gorm:"not null;size:255"`
	Url       string    `gorm:"not null;size:2048"`
	IsDefault bool      `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (ProductImageModel) TableName() string {
	return "product_images"
}
