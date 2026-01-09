package test

import (
	"tech_challenge/internal/product/daos"
)

type MockProductDataSource struct {
	FindAllFunc                  func() ([]daos.ProductDAO, error)
	FindByIDFunc                 func(string) (daos.ProductDAO, error)
	FindAllImagesProductByIdFunc func(string) ([]daos.ProductImageDAO, error)
	InsertFunc                   func(daos.ProductDAO) error
	UpdateFunc                   func(daos.ProductDAO) error
	DeleteFunc                   func(string) error
	DeleteImageFunc              func(string) error
	AddProductImageFunc          func(daos.ProductImageDAO) error
}

func (m *MockProductDataSource) FindAll() ([]daos.ProductDAO, error) {
	if m.FindAllFunc != nil {
		return m.FindAllFunc()
	}
	return nil, nil
}
func (m *MockProductDataSource) FindByID(id string) (daos.ProductDAO, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}
	return daos.ProductDAO{}, nil
}
func (m *MockProductDataSource) FindAllImagesProductById(id string) ([]daos.ProductImageDAO, error) {
	if m.FindAllImagesProductByIdFunc != nil {
		return m.FindAllImagesProductByIdFunc(id)
	}
	return nil, nil
}
func (m *MockProductDataSource) Insert(p daos.ProductDAO) error {
	if m.InsertFunc != nil {
		return m.InsertFunc(p)
	}
	return nil
}
func (m *MockProductDataSource) Update(p daos.ProductDAO) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(p)
	}
	return nil
}
func (m *MockProductDataSource) Delete(id string) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(id)
	}
	return nil
}
func (m *MockProductDataSource) DeleteImage(fileName string) error {
	if m.DeleteImageFunc != nil {
		return m.DeleteImageFunc(fileName)
	}
	return nil
}
func (m *MockProductDataSource) AddProductImage(img daos.ProductImageDAO) error {
	if m.AddProductImageFunc != nil {
		return m.AddProductImageFunc(img)
	}
	return nil
}

type MockCategoryDataSource struct {
	FindByIDFunc func(string) (daos.CategoryDAO, error)
}

func (m *MockCategoryDataSource) FindByID(id string) (daos.CategoryDAO, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}
	return daos.CategoryDAO{}, nil
}
