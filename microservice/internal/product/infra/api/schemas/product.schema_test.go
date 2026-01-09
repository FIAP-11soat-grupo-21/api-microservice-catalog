package schemas

import (
	"tech_challenge/internal/product/application/dtos"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateProductSchema_ToDTO(t *testing.T) {
	schema := CreateProductSchema{
		CategoryID:  "catid",
		Name:        "Coca-Cola",
		Description: "desc",
		Price:       5.99,
		Active:      true,
	}
	dto := schema.ToDTO()
	require.Equal(t, schema.CategoryID, dto.CategoryID)
	require.Equal(t, schema.Name, dto.Name)
	require.Equal(t, schema.Description, dto.Description)
	require.Equal(t, schema.Price, dto.Price)
	require.Equal(t, schema.Active, dto.Active)
}

func TestUpdateProductRequestBodySchema_ToDTO(t *testing.T) {
	schema := UpdateProductRequestBodySchema{
		CategoryID:  "catid",
		Name:        "Coca-Cola",
		Description: "desc",
		Price:       5.99,
		Active:      true,
	}
	dto := schema.ToDTO("pid")
	require.Equal(t, "pid", dto.ID)
	require.Equal(t, schema.CategoryID, dto.CategoryID)
	require.Equal(t, schema.Name, dto.Name)
	require.Equal(t, schema.Description, dto.Description)
	require.Equal(t, schema.Price, dto.Price)
	require.Equal(t, schema.Active, dto.Active)
}

func TestToProductResponseSchema(t *testing.T) {
	product := dtos.ProductResultDTO{
		ID:          "pid",
		Name:        "Coca-Cola",
		Description: "desc",
		Price:       5.99,
		Active:      true,
		CategoryID:  "catid",
		Images:      []dtos.ProductImageDTO{{FileName: "img.jpg", Url: "http://host/img.jpg"}},
	}
	resp := ToProductResponseSchema(product)
	require.Equal(t, product.ID, resp.ID)
	require.Equal(t, product.Name, resp.Name)
	require.Equal(t, product.Description, resp.Description)
	require.Equal(t, product.Price, resp.Price)
	require.Equal(t, product.Active, resp.Active)
	require.Equal(t, product.CategoryID, resp.CategoryID)
	require.Len(t, resp.Images, 1)
	require.Equal(t, "img.jpg", resp.Images[0].FileName)
}

func TestListProductsResponseSchema(t *testing.T) {
	products := []dtos.ProductResultDTO{
		{ID: "pid1", Name: "Coca-Cola", Description: "desc", Price: 5.99, Active: true, CategoryID: "catid", Images: []dtos.ProductImageDTO{}},
		{ID: "pid2", Name: "Pepsi", Description: "desc2", Price: 4.99, Active: false, CategoryID: "catid", Images: []dtos.ProductImageDTO{}},
	}
	resp := ListProductsResponseSchema(products)
	require.Len(t, resp, 2)
	require.Equal(t, "pid1", resp[0].ID)
	require.Equal(t, "pid2", resp[1].ID)
}
