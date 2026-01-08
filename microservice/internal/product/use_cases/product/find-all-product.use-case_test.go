package use_cases

import (
	"errors"
	"testing"

	"tech_challenge/internal/product/application/gateways"
	"tech_challenge/internal/product/daos"
	"tech_challenge/internal/product/domain/exceptions"
	mock_interfaces "tech_challenge/internal/product/interfaces/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestFindAllProductsUseCase_Success_WithCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockProductDataSource := mock_interfaces.NewMockIProductDataSource(ctrl)
	mockCategoryDataSource := mock_interfaces.NewMockICategoryDataSource(ctrl)
	mockFileProvider := mock_interfaces.NewMockIFileProvider(ctrl)

	categoryID := "cat-1"
	mockCategoryDataSource.EXPECT().FindByID(categoryID).Return(daos.CategoryDAO{ID: categoryID, Name: "Categoria Teste", Active: true}, nil)
	mockProductDataSource.EXPECT().FindAllByCategoryID(categoryID).Return(
		[]daos.ProductDAO{
			{ID: "pid", Name: "Coca-Cola", CategoryID: categoryID, Price: 5.99, Active: true, Images: []daos.ProductImageDAO{}},
		},
		nil,
	)
	categoryGateway := gateways.NewCategoryGateway(mockCategoryDataSource)
	productGateway := gateways.NewProductGateway(mockProductDataSource, mockFileProvider)
	uc := NewFindAllProductsUseCase(*productGateway, categoryGateway)

	products, err := uc.Execute(&categoryID)
	require.NoError(t, err)
	require.Len(t, products, 1)
	require.Equal(t, "pid", products[0].ID)
}

func TestFindAllProductsUseCase_Success_WithoutCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockProductDataSource := mock_interfaces.NewMockIProductDataSource(ctrl)
	mockCategoryDataSource := mock_interfaces.NewMockICategoryDataSource(ctrl)
	mockFileProvider := mock_interfaces.NewMockIFileProvider(ctrl)

	mockProductDataSource.EXPECT().FindAll().Return(
		[]daos.ProductDAO{
			{ID: "pid", Name: "Coca-Cola", CategoryID: "cat-1", Price: 5.99, Active: true, Images: []daos.ProductImageDAO{}},
		},
		nil,
	)
	categoryGateway := gateways.NewCategoryGateway(mockCategoryDataSource)
	productGateway := gateways.NewProductGateway(mockProductDataSource, mockFileProvider)
	uc := NewFindAllProductsUseCase(*productGateway, categoryGateway)

	products, err := uc.Execute(nil)
	require.NoError(t, err)
	require.Len(t, products, 1)
	require.Equal(t, "pid", products[0].ID)
	require.Equal(t, "cat-1", products[0].CategoryID)
}

func TestFindAllProductsUseCase_Error_CategoryNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockProductDataSource := mock_interfaces.NewMockIProductDataSource(ctrl)
	mockCategoryDataSource := mock_interfaces.NewMockICategoryDataSource(ctrl)
	mockFileProvider := mock_interfaces.NewMockIFileProvider(ctrl)

	categoryID := "cat-1"
	// Simula erro ao buscar categoria
	mockCategoryDataSource.EXPECT().FindByID(categoryID).Return(daos.CategoryDAO{}, errors.New("not found"))

	categoryGateway := gateways.NewCategoryGateway(mockCategoryDataSource)
	productGateway := gateways.NewProductGateway(mockProductDataSource, mockFileProvider)
	uc := NewFindAllProductsUseCase(*productGateway, categoryGateway)

	products, err := uc.Execute(&categoryID)
	_, ok := err.(*exceptions.CategoryNotFoundException)
	require.True(t, ok)
	require.Nil(t, products)
}
