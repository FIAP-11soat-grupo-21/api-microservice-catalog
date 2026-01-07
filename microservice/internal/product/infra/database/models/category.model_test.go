package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCategoryModel_TableName(t *testing.T) {
	model := CategoryModel{}
	require.Equal(t, "category", model.TableName())
}

func TestCategoryModel_Fields(t *testing.T) {
	model := CategoryModel{ID: "id1", Name: "Bebidas", Active: true}
	require.Equal(t, "id1", model.ID)
	require.Equal(t, "Bebidas", model.Name)
	require.True(t, model.Active)
}
