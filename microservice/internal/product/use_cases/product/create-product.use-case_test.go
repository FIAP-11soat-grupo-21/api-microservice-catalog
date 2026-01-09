package use_cases

import (
	"errors"
	"os"
	"testing"

	"tech_challenge/internal/product/application/dtos"
	"tech_challenge/internal/product/application/gateways"
	"tech_challenge/internal/product/daos"
	"tech_challenge/internal/product/domain/exceptions"
	mock_interfaces "tech_challenge/internal/product/interfaces/mocks"
	testenv "tech_challenge/internal/shared/test"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	testenv.SetupTestEnv()
	code := m.Run()
	os.Exit(code)
}

func setupCreateProductTest(t *testing.T, name string) (dtos.CreateProductDTO, *mock_interfaces.MockIProductDataSource, *mock_interfaces.MockICategoryDataSource, *mock_interfaces.MockIFileProvider, string, *gomock.Controller) {
	ctrl := gomock.NewController(t)
	mockProductDataSource := mock_interfaces.NewMockIProductDataSource(ctrl)
	mockCategoryDataSource := mock_interfaces.NewMockICategoryDataSource(ctrl)
	mockFileProvider := mock_interfaces.NewMockIFileProvider(ctrl)
	categoryID := "cat-1"
	productDTO := dtos.CreateProductDTO{
		CategoryID:  categoryID,
		Name:        name,
		Description: "Descrição",
		Price:       10.0,
		Active:      true,
	}
	return productDTO, mockProductDataSource, mockCategoryDataSource, mockFileProvider, categoryID, ctrl
}

func TestCreateProductUseCase_Success(t *testing.T) {
	productDTO, mockProductDataSource, mockCategoryDataSource, mockFileProvider, categoryID, ctrl := setupCreateProductTest(t, "Produto Teste")
	defer ctrl.Finish()
	mockCategoryDataSource.EXPECT().FindByID(categoryID).Return(daos.CategoryDAO{ID: categoryID, Name: "Categoria Teste", Active: true}, nil)
	mockProductDataSource.EXPECT().Insert(gomock.Any()).Return(nil)
	categoryGateway := gateways.NewCategoryGateway(mockCategoryDataSource)
	productGateway := gateways.NewProductGateway(mockProductDataSource, mockFileProvider)
	uc := NewCreateProductUseCase(*productGateway, categoryGateway)
	product, err := uc.Execute(productDTO)
	require.NoError(t, err)
	require.Equal(t, productDTO.Name, product.Name.Value())
	require.Equal(t, productDTO.CategoryID, product.CategoryID)
	require.Equal(t, productDTO.Description, product.Description)
	require.Equal(t, productDTO.Price, product.Price.Value())
	require.Equal(t, productDTO.Active, product.Active)
}

func TestCreateProductUseCase_CategoryNotFound(t *testing.T) {
	productDTO, mockProductDataSource, mockCategoryDataSource, mockFileProvider, categoryID, ctrl := setupCreateProductTest(t, "Produto Teste")
	defer ctrl.Finish()
	mockCategoryDataSource.EXPECT().FindByID(categoryID).Return(daos.CategoryDAO{}, errors.New("not found"))
	categoryGateway := gateways.NewCategoryGateway(mockCategoryDataSource)
	productGateway := gateways.NewProductGateway(mockProductDataSource, mockFileProvider)
	uc := NewCreateProductUseCase(*productGateway, categoryGateway)
	_, err := uc.Execute(productDTO)
	_, ok := err.(*exceptions.CategoryNotFoundException)
	require.True(t, ok)
}

func TestCreateProductUseCase_InvalidProduct(t *testing.T) {
	productDTO, mockProductDataSource, mockCategoryDataSource, mockFileProvider, _, ctrl := setupCreateProductTest(t, "")
	defer ctrl.Finish()
	categoryGateway := gateways.NewCategoryGateway(mockCategoryDataSource)
	productGateway := gateways.NewProductGateway(mockProductDataSource, mockFileProvider)
	uc := NewCreateProductUseCase(*productGateway, categoryGateway)
	_, err := uc.Execute(productDTO)
	require.Error(t, err)
}

func TestCreateProductUseCase_InsertError(t *testing.T) {
	productDTO, mockProductDataSource, mockCategoryDataSource, mockFileProvider, categoryID, ctrl := setupCreateProductTest(t, "Produto Teste")
	defer ctrl.Finish()
	mockCategoryDataSource.EXPECT().FindByID(categoryID).Return(daos.CategoryDAO{ID: categoryID, Name: "Categoria Teste", Active: true}, nil)
	mockProductDataSource.EXPECT().Insert(gomock.Any()).Return(errors.New("insert error"))
	categoryGateway := gateways.NewCategoryGateway(mockCategoryDataSource)
	productGateway := gateways.NewProductGateway(mockProductDataSource, mockFileProvider)
	uc := NewCreateProductUseCase(*productGateway, categoryGateway)
	_, err := uc.Execute(productDTO)
	require.EqualError(t, err, "insert error")
}
