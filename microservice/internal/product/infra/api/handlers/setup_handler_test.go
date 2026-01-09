package handlers

import (
	"os"
	"tech_challenge/internal/product/application/controllers"
	mock_interfaces "tech_challenge/internal/product/interfaces/mocks"
	testmocks "tech_challenge/internal/shared/test"

	"testing"
)

func TestMain(m *testing.M) {
	testmocks.SetupTestEnv()
	code := m.Run()
	os.Exit(code)
}

func setupProductHandlerWithFakeGateway(productDs *testmocks.MockProductDataSource, categoryDs *testmocks.MockCategoryDataSource, fileProvider *mock_interfaces.MockIFileProvider) *ProductHandler {
	ctrl := controllers.NewProductController(productDs, categoryDs, fileProvider)
	return &ProductHandler{productController: *ctrl}
}
func setupCategoryHandlerWithFakeGateway(categoryDs *testmocks.MockCategoryDataSource) *CategoryHandler {
	ctrl := controllers.NewCategoryController(categoryDs)
	return &CategoryHandler{categoryController: *ctrl}
}
