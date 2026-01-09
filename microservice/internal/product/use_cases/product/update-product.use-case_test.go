package use_cases_test

import (
	"tech_challenge/internal/product/application/dtos"
	"tech_challenge/internal/product/application/gateways"
	"tech_challenge/internal/product/daos"
	"tech_challenge/internal/product/domain/exceptions"
	mock_interfaces "tech_challenge/internal/product/interfaces/mocks"
	use_cases "tech_challenge/internal/product/use_cases/product"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func makeProductDTO(id, categoryID, name, description string, price float64, active bool) dtos.UpdateProductDTO {
	return dtos.UpdateProductDTO{
		ID:          id,
		CategoryID:  categoryID,
		Name:        name,
		Description: description,
		Price:       price,
		Active:      active,
	}
}

func TestUpdateProductUseCase_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockProductDataSource := mock_interfaces.NewMockIProductDataSource(ctrl)
	mockFileProvider := mock_interfaces.NewMockIFileProvider(ctrl)
	categoryID := "cat-1"
	productDTO := makeProductDTO("pid", categoryID, "Produto Teste", "Descrição", 10.0, true)
	mockProductDataSource.EXPECT().FindByID("pid").Return(daos.ProductDAO{ID: "pid", CategoryID: categoryID, Name: "Produto Teste", Description: "Descrição", Price: 10.0, Active: true}, nil)
	mockProductDataSource.EXPECT().Update(gomock.Any()).Return(nil)
	productGateway := gateways.NewProductGateway(mockProductDataSource, mockFileProvider)
	uc := use_cases.NewUpdateProductUseCase(*productGateway)
	_, err := uc.Execute(productDTO)
	require.NoError(t, err)
}

func TestUpdateProductUseCase_ProductNotFoundException(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockProductDataSource := mock_interfaces.NewMockIProductDataSource(ctrl)
	mockFileProvider := mock_interfaces.NewMockIFileProvider(ctrl)
	categoryID := "cat-1"
	productDTO := makeProductDTO("pid", categoryID, "Produto Teste", "Descrição", 10.0, true)
	mockProductDataSource.EXPECT().FindByID("pid").Return(daos.ProductDAO{}, &exceptions.ProductNotFoundException{})
	mockProductDataSource.EXPECT().Update(gomock.Any()).Return(nil).AnyTimes()
	productGateway := gateways.NewProductGateway(mockProductDataSource, mockFileProvider)
	uc := use_cases.NewUpdateProductUseCase(*productGateway)
	_, err := uc.Execute(productDTO)
	require.Error(t, err)
	_, ok := err.(*exceptions.ProductNotFoundException)
	require.True(t, ok)
}

func TestUpdateProductUseCase_SetNameError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockProductDataSource := mock_interfaces.NewMockIProductDataSource(ctrl)
	mockFileProvider := mock_interfaces.NewMockIFileProvider(ctrl)
	categoryID := "cat-1"
	productDTO := makeProductDTO("pid", categoryID, "", "Descrição", 10.0, true)
	mockProductDataSource.EXPECT().FindByID("pid").Return(daos.ProductDAO{ID: "pid", CategoryID: categoryID, Name: "Produto Teste", Description: "Descrição", Price: 10.0, Active: true}, nil)
	mockProductDataSource.EXPECT().Update(gomock.Any()).Return(nil).AnyTimes()
	productGateway := gateways.NewProductGateway(mockProductDataSource, mockFileProvider)
	uc := use_cases.NewUpdateProductUseCase(*productGateway)
	_, err := uc.Execute(productDTO)
	require.Error(t, err)
}
