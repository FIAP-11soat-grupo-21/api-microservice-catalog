package gateways

import (
	"errors"
	"tech_challenge/internal/product/daos"
	"tech_challenge/internal/product/domain/entities"
	"testing"

	"github.com/stretchr/testify/require"
)

type mockCategoryDataSource struct {
	insertFunc   func(dao daos.CategoryDAO) error
	findAllFunc  func() ([]daos.CategoryDAO, error)
	findByIDFunc func(id string) (daos.CategoryDAO, error)
	updateFunc   func(dao daos.CategoryDAO) error
	deleteFunc   func(id string) error
}

func (m *mockCategoryDataSource) Insert(dao daos.CategoryDAO) error {
	return m.insertFunc(dao)
}
func (m *mockCategoryDataSource) FindAll() ([]daos.CategoryDAO, error) {
	return m.findAllFunc()
}
func (m *mockCategoryDataSource) FindByID(id string) (daos.CategoryDAO, error) {
	return m.findByIDFunc(id)
}
func (m *mockCategoryDataSource) Update(dao daos.CategoryDAO) error {
	return m.updateFunc(dao)
}
func (m *mockCategoryDataSource) Delete(id string) error {
	return m.deleteFunc(id)
}

func TestCategoryGateway_Insert(t *testing.T) {
	gw := NewCategoryGateway(&mockCategoryDataSource{
		insertFunc: func(dao daos.CategoryDAO) error { return nil },
	})
	cat, _ := entities.NewCategory("id", "Bebidas", true)
	require.NoError(t, gw.Insert(*cat))
}

func TestCategoryGateway_FindAll_Success(t *testing.T) {
	gw := NewCategoryGateway(&mockCategoryDataSource{
		findAllFunc: func() ([]daos.CategoryDAO, error) {
			return []daos.CategoryDAO{{ID: "id", Name: "Bebidas", Active: true}}, nil
		},
	})
	cats, err := gw.FindAll()
	require.NoError(t, err)
	require.Len(t, cats, 1)
	require.Equal(t, "id", cats[0].ID)
}

func TestCategoryGateway_FindAll_Error(t *testing.T) {
	gw := NewCategoryGateway(&mockCategoryDataSource{
		findAllFunc: func() ([]daos.CategoryDAO, error) { return nil, errors.New("fail") },
	})
	cats, err := gw.FindAll()
	require.Error(t, err)
	require.Nil(t, cats)
}

func TestCategoryGateway_FindAll_Error_Entity(t *testing.T) {
	gw := NewCategoryGateway(&mockCategoryDataSource{
		findAllFunc: func() ([]daos.CategoryDAO, error) {
			return []daos.CategoryDAO{{ID: "", Name: "", Active: true}}, nil
		},
	})
	cats, err := gw.FindAll()
	require.Error(t, err)
	require.Nil(t, cats)
}

func TestCategoryGateway_FindByID_Success(t *testing.T) {
	gw := NewCategoryGateway(&mockCategoryDataSource{
		findByIDFunc: func(id string) (daos.CategoryDAO, error) {
			return daos.CategoryDAO{ID: "id", Name: "Bebidas", Active: true}, nil
		},
	})
	cat, err := gw.FindByID("id")
	require.NoError(t, err)
	require.Equal(t, "id", cat.ID)
}

func TestCategoryGateway_FindByID_Error(t *testing.T) {
	gw := NewCategoryGateway(&mockCategoryDataSource{
		findByIDFunc: func(id string) (daos.CategoryDAO, error) { return daos.CategoryDAO{}, errors.New("fail") },
	})
	cat, err := gw.FindByID("id")
	require.Error(t, err)
	require.Nil(t, cat)
}

func TestCategoryGateway_FindByID_Error_Entity(t *testing.T) {
	gw := NewCategoryGateway(&mockCategoryDataSource{
		findByIDFunc: func(id string) (daos.CategoryDAO, error) {
			return daos.CategoryDAO{ID: "", Name: "", Active: true}, nil
		},
	})
	cat, err := gw.FindByID("id")
	require.Error(t, err)
	require.Nil(t, cat)
}

func TestCategoryGateway_Update(t *testing.T) {
	gw := NewCategoryGateway(&mockCategoryDataSource{
		updateFunc: func(dao daos.CategoryDAO) error { return nil },
	})
	cat, _ := entities.NewCategory("id", "Bebidas", true)
	require.NoError(t, gw.Update(*cat))
}

func TestCategoryGateway_Delete(t *testing.T) {
	gw := NewCategoryGateway(&mockCategoryDataSource{
		deleteFunc: func(id string) error { return nil },
	})
	require.NoError(t, gw.Delete("id"))
}
