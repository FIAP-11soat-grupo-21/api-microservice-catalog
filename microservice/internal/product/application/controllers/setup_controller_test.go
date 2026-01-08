package controllers

import (
	"os"
	"tech_challenge/internal/product/application/dtos"
	"tech_challenge/internal/product/daos"
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
	uploadImageFunc              func(dto dtos.UploadProductImageDTO) error
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
