package presenters

import (
	"tech_challenge/internal/product/domain/entities"
	value_objects "tech_challenge/internal/product/domain/value-objects"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCategoryFromDomainToResultDTO(t *testing.T) {
	catName, _ := value_objects.NewCategoryName("Bebidas")
	cat := entities.Category{
		ID:     "catid",
		Name:   catName,
		Active: true,
	}
	dto := CategoryFromDomainToResultDTO(cat)
	require.Equal(t, "catid", dto.ID)
	require.Equal(t, "Bebidas", dto.Name)
	require.True(t, dto.Active)
}

func TestCategoriesFromDomainToResultDTO(t *testing.T) {
	catName1, _ := value_objects.NewCategoryName("Bebidas")
	catName2, _ := value_objects.NewCategoryName("Comidas")
	cat1 := &entities.Category{ID: "1", Name: catName1, Active: true}
	cat2 := &entities.Category{ID: "2", Name: catName2, Active: false}
	list := []*entities.Category{cat1, cat2}
	dtos := CategoriesFromDomainToResultDTO(list)
	require.Len(t, dtos, 2)
	require.Equal(t, "1", dtos[0].ID)
	require.Equal(t, "Bebidas", dtos[0].Name)
	require.True(t, dtos[0].Active)
	require.Equal(t, "2", dtos[1].ID)
	require.Equal(t, "Comidas", dtos[1].Name)
	require.False(t, dtos[1].Active)
}
