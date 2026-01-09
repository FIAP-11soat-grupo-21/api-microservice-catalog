package use_cases_test

import (
	"errors"
	"testing"

	"tech_challenge/internal/product/application/dtos"
	"tech_challenge/internal/product/application/gateways"
	"tech_challenge/internal/product/daos"
	mock_interfaces "tech_challenge/internal/product/interfaces/mocks"
	use_cases "tech_challenge/internal/product/use_cases/product"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func makeUploadProductImageDTO() dtos.UploadProductImageDTO {
	return dtos.UploadProductImageDTO{
		ProductID:   "pid",
		FileName:    "img.jpg",
		FileContent: []byte("filedata"),
	}
}

func TestUploadProductImageUseCase_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockProductDataSource := mock_interfaces.NewMockIProductDataSource(ctrl)
	mockFileProvider := mock_interfaces.NewMockIFileProvider(ctrl)
	mockProductDataSource.EXPECT().FindByID("pid").Return(daos.ProductDAO{ID: "pid", Name: "Produto Teste", Description: "desc", Price: 10.0, CategoryID: "cat1", Active: true}, nil)
	mockFileProvider.EXPECT().UploadFile(gomock.Any(), gomock.Any()).Return(nil)
	mockFileProvider.EXPECT().GetPresignedURL(gomock.Any()).Return("http://localhost:8080/uploads/img.jpg", nil)
	mockProductDataSource.EXPECT().SetAllPreviousImagesAsNotDefault("pid", gomock.Any()).Return(nil).AnyTimes()
	mockProductDataSource.EXPECT().AddProductImage(gomock.Any()).Return(nil)
	mockProductDataSource.EXPECT().SetImageAsDefault("pid", gomock.Any()).Return(nil).AnyTimes()
	productGateway := gateways.NewProductGateway(mockProductDataSource, mockFileProvider)
	uc := use_cases.NewUploadProductImageUseCase(*productGateway)
	productDTO := makeUploadProductImageDTO()
	err := uc.Execute(productDTO)
	require.NoError(t, err)
}

func TestUploadProductImageUseCase_ProductNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockProductDataSource := mock_interfaces.NewMockIProductDataSource(ctrl)
	mockFileProvider := mock_interfaces.NewMockIFileProvider(ctrl)
	mockProductDataSource.EXPECT().FindByID("pid").Return(daos.ProductDAO{}, errors.New("not found"))
	productGateway := gateways.NewProductGateway(mockProductDataSource, mockFileProvider)
	uc := use_cases.NewUploadProductImageUseCase(*productGateway)
	productDTO := makeUploadProductImageDTO()
	err := uc.Execute(productDTO)
	require.Error(t, err)
}

func TestUploadProductImageUseCase_InvalidImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockProductDataSource := mock_interfaces.NewMockIProductDataSource(ctrl)
	mockFileProvider := mock_interfaces.NewMockIFileProvider(ctrl)
	mockProductDataSource.EXPECT().FindByID("pid").Return(daos.ProductDAO{ID: "pid", Name: "Produto Teste", Description: "desc", Price: 10.0, CategoryID: "cat1", Active: true}, nil)
	productGateway := gateways.NewProductGateway(mockProductDataSource, mockFileProvider)
	uc := use_cases.NewUploadProductImageUseCase(*productGateway)
	productDTO := dtos.UploadProductImageDTO{ProductID: "pid", FileName: "", FileContent: []byte("filedata")}
	err := uc.Execute(productDTO)
	require.Error(t, err)
}

func TestUploadProductImageUseCase_UploadFileError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockProductDataSource := mock_interfaces.NewMockIProductDataSource(ctrl)
	mockFileProvider := mock_interfaces.NewMockIFileProvider(ctrl)
	mockProductDataSource.EXPECT().FindByID("pid").Return(daos.ProductDAO{ID: "pid", Name: "Produto Teste", Description: "desc", Price: 10.0, CategoryID: "cat1", Active: true}, nil)
	mockFileProvider.EXPECT().UploadFile(gomock.Any(), gomock.Any()).Return(errors.New("upload error"))
	productGateway := gateways.NewProductGateway(mockProductDataSource, mockFileProvider)
	uc := use_cases.NewUploadProductImageUseCase(*productGateway)
	productDTO := makeUploadProductImageDTO()
	err := uc.Execute(productDTO)
	require.EqualError(t, err, "upload error")
}

func TestUploadProductImageUseCase_GetPresignedURLReturnNil(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockProductDataSource := mock_interfaces.NewMockIProductDataSource(ctrl)
	mockFileProvider := mock_interfaces.NewMockIFileProvider(ctrl)
	mockProductDataSource.EXPECT().FindByID("pid").Return(daos.ProductDAO{ID: "pid", Name: "Produto Teste", Description: "desc", Price: 10.0, CategoryID: "cat1", Active: true}, nil)
	mockFileProvider.EXPECT().UploadFile(gomock.Any(), gomock.Any()).Return(nil)
	mockFileProvider.EXPECT().GetPresignedURL(gomock.Any()).DoAndReturn(
		func(_ interface{}) (string, error) {
			return "", nil
		},
	)
	mockProductDataSource.EXPECT().AddProductImage(gomock.Any()).AnyTimes()
	mockProductDataSource.EXPECT().SetAllPreviousImagesAsNotDefault(gomock.Any(), gomock.Any()).AnyTimes()
	mockProductDataSource.EXPECT().SetImageAsDefault(gomock.Any(), gomock.Any()).AnyTimes()
	productGateway := gateways.NewProductGateway(mockProductDataSource, mockFileProvider)
	uc := use_cases.NewUploadProductImageUseCase(*productGateway)
	productDTO := makeUploadProductImageDTO()
	err := uc.Execute(productDTO)
	require.Nil(t, err)
}

func TestUploadProductImageUseCase_AddProductImageError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockProductDataSource := mock_interfaces.NewMockIProductDataSource(ctrl)
	mockFileProvider := mock_interfaces.NewMockIFileProvider(ctrl)
	mockProductDataSource.EXPECT().FindByID("pid").Return(daos.ProductDAO{ID: "pid", Name: "Produto Teste", Description: "desc", Price: 10.0, CategoryID: "cat1", Active: true}, nil)
	mockFileProvider.EXPECT().UploadFile(gomock.Any(), gomock.Any()).Return(nil)
	mockFileProvider.EXPECT().GetPresignedURL(gomock.Any()).Return("http://localhost:8080/uploads/img.jpg", nil)
	mockProductDataSource.EXPECT().SetAllPreviousImagesAsNotDefault("pid", gomock.Any()).Return(nil).AnyTimes()
	mockProductDataSource.EXPECT().AddProductImage(gomock.Any()).Return(errors.New("add error"))
	mockProductDataSource.EXPECT().SetImageAsDefault("pid", gomock.Any()).Return(nil).AnyTimes()
	productGateway := gateways.NewProductGateway(mockProductDataSource, mockFileProvider)
	uc := use_cases.NewUploadProductImageUseCase(*productGateway)
	productDTO := makeUploadProductImageDTO()
	err := uc.Execute(productDTO)
	require.Error(t, err)
	require.Contains(t, err.Error(), "Invalid product data")
}
