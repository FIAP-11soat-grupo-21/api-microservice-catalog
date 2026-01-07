package mappers

import (
	"tech_challenge/internal/product/daos"
	"tech_challenge/internal/product/infra/database/models"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFromProductDAOToProductModel(t *testing.T) {
	da := daos.ProductDAO{
		ID:          "pid",
		CategoryID:  "catid",
		Name:        "Coca-Cola",
		Description: "desc",
		Price:       5.99,
		Active:      true,
	}
	model := FromProductDAOToProductModel(da)
	require.Equal(t, da.ID, model.ID)
	require.Equal(t, da.CategoryID, model.CategoryID)
	require.Equal(t, da.Name, model.Name)
	require.Equal(t, da.Description, model.Description)
	require.Equal(t, da.Price, model.Price)
	require.True(t, model.Active)
}

func TestFromProductModelToProductDAO(t *testing.T) {
	created := time.Now()
	img := models.ProductImageModel{ID: "imgid", ProductID: "pid", FileName: "img.jpg", Url: "http://host/img.jpg", IsDefault: true, CreatedAt: created}
	model := &models.ProductModel{
		ID:          "pid",
		CategoryID:  "catid",
		Name:        "Coca-Cola",
		Description: "desc",
		Price:       5.99,
		Active:      true,
		Images:      []models.ProductImageModel{img},
	}
	da, err := FromProductModelToProductDAO(model)
	require.NoError(t, err)
	require.Equal(t, model.ID, da.ID)
	require.Equal(t, model.CategoryID, da.CategoryID)
	require.Equal(t, model.Name, da.Name)
	require.Equal(t, model.Description, da.Description)
	require.Equal(t, model.Price, da.Price)
	require.True(t, da.Active)
	require.Len(t, da.Images, 1)
	require.Equal(t, img.ID, da.Images[0].ID)
}

func TestArrayFromProductModelToProductDAO(t *testing.T) {
	model1 := &models.ProductModel{ID: "pid1", CategoryID: "catid", Name: "Coca-Cola", Description: "desc", Price: 5.99, Active: true}
	model2 := &models.ProductModel{ID: "pid2", CategoryID: "catid", Name: "Pepsi", Description: "desc2", Price: 4.99, Active: false}
	arr, err := ArrayFromProductModelToProductDAO([]*models.ProductModel{model1, model2})
	require.NoError(t, err)
	require.Len(t, arr, 2)
	require.Equal(t, "pid1", arr[0].ID)
	require.Equal(t, "pid2", arr[1].ID)
}
