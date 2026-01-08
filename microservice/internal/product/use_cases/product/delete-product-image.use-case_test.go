package use_cases

import (
	"testing"

	"tech_challenge/internal/product/application/gateways"
	"tech_challenge/internal/product/daos"
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

// func TestDeleteProductImageUseCase_ProductNotFound(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	mockGateway := mock_interfaces.NewMockIProductGateway(ctrl)
// 	productID := "notfound"
// 	imageFileName := "img1.jpg"
// 	mockGateway.EXPECT().FindByID(productID).Return(errors.New("not found"), errors.New("not found"))
// 	uc := NewDeleteProductImageUseCase(mockGateway)
// 	err := uc.Execute(productID, imageFileName)
// 	_, ok := err.(*exceptions.ProductNotFoundException)
// 	require.True(t, ok)
// }

// func TestDeleteProductImageUseCase_ImagesNotFound(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	mockGateway := mock_interfaces.NewMockIProductGateway(ctrl)
// 	productID := "prod-1"
// 	imageFileName := "img1.jpg"
// 	mockGateway.EXPECT().FindByID(productID).Return(nil, nil)
// 	mockGateway.EXPECT().FindAllImagesProductById(productID).Return(&value_objects.ProductImages{Images: []value_objects.Image{}}, errors.New("not found"))
// 	uc := NewDeleteProductImageUseCase(mockGateway)
// 	err := uc.Execute(productID, imageFileName)
// 	_, ok := err.(*exceptions.ProductImagesNotFoundException)
// 	require.True(t, ok)
// }

// func TestDeleteProductImageUseCase_CannotBeEmpty(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	mockGateway := mock_interfaces.NewMockIProductGateway(ctrl)
// 	productID := "prod-1"
// 	imageFileName := "img1.jpg"
// 	mockGateway.EXPECT().FindByID(productID).Return(nil, nil)
// 	mockGateway.EXPECT().FindAllImagesProductById(productID).Return(&value_objects.ProductImages{Images: []value_objects.Image{{FileName: imageFileName}}}, nil)
// 	uc := NewDeleteProductImageUseCase(mockGateway)
// 	err := uc.Execute(productID, imageFileName)
// 	_, ok := err.(*exceptions.ProductImageCannotBeEmptyException)
// 	require.True(t, ok)
// }

// func TestDeleteProductImageUseCase_DefaultImage(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	mockGateway := mock_interfaces.NewMockIProductGateway(ctrl)
// 	productID := "prod-1"
// 	imageFileName := "img1.jpg"
// 	mockGateway.EXPECT().FindByID(productID).Return(nil, nil)
// 	mockGateway.EXPECT().FindAllImagesProductById(productID).Return(&value_objects.ProductImages{Images: []value_objects.Image{{FileName: imageFileName}, {FileName: "img2.jpg"}}}, nil)
// 	mockGateway.EXPECT().ImageIsDefault(imageFileName).Return(true)
// 	mockGateway.EXPECT().SetLastImageAsDefault(productID, imageFileName).Return(nil)
// 	mockGateway.EXPECT().DeleteProductImage(imageFileName).Return(nil)
// 	mockGateway.EXPECT().DeleteImage(imageFileName).Return(nil)
// 	uc := NewDeleteProductImageUseCase(mockGateway)
// 	require.NoError(t, uc.Execute(productID, imageFileName))
// }

// func TestDeleteProductImageUseCase_DeleteProductImageError(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	mockGateway := mock_interfaces.NewMockIProductGateway(ctrl)
// 	productID := "prod-1"
// 	imageFileName := "img1.jpg"
// 	mockGateway.EXPECT().FindByID(productID).Return(nil, nil)
// 	mockGateway.EXPECT().FindAllImagesProductById(productID).Return(&value_objects.ProductImages{Images: []value_objects.Image{{FileName: imageFileName}, {FileName: "img2.jpg"}}}, nil)
// 	mockGateway.EXPECT().ImageIsDefault(imageFileName).Return(false)
// 	mockGateway.EXPECT().DeleteProductImage(imageFileName).Return(errors.New("db error"))
// 	uc := NewDeleteProductImageUseCase(mockGateway)
// 	err := uc.Execute(productID, imageFileName)
// 	_, ok := err.(*exceptions.InvalidProductImageException)
// 	require.True(t, ok)
// }

// func TestDeleteProductImageUseCase_DeleteImageError(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	mockGateway := mock_interfaces.NewMockIProductGateway(ctrl)
// 	productID := "prod-1"
// 	imageFileName := "img1.jpg"
// 	mockGateway.EXPECT().FindByID(productID).Return(nil, nil)
// 	mockGateway.EXPECT().FindAllImagesProductById(productID).Return(&value_objects.ProductImages{Images: []value_objects.Image{{FileName: imageFileName}, {FileName: "img2.jpg"}}}, nil)
// 	mockGateway.EXPECT().ImageIsDefault(imageFileName).Return(false)
// 	mockGateway.EXPECT().DeleteProductImage(imageFileName).Return(nil)
// 	mockGateway.EXPECT().DeleteImage(imageFileName).Return(fmt.Errorf("bucket error"))
// 	uc := NewDeleteProductImageUseCase(mockGateway)
// 	err := uc.Execute(productID, imageFileName)
// 	require.EqualError(t, err, "failed to delete file from bucket: bucket error")
// }
