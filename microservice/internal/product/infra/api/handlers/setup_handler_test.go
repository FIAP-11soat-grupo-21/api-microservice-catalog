package handlers

import (
	"os"
	"tech_challenge/internal/product/application/controllers"
	"tech_challenge/internal/product/daos"
	mock_interfaces "tech_challenge/internal/product/interfaces/mocks"
	testenv "tech_challenge/internal/shared/test"
	"testing"
)

func TestMain(m *testing.M) {
	testenv.SetupTestEnv()
	code := m.Run()
	os.Exit(code)
}

type mockProductDataSource struct {
	insertFunc                   func(dao daos.ProductDAO) error
	findAllFunc                  func() ([]daos.ProductDAO, error)
	findByIDFunc                 func(id string) (daos.ProductDAO, error)
	updateFunc                   func(dao daos.ProductDAO) error
	deleteFunc                   func(id string) error
	addProductImageFunc          func(img daos.ProductImageDAO) error
	deleteImageFunc              func(imageFileName string) error
	findAllImagesProductByIdFunc func(productID string) ([]daos.ProductImageDAO, error)
	// uploadImageFunc              func(dto dtos.UploadProductImageDTO) error
}

func (m *mockProductDataSource) Insert(dao daos.ProductDAO) error {
	if m.insertFunc != nil {
		return m.insertFunc(dao)
	}
	return nil
}
func (m *mockProductDataSource) FindAll() ([]daos.ProductDAO, error) {
	if m.findAllFunc != nil {
		return m.findAllFunc()
	}
	return nil, nil
}
func (m *mockProductDataSource) FindByID(id string) (daos.ProductDAO, error) {
	if m.findByIDFunc != nil {
		return m.findByIDFunc(id)
	}
	return daos.ProductDAO{}, nil
}
func (m *mockProductDataSource) Update(dao daos.ProductDAO) error {
	if m.updateFunc != nil {
		return m.updateFunc(dao)
	}
	return nil
}
func (m *mockProductDataSource) Delete(id string) error {
	if m.deleteFunc != nil {
		return m.deleteFunc(id)
	}
	return nil
}
func (m *mockProductDataSource) AddProductImage(img daos.ProductImageDAO) error {
	if m.addProductImageFunc != nil {
		return m.addProductImageFunc(img)
	}
	return nil
}
func (m *mockProductDataSource) DeleteImage(imageFileName string) error {
	if m.deleteImageFunc != nil {
		return m.deleteImageFunc(imageFileName)
	}
	return nil
}
func (m *mockProductDataSource) SetAllPreviousImagesAsNotDefault(productID, exceptImageID string) error {
	return nil
}
func (m *mockProductDataSource) FindAllImagesProductById(productID string) ([]daos.ProductImageDAO, error) {
	if m.findAllImagesProductByIdFunc != nil {
		imgs, err := m.findAllImagesProductByIdFunc(productID)
		if err != nil {
			return nil, err
		}
		if len(imgs) == 0 {
			return []daos.ProductImageDAO{{ID: "imgid", ProductID: productID, FileName: "img.jpg", IsDefault: true}}, nil
		}
		return imgs, nil
	}
	return []daos.ProductImageDAO{{ID: "imgid", ProductID: productID, FileName: "img.jpg", IsDefault: true}}, nil
}
func (m *mockProductDataSource) SetImageAsDefault(productID, imageID string) error {
	return nil
}
func (m *mockProductDataSource) FindAllByCategoryID(categoryID string) ([]daos.ProductDAO, error) {
	return nil, nil
}

// Mock de CategoryDataSource

type mockCategoryDataSource struct {
	insertFunc   func(dao daos.CategoryDAO) error
	findAllFunc  func() ([]daos.CategoryDAO, error)
	findByIDFunc func(id string) (daos.CategoryDAO, error)
	updateFunc   func(dao daos.CategoryDAO) error
	deleteFunc   func(id string) error
}

func (m *mockCategoryDataSource) Insert(dao daos.CategoryDAO) error {
	return m.insertFunc(dao)
}
func (m *mockCategoryDataSource) FindAll() ([]daos.CategoryDAO, error) {
	return m.findAllFunc()
}
func (m *mockCategoryDataSource) FindByID(id string) (daos.CategoryDAO, error) {
	return m.findByIDFunc(id)
}
func (m *mockCategoryDataSource) Update(dao daos.CategoryDAO) error {
	return m.updateFunc(dao)
}
func (m *mockCategoryDataSource) Delete(id string) error {
	return m.deleteFunc(id)
}

// Helper para abrir imagem de teste
// func openTestImage(t *testing.T) *os.File {
// 	imgPaths := []string{
// 		"microservice/uploads/default_product_image.jpg",
// 		"./microservice/uploads/default_product_image.jpg",
// 		"uploads/default_product_image.jpg",
// 		"./uploads/default_product_image.jpg",
// 		"../uploads/default_product_image.jpg",
// 		"../microservice/uploads/default_product_image.jpg",
// 		"C:/Users/thali/fiap/api-microservice-catalog/microservice/uploads/default_product_image.jpg",
// 	}
// 	var f *os.File
// 	var err error
// 	for _, path := range imgPaths {
// 		f, err = os.Open(path)
// 		if err == nil {
// 			return f
// 		}
// 	}
// 	return createFakeJPEG(t)
// }

// func createFakeJPEG(t *testing.T) *os.File {
// 	tempFile, err := os.CreateTemp("", "default_product_image_*.jpg")
// 	if err != nil {
// 		t.Fatalf("Não foi possível criar arquivo temporário de imagem: %v", err)
// 	}
// 	_, err = tempFile.Write([]byte{0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46, 0x49, 0x46, 0x00, 0x01})
// 	if err != nil {
// 		t.Fatalf("Não foi possível escrever no arquivo temporário de imagem: %v", err)
// 	}
// 	if _, err := tempFile.Seek(0, 0); err != nil {
// 		t.Fatalf("Não foi possível reposicionar o ponteiro do arquivo: %v", err)
// 	}
// 	return tempFile
// }

func setupProductHandlerWithFakeGateway(productDs *mockProductDataSource, categoryDs *mockCategoryDataSource, fileProvider *mock_interfaces.MockIFileProvider) *ProductHandler {
	ctrl := controllers.NewProductController(productDs, categoryDs, fileProvider)
	return &ProductHandler{productController: *ctrl}
}
func setupCategoryHandlerWithFakeGateway(categoryDs *mockCategoryDataSource) *CategoryHandler {
	ctrl := controllers.NewCategoryController(categoryDs)
	return &CategoryHandler{categoryController: *ctrl}
}
