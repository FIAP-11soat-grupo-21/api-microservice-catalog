package use_cases_test

import (
	"testing"

	"tech_challenge/internal/product/application/gateways"
	mock_interfaces "tech_challenge/internal/product/interfaces/mocks"
	category "tech_challenge/internal/product/use_cases/category"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreateCategoryUseCase_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockCategoryDataSource := mock_interfaces.NewMockICategoryDataSource(ctrl)
	mockCategoryDataSource.EXPECT().Insert(gomock.Any()).Return(nil)

	categoryGateway := gateways.NewCategoryGateway(mockCategoryDataSource)
	uc := category.NewCreateCategoryUseCase(categoryGateway)
	cat, err := uc.Execute("Bebidas", true)
	require.NoError(t, err)
	require.Equal(t, "Bebidas", cat.Name.Value())
	require.True(t, cat.Active)
}
