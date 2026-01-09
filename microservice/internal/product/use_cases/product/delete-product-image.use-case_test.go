package use_cases

import (
	"fmt"
	"testing"

	"tech_challenge/internal/product/application/gateways"
	"tech_challenge/internal/product/daos"
	"tech_challenge/internal/product/domain/exceptions"
	mock_interfaces "tech_challenge/internal/product/interfaces/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestDeleteProductImageUseCase_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockProductDataSource := mock_interfaces.NewMockIProductDataSource(ctrl)
	mockFileProvider := mock_interfaces.NewMockIFileProvider(ctrl)
	productID := "a3bb189e-8bf9-3888-9912-ace4e6543002"
	imageFileName := "img1.jpg"

	// Assegura a ordem correta das chamadas esperadas
	mockProductDataSource.EXPECT().FindByID(productID).Return(daos.ProductDAO{ID: productID, Name: "Produto Teste", Description: "desc", Price: 10.0}, nil)
	mockProductDataSource.EXPECT().FindAllImagesProductById(productID).Return(
		[]daos.ProductImageDAO{
			{FileName: imageFileName},
			{FileName: "img2.jpg"},
		}, nil)
	mockProductDataSource.EXPECT().ImageIsDefault(imageFileName).Return(false).AnyTimes()
	mockProductDataSource.EXPECT().DeleteProductImage(imageFileName).Return(nil).AnyTimes()
	mockProductDataSource.EXPECT().DeleteImage(imageFileName).Return(nil)
	mockFileProvider.EXPECT().DeleteFile(imageFileName).Return(nil)

	gw := gateways.NewProductGateway(mockProductDataSource, mockFileProvider)
	uc := NewDeleteProductImageUseCase(*gw)
	err := uc.Execute(productID, imageFileName)
	require.NoError(t, err)
}

func TestDeleteProductImageUseCase_ProductNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockProductDataSource := mock_interfaces.NewMockIProductDataSource(ctrl)
	mockFileProvider := mock_interfaces.NewMockIFileProvider(ctrl)
	productID := "notfound"
	imageFileName := "img1.jpg"

	mockProductDataSource.EXPECT().FindByID(productID).Return(daos.ProductDAO{}, fmt.Errorf("not found"))

	gw := gateways.NewProductGateway(mockProductDataSource, mockFileProvider)
	uc := NewDeleteProductImageUseCase(*gw)
	err := uc.Execute(productID, imageFileName)
	_, ok := err.(*exceptions.ProductNotFoundException)
	require.True(t, ok)
}

func TestDeleteProductImageUseCase_ImagesNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockProductDataSource := mock_interfaces.NewMockIProductDataSource(ctrl)
	mockFileProvider := mock_interfaces.NewMockIFileProvider(ctrl)
	productID := "prod-1"
	imageFileName := "img1.jpg"

	mockProductDataSource.EXPECT().FindByID(productID).Return(daos.ProductDAO{ID: productID, Name: "Produto Teste", Description: "desc", Price: 10.0}, nil)
	// Simula erro ao buscar imagens (err != nil)
	mockProductDataSource.EXPECT().FindAllImagesProductById(productID).Return(nil, fmt.Errorf("not found"))

	gw := gateways.NewProductGateway(mockProductDataSource, mockFileProvider)
	uc := NewDeleteProductImageUseCase(*gw)
	err := uc.Execute(productID, imageFileName)
	_, ok := err.(*exceptions.ProductImagesNotFoundException)
	require.True(t, ok)
}

func TestDeleteProductImageUseCase_CannotBeEmpty(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockProductDataSource := mock_interfaces.NewMockIProductDataSource(ctrl)
	mockFileProvider := mock_interfaces.NewMockIFileProvider(ctrl)
	productID := "prod-1"
	imageFileName := "img1.jpg"

	mockProductDataSource.EXPECT().FindByID(productID).Return(daos.ProductDAO{ID: productID, Name: "Produto Teste", Description: "desc", Price: 10.0}, nil)
	// Simula retorno de apenas uma imagem
	mockProductDataSource.EXPECT().FindAllImagesProductById(productID).Return(
		[]daos.ProductImageDAO{
			{FileName: imageFileName},
		}, nil)

	gw := gateways.NewProductGateway(mockProductDataSource, mockFileProvider)
	uc := NewDeleteProductImageUseCase(*gw)
	err := uc.Execute(productID, imageFileName)
	_, ok := err.(*exceptions.ProductImageCannotBeEmptyException)
	require.True(t, ok)
}
