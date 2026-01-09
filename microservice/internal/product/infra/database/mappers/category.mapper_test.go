package mappers

import (
	"tech_challenge/internal/product/daos"
	"tech_challenge/internal/product/infra/database/models"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFromCategoryDAOToCategoryModel(t *testing.T) {
	dao := daos.CategoryDAO{ID: "catid", Name: "Bebidas", Active: true}
	model := FromCategoryDAOToCategoryModel(dao)
	require.Equal(t, dao.ID, model.ID)
	require.Equal(t, dao.Name, model.Name)
	require.Equal(t, dao.Active, model.Active)
}

func TestFromCategoryModelToCategoryDAO(t *testing.T) {
	model := &models.CategoryModel{ID: "catid", Name: "Bebidas", Active: true}
	dao := FromCategoryModelToCategoryDAO(model)
	require.Equal(t, model.ID, dao.ID)
	require.Equal(t, model.Name, dao.Name)
	require.Equal(t, model.Active, dao.Active)
}

func TestArrayFromCategoryModelToCategoryDAO(t *testing.T) {
	model1 := &models.CategoryModel{ID: "catid1", Name: "Bebidas", Active: true}
	model2 := &models.CategoryModel{ID: "catid2", Name: "Comidas", Active: false}
	arr := ArrayFromCategoryModelToCategoryDAO([]*models.CategoryModel{model1, model2})
	require.Len(t, arr, 2)
	require.Equal(t, "catid1", arr[0].ID)
	require.Equal(t, "catid2", arr[1].ID)
}
