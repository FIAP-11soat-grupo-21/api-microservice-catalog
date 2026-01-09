package testenv

import (
	"tech_challenge/internal/product/application/dtos"
	"tech_challenge/internal/product/daos"
)

type MockProductDataSource struct {
	FindAllFunc                          func() ([]daos.ProductDAO, error)
	FindByIDFunc                         func(string) (daos.ProductDAO, error)
	FindAllImagesProductByIdFunc         func(string) ([]daos.ProductImageDAO, error)
	FindAllByCategoryIDFunc              func(string) ([]daos.ProductDAO, error)
	InsertFunc                           func(daos.ProductDAO) error
	UpdateFunc                           func(daos.ProductDAO) error
	DeleteFunc                           func(string) error
	DeleteImageFunc                      func(string) error
	AddProductImageFunc                  func(daos.ProductImageDAO) error
	SetAllPreviousImagesAsNotDefaultFunc func(productID, exceptImageID string) error
	SetImageAsDefaultFunc                func(productID, imageID string) error
	UploadImageFunc                      func(uploadDTO dtos.UploadProductImageDTO) error
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
	// Retorna produto com imagem para testes de sucesso
	return daos.ProductDAO{
		ID:          id,
		Name:        "prod",
		Description: "desc",
		Price:       1.0,
		Active:      true,
		CategoryID:  "catid",
		Images:      []daos.ProductImageDAO{{ID: "imgid", ProductID: id, FileName: "img.jpg", IsDefault: true}},
	}, nil
}
func (m *MockProductDataSource) FindAllImagesProductById(id string) ([]daos.ProductImageDAO, error) {
	if m.FindAllImagesProductByIdFunc != nil {
		return m.FindAllImagesProductByIdFunc(id)
	}
	return nil, nil
}
func (m *MockProductDataSource) FindAllByCategoryID(categoryID string) ([]daos.ProductDAO, error) {
	if m.FindAllByCategoryIDFunc != nil {
		return m.FindAllByCategoryIDFunc(categoryID)
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
func (m *MockProductDataSource) SetAllPreviousImagesAsNotDefault(productID, exceptImageID string) error {
	if m.SetAllPreviousImagesAsNotDefaultFunc != nil {
		return m.SetAllPreviousImagesAsNotDefaultFunc(productID, exceptImageID)
	}
	return nil
}
func (m *MockProductDataSource) SetImageAsDefault(productID, imageID string) error {
	if m.SetImageAsDefaultFunc != nil {
		return m.SetImageAsDefaultFunc(productID, imageID)
	}
	return nil
}
func (m *MockProductDataSource) UploadImage(uploadDTO dtos.UploadProductImageDTO) error {
	if m.UploadImageFunc != nil {
		return m.UploadImageFunc(uploadDTO)
	}
	return nil
}

type MockCategoryDataSource struct {
	FindByIDFunc func(string) (daos.CategoryDAO, error)
	DeleteFunc   func(string) error
	InsertFunc   func(daos.CategoryDAO) error
	FindAllFunc  func() ([]daos.CategoryDAO, error)
	UpdateFunc   func(daos.CategoryDAO) error
}

func (m *MockCategoryDataSource) FindByID(id string) (daos.CategoryDAO, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}
	return daos.CategoryDAO{}, nil
}
func (m *MockCategoryDataSource) Delete(id string) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(id)
	}
	return nil
}
func (m *MockCategoryDataSource) Insert(dao daos.CategoryDAO) error {
	if m.InsertFunc != nil {
		return m.InsertFunc(dao)
	}
	return nil
}
func (m *MockCategoryDataSource) FindAll() ([]daos.CategoryDAO, error) {
	if m.FindAllFunc != nil {
		return m.FindAllFunc()
	}
	return nil, nil
}
func (m *MockCategoryDataSource) Update(dao daos.CategoryDAO) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(dao)
	}
	return nil
}
