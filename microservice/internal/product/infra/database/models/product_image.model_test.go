package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestProductImageModel_TableName(t *testing.T) {
	model := ProductImageModel{}
	require.Equal(t, "product_images", model.TableName())
}

func TestProductImageModel_Fields(t *testing.T) {
	created := time.Now()
	model := ProductImageModel{
		ID:        "imgid",
		ProductID: "pid",
		FileName:  "img.jpg",
		Url:       "http://host/img.jpg",
		IsDefault: true,
		CreatedAt: created,
	}
	require.Equal(t, "imgid", model.ID)
	require.Equal(t, "pid", model.ProductID)
	require.Equal(t, "img.jpg", model.FileName)
	require.Equal(t, "http://host/img.jpg", model.Url)
	require.True(t, model.IsDefault)
	require.Equal(t, created, model.CreatedAt)
}
