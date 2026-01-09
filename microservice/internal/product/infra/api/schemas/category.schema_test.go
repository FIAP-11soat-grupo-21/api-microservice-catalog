package schemas

import (
	"tech_challenge/internal/product/application/dtos"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateCategorySchema_ToDTO(t *testing.T) {
	schema := CreateCategorySchema{Name: "Bebidas", Active: true}
	dto := schema.ToDTO()
	require.Equal(t, schema.Name, dto.Name)
	require.Equal(t, schema.Active, dto.Active)
}

func TestUpdateCategoryRequestBodySchema_ToDTO(t *testing.T) {
	active := true
	schema := UpdateCategoryRequestBodySchema{Name: "Bebidas", Active: &active}
	dto := schema.ToDTO("catid")
	require.Equal(t, "catid", dto.ID)
	require.Equal(t, schema.Name, dto.Name)
	require.Equal(t, *schema.Active, dto.Active)
}

func TestToCategoryResponseSchema(t *testing.T) {
	dto := dtos.CategoryResultDTO{ID: "catid", Name: "Bebidas", Active: true}
	resp := ToCategoryResponseSchema(dto)
	require.Equal(t, dto.ID, resp.ID)
	require.Equal(t, dto.Name, resp.Name)
	require.Equal(t, dto.Active, resp.Active)
}

func TestListToCategoryResponseSchema(t *testing.T) {
	dtosArr := []dtos.CategoryResultDTO{
		{ID: "catid1", Name: "Bebidas", Active: true},
		{ID: "catid2", Name: "Comidas", Active: false},
	}
	resp := ListToCategoryResponseSchema(dtosArr)
	require.Len(t, resp, 2)
	require.Equal(t, "catid1", resp[0].ID)
	require.Equal(t, "catid2", resp[1].ID)
}
