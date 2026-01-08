package use_cases

import (
	"errors"
	"testing"

	"tech_challenge/internal/product/application/gateways"
	"tech_challenge/internal/product/daos"
	"tech_challenge/internal/product/domain/entities"
	"tech_challenge/internal/product/domain/exceptions"
	mock_interfaces "tech_challenge/internal/product/interfaces/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestFindProductByIDUseCase_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockProductDataSource := mock_interfaces.NewMockIProductDataSource(ctrl)
	mockFileProvider := mock_interfaces.NewMockIFileProvider(ctrl)
	id := "pid"
	mockProductDataSource.EXPECT().FindByID(id).Return(
		daos.ProductDAO{
			ID:         id,
			Name:       "Coca-Cola",
			CategoryID: "cat-1",
			Price:      5.99,
			Active:     true,
			Images:     []daos.ProductImageDAO{},
		},
		nil,
	)
	productGateway := gateways.NewProductGateway(mockProductDataSource, mockFileProvider)
	uc := NewFindProductByIDUseCase(*productGateway)
	product, err := uc.Execute(id)
	require.NoError(t, err)
	require.Equal(t, id, product.ID)
}

func TestFindProductByIDUseCase_Error_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockProductDataSource := mock_interfaces.NewMockIProductDataSource(ctrl)
	mockFileProvider := mock_interfaces.NewMockIFileProvider(ctrl)
	id := "notfound"
	mockProductDataSource.EXPECT().FindByID(id).Return(daos.ProductDAO{}, errors.New("not found"))
	productGateway := gateways.NewProductGateway(mockProductDataSource, mockFileProvider)
	uc := NewFindProductByIDUseCase(*productGateway)
	product, err := uc.Execute(id)
	_, ok := err.(*exceptions.ProductNotFoundException)
	require.True(t, ok)
	require.Equal(t, entities.Product{}, product)
}

func TestFindProductByIDUseCase_Error_EmptyProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockProductDataSource := mock_interfaces.NewMockIProductDataSource(ctrl)
	mockFileProvider := mock_interfaces.NewMockIFileProvider(ctrl)
	id := "empty"
	mockProductDataSource.EXPECT().FindByID(id).Return(daos.ProductDAO{}, nil)
	productGateway := gateways.NewProductGateway(mockProductDataSource, mockFileProvider)
	uc := NewFindProductByIDUseCase(*productGateway)
	product, err := uc.Execute(id)
	_, ok := err.(*exceptions.ProductNotFoundException)
	require.True(t, ok)
	require.Equal(t, entities.Product{}, product)
}
