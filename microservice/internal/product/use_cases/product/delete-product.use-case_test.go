package use_cases

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"tech_challenge/internal/product/application/gateways"
	"tech_challenge/internal/product/daos"
	"tech_challenge/internal/product/domain/exceptions"
	mock_interfaces "tech_challenge/internal/product/interfaces/mocks"
)

func TestDeleteProductUseCase_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockProductDataSource := mock_interfaces.NewMockIProductDataSource(ctrl)
	mockFileProvider := mock_interfaces.NewMockIFileProvider(ctrl)
	id := "a3bb189e-8bf9-3888-9912-ace4e6543002"
	gomock.InOrder(
		mockProductDataSource.EXPECT().FindByID(id).Return(daos.ProductDAO{ID: id, Name: "Product 1", Description: "Description 1", Price: 100}, nil),
		mockProductDataSource.EXPECT().FindAllImagesProductById(id).Return([]daos.ProductImageDAO{{FileName: "img1.jpg"}}, nil),
		mockFileProvider.EXPECT().DeleteFiles([]string{"img1.jpg"}).Return(nil),
		mockProductDataSource.EXPECT().Delete(id).Return(nil),
	)
	gw := gateways.NewProductGateway(mockProductDataSource, mockFileProvider)
	uc := NewDeleteProductUseCase(*gw)
	require.NoError(t, uc.Execute(id))
}

func TestDeleteProductUseCase_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockProductDataSource := mock_interfaces.NewMockIProductDataSource(ctrl)
	mockFileProvider := mock_interfaces.NewMockIFileProvider(ctrl)
	id := "not-found-id"
	mockProductDataSource.EXPECT().FindByID(id).Return(daos.ProductDAO{}, errors.New("not found"))

	gw := gateways.NewProductGateway(mockProductDataSource, mockFileProvider)
	uc := NewDeleteProductUseCase(*gw)
	err := uc.Execute(id)
	_, ok := err.(*exceptions.ProductNotFoundException)
	require.True(t, ok)
}

func TestDeleteProductUseCase_ImagesNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockProductDataSource := mock_interfaces.NewMockIProductDataSource(ctrl)
	mockFileProvider := mock_interfaces.NewMockIFileProvider(ctrl)
	id := "img-not-found-id"
	mockProductDataSource.EXPECT().FindByID(id).Return(daos.ProductDAO{ID: id, Name: "Product 1", Description: "Description 1", Price: 100}, nil)
	mockProductDataSource.EXPECT().FindAllImagesProductById(gomock.Any()).DoAndReturn(
		func(productID string) ([]daos.ProductImageDAO, error) {
			t.Logf("FindAllImagesProductById called with: %s", productID)
			return nil, &exceptions.ProductImagesNotFoundException{}
		},
	).Times(1)
	gw := gateways.NewProductGateway(mockProductDataSource, mockFileProvider)
	uc := NewDeleteProductUseCase(*gw)
	err := uc.Execute(id)
	t.Logf("Error returned: %v", err)
	_, ok := err.(*exceptions.ProductImagesNotFoundException)
	require.True(t, ok)
}

func TestDeleteProductUseCase_DeleteFilesError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockProductDataSource := mock_interfaces.NewMockIProductDataSource(ctrl)
	mockFileProvider := mock_interfaces.NewMockIFileProvider(ctrl)
	id := "delete-files-error-id"
	mockProductDataSource.EXPECT().FindByID(id).Return(daos.ProductDAO{ID: id, Name: "Product 1", Description: "Description 1", Price: 100}, nil)
	mockProductDataSource.EXPECT().FindAllImagesProductById(id).Return([]daos.ProductImageDAO{{FileName: "img1.jpg"}}, nil)
	mockFileProvider.EXPECT().DeleteFiles(gomock.Any()).Return(errors.New("delete files error"))
	gw := gateways.NewProductGateway(mockProductDataSource, mockFileProvider)
	uc := NewDeleteProductUseCase(*gw)
	err := uc.Execute(id)
	_, ok := err.(*exceptions.DeleteImagesStorageException)
	require.True(t, ok)
}

func TestDeleteProductUseCase_DeleteError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockProductDataSource := mock_interfaces.NewMockIProductDataSource(ctrl)
	mockFileProvider := mock_interfaces.NewMockIFileProvider(ctrl)
	id := "delete-error-id"
	gomock.InOrder(
		mockProductDataSource.EXPECT().FindByID(id).Return(daos.ProductDAO{ID: id, Name: "Product 1", Description: "Description 1", Price: 100}, nil),
		mockProductDataSource.EXPECT().FindAllImagesProductById(id).Return([]daos.ProductImageDAO{{FileName: "img1.jpg"}}, nil),
		mockFileProvider.EXPECT().DeleteFiles([]string{"img1.jpg"}).Return(nil),
		mockProductDataSource.EXPECT().Delete(id).Return(errors.New("delete error")),
	)
	gw := gateways.NewProductGateway(mockProductDataSource, mockFileProvider)
	uc := NewDeleteProductUseCase(*gw)
	err := uc.Execute(id)
	require.EqualError(t, err, "delete error")
}
