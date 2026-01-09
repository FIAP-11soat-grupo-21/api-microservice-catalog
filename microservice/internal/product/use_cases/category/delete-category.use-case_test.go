package use_cases_test

import (
	"testing"

	"tech_challenge/internal/product/application/gateways"
	"tech_challenge/internal/product/daos"
	mock_interfaces "tech_challenge/internal/product/interfaces/mocks"
	category "tech_challenge/internal/product/use_cases/category"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestDeleteCategoryUseCase_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategoryDataSource := mock_interfaces.NewMockICategoryDataSource(ctrl)
	categoryID := "cat-1"
	mockCategoryDataSource.EXPECT().FindByID(categoryID).Return(daos.CategoryDAO{ID: categoryID, Name: "Categoria Teste", Active: true}, nil)
	mockCategoryDataSource.EXPECT().Delete(gomock.Any()).Return(nil)

	categoryGateway := gateways.NewCategoryGateway(mockCategoryDataSource)
	uc := category.NewDeleteCategoryUseCase(categoryGateway)
	err := uc.Execute(categoryID)
	require.NoError(t, err)
}
