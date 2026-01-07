package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestProductModel_TableName(t *testing.T) {
	model := ProductModel{}
	require.Equal(t, "products", model.TableName())
}

func TestProductModel_Fields(t *testing.T) {
	created := time.Now()
	cat := CategoryModel{ID: "catid", Name: "Bebidas", Active: true}
	img := ProductImageModel{ID: "imgid", ProductID: "pid", FileName: "img.jpg", Url: "http://host/img.jpg", IsDefault: true, CreatedAt: created}
	model := ProductModel{
		ID:          "pid",
		CategoryID:  "catid",
		Category:    cat,
		Name:        "Coca-Cola",
		Description: "desc",
		Price:       5.99,
		Active:      true,
		CreatedAt:   created,
		Images:      []ProductImageModel{img},
	}
	require.Equal(t, "pid", model.ID)
	require.Equal(t, "catid", model.CategoryID)
	require.Equal(t, cat, model.Category)
	require.Equal(t, "Coca-Cola", model.Name)
	require.Equal(t, "desc", model.Description)
	require.Equal(t, 5.99, model.Price)
	require.True(t, model.Active)
	require.Equal(t, created, model.CreatedAt)
	require.Len(t, model.Images, 1)
	require.Equal(t, img, model.Images[0])
}
