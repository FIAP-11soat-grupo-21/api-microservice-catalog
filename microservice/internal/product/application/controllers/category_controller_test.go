package controllers

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"tech_challenge/internal/product/application/dtos"
	"tech_challenge/internal/product/daos"
	testmocks "tech_challenge/internal/shared/test"
)

func TestCategoryController_Create_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDS := &testmocks.MockCategoryDataSource{
		InsertFunc: func(dao daos.CategoryDAO) error { return nil },
	}
	c := NewCategoryController(mockDS)
	dto := dtos.CreateCategoryDTO{Name: "Bebidas", Active: true}
	res, err := c.Create(dto)
	require.NoError(t, err)
	require.Equal(t, "Bebidas", res.Name)
}

func TestCategoryController_Create_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDS := &testmocks.MockCategoryDataSource{
		InsertFunc: func(dao daos.CategoryDAO) error { return errors.New("fail") },
	}
	c := NewCategoryController(mockDS)
	dto := dtos.CreateCategoryDTO{Name: "Bebidas", Active: true}
	_, err := c.Create(dto)
	require.Error(t, err)
}

func TestCategoryController_FindByID_Success(t *testing.T) {
	mockDS := &testmocks.MockCategoryDataSource{
		FindByIDFunc: func(id string) (daos.CategoryDAO, error) {
			return daos.CategoryDAO{ID: id, Name: "Bebidas", Active: true}, nil
		},
	}
	c := NewCategoryController(mockDS)
	res, err := c.FindByID("catid")
	require.NoError(t, err)
	require.Equal(t, "catid", res.ID)
}

func TestCategoryController_FindByID_Error(t *testing.T) {
	mockDS := &testmocks.MockCategoryDataSource{
		FindByIDFunc: func(id string) (daos.CategoryDAO, error) { return daos.CategoryDAO{}, errors.New("fail") },
	}
	c := NewCategoryController(mockDS)
	_, err := c.FindByID("catid")
	require.Error(t, err)
}

func TestCategoryController_FindAll_Success(t *testing.T) {
	mockDS := &testmocks.MockCategoryDataSource{
		FindAllFunc: func() ([]daos.CategoryDAO, error) {
			return []daos.CategoryDAO{{ID: "catid", Name: "Bebidas", Active: true}}, nil
		},
	}
	c := NewCategoryController(mockDS)
	res, err := c.FindAll()
	require.NoError(t, err)
	require.Len(t, res, 1)
	require.Equal(t, "catid", res[0].ID)
}

func TestCategoryController_FindAll_Error(t *testing.T) {
	mockDS := &testmocks.MockCategoryDataSource{
		FindAllFunc: func() ([]daos.CategoryDAO, error) { return nil, errors.New("fail") },
	}
	c := NewCategoryController(mockDS)
	_, err := c.FindAll()
	require.Error(t, err)
}

func TestCategoryController_Update_Success(t *testing.T) {
	mockDS := &testmocks.MockCategoryDataSource{
		UpdateFunc: func(dao daos.CategoryDAO) error { return nil },
		FindByIDFunc: func(id string) (daos.CategoryDAO, error) {
			return daos.CategoryDAO{ID: id, Name: "Bebidas", Active: true}, nil
		},
	}
	c := NewCategoryController(mockDS)
	dto := dtos.UpdateCategoryDTO{ID: "catid", Name: "Bebidas", Active: true}
	res, err := c.Update(dto)
	require.NoError(t, err)
	require.Equal(t, "catid", res.ID)
}

func TestCategoryController_Update_Error(t *testing.T) {
	mockDS := &testmocks.MockCategoryDataSource{
		UpdateFunc: func(dao daos.CategoryDAO) error { return errors.New("fail") },
		FindByIDFunc: func(id string) (daos.CategoryDAO, error) {
			return daos.CategoryDAO{ID: id, Name: "Bebidas", Active: true}, nil
		},
	}
	c := NewCategoryController(mockDS)
	dto := dtos.UpdateCategoryDTO{ID: "catid", Name: "Bebidas", Active: true}
	_, err := c.Update(dto)
	require.Error(t, err)
}

func TestCategoryController_Delete_Success(t *testing.T) {
	mockDS := &testmocks.MockCategoryDataSource{
		DeleteFunc: func(id string) error { return nil },
		FindByIDFunc: func(id string) (daos.CategoryDAO, error) {
			return daos.CategoryDAO{ID: id, Name: "Bebidas", Active: true}, nil
		},
	}
	c := NewCategoryController(mockDS)
	require.NoError(t, c.Delete("catid"))
}

func TestCategoryController_Delete_Error(t *testing.T) {
	mockDS := &testmocks.MockCategoryDataSource{
		DeleteFunc: func(id string) error { return errors.New("fail") },
		FindByIDFunc: func(id string) (daos.CategoryDAO, error) {
			return daos.CategoryDAO{ID: id, Name: "Bebidas", Active: true}, nil
		},
	}
	c := NewCategoryController(mockDS)
	require.Error(t, c.Delete("catid"))
}
