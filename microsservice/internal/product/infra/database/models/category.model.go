package models

type CategoryModel struct {
	ID     string `gorm:"primaryKey; size:36"`
	Name   string `gorm:"not null;size:100;"`
	Active bool   `gorm:"not null;"`
}

func (CategoryModel) TableName() string {
	return "category"
}
