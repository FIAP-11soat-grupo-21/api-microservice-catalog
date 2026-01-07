package controllers

import (
	"errors"
	"os"
	"tech_challenge/internal/product/application/dtos"
	"tech_challenge/internal/product/daos"
	mock_interfaces "tech_challenge/internal/product/interfaces/mocks"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	os.Setenv("GO_ENV", "test")
	os.Setenv("API_PORT", "8080")
	os.Setenv("API_HOST", "localhost")
	os.Setenv("API_UPLOAD_URL", "http://localhost:8080/uploads")
	os.Setenv("DB_RUN_MIGRATIONS", "false")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_NAME", "test_db")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USERNAME", "test_user")
	os.Setenv("DB_PASSWORD", "test_pass")
	os.Setenv("AWS_REGION", "us-east-1")
	os.Setenv("AWS_S3_BUCKET_NAME", "test-bucket")
	os.Setenv("AWS_S3_PRESIGN_EXPIRATION", "3600")
	code := m.Run()
	os.Exit(code)
}

// Mocks para ProductDataSource, CategoryDataSource e FileProvider

type mockProductDataSource struct {
	insertFunc                   func(dao daos.ProductDAO) error
	findAllFunc                  func() ([]daos.ProductDAO, error)
	findByIDFunc                 func(id string) (daos.ProductDAO, error)
	updateFunc                   func(dao daos.ProductDAO) error
	deleteFunc                   func(id string) error
	addProductImageFunc          func(img daos.ProductImageDAO) error
	deleteImageFunc              func(imageFileName string) error
	findAllImagesProductByIdFunc func(productID string) ([]daos.ProductImageDAO, error)
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

// Corrige a assinatura do método DeleteImage para bater com a interface
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
		// Garante que sempre retorna pelo menos uma imagem para o produto
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

func TestProductController_Create_Success(t *testing.T) {
	mockProductDs := &mockProductDataSource{
		insertFunc: func(dao daos.ProductDAO) error { return nil },
	}
	mockCategoryDs := &mockCategoryDataSource{
		findByIDFunc: func(id string) (daos.CategoryDAO, error) {
			return daos.CategoryDAO{ID: id, Name: "Bebidas", Active: true}, nil
		},
	}
	mockFileProvider := &mock_interfaces.MockFileProvider{}
	c := NewProductController(mockProductDs, mockCategoryDs, mockFileProvider)
	dto := dtos.CreateProductDTO{Name: "Coca-Cola", CategoryID: "catid", Price: 5.99, Active: true}
	res, err := c.Create(dto)
	require.NoError(t, err)
	require.Equal(t, "Coca-Cola", res.Name)
}

func TestProductController_Create_Error(t *testing.T) {
	mockProductDs := &mockProductDataSource{
		insertFunc: func(dao daos.ProductDAO) error { return errors.New("fail") },
	}
	mockCategoryDs := &mockCategoryDataSource{
		findByIDFunc: func(id string) (daos.CategoryDAO, error) {
			return daos.CategoryDAO{ID: id, Name: "Bebidas", Active: true}, nil
		},
	}
	c := NewProductController(mockProductDs, mockCategoryDs, &mock_interfaces.MockFileProvider{})
	dto := dtos.CreateProductDTO{Name: "Coca-Cola", CategoryID: "catid", Price: 5.99, Active: true}
	_, err := c.Create(dto)
	require.Error(t, err)
}

func TestProductController_FindByID_Success(t *testing.T) {
	mockProductDs := &mockProductDataSource{
		findByIDFunc: func(id string) (daos.ProductDAO, error) {
			return daos.ProductDAO{ID: id, Name: "Coca-Cola", CategoryID: "catid", Price: 5.99, Active: true}, nil
		},
	}
	mockCategoryDs := &mockCategoryDataSource{
		findByIDFunc: func(id string) (daos.CategoryDAO, error) {
			return daos.CategoryDAO{ID: id, Name: "Bebidas", Active: true}, nil
		},
	}
	c := NewProductController(mockProductDs, mockCategoryDs, &mock_interfaces.MockFileProvider{})
	res, err := c.FindByID("pid")
	require.NoError(t, err)
	require.Equal(t, "pid", res.ID)
}

func TestProductController_FindByID_Error(t *testing.T) {
	mockProductDs := &mockProductDataSource{
		findByIDFunc: func(id string) (daos.ProductDAO, error) { return daos.ProductDAO{}, errors.New("fail") },
	}
	mockCategoryDs := &mockCategoryDataSource{
		findByIDFunc: func(id string) (daos.CategoryDAO, error) {
			return daos.CategoryDAO{ID: id, Name: "Bebidas", Active: true}, nil
		},
	}
	c := NewProductController(mockProductDs, mockCategoryDs, &mock_interfaces.MockFileProvider{})
	_, err := c.FindByID("pid")
	require.Error(t, err)
}

func TestProductController_FindAll_Success(t *testing.T) {
	mockProductDs := &mockProductDataSource{
		findAllFunc: func() ([]daos.ProductDAO, error) {
			return []daos.ProductDAO{{ID: "pid", Name: "Coca-Cola", CategoryID: "catid", Price: 5.99, Active: true}}, nil
		},
	}
	mockCategoryDs := &mockCategoryDataSource{
		findByIDFunc: func(id string) (daos.CategoryDAO, error) {
			return daos.CategoryDAO{ID: id, Name: "Bebidas", Active: true}, nil
		},
	}
	c := NewProductController(mockProductDs, mockCategoryDs, &mock_interfaces.MockFileProvider{})
	res, err := c.FindAll(nil)
	require.NoError(t, err)
	require.Len(t, res, 1)
	require.Equal(t, "pid", res[0].ID)
}

func TestProductController_FindAll_Error(t *testing.T) {
	mockProductDs := &mockProductDataSource{
		findAllFunc: func() ([]daos.ProductDAO, error) { return nil, errors.New("fail") },
	}
	mockCategoryDs := &mockCategoryDataSource{
		findByIDFunc: func(id string) (daos.CategoryDAO, error) {
			return daos.CategoryDAO{ID: id, Name: "Bebidas", Active: true}, nil
		},
	}
	c := NewProductController(mockProductDs, mockCategoryDs, &mock_interfaces.MockFileProvider{})
	_, err := c.FindAll(nil)
	require.Error(t, err)
}

func TestProductController_Update_Success(t *testing.T) {
	mockProductDs := &mockProductDataSource{
		updateFunc: func(dao daos.ProductDAO) error { return nil },
		findByIDFunc: func(id string) (daos.ProductDAO, error) {
			return daos.ProductDAO{ID: id, Name: "Coca-Cola", CategoryID: "catid", Price: 5.99, Active: true}, nil
		},
	}
	mockCategoryDs := &mockCategoryDataSource{
		findByIDFunc: func(id string) (daos.CategoryDAO, error) {
			return daos.CategoryDAO{ID: id, Name: "Bebidas", Active: true}, nil
		},
	}
	c := NewProductController(mockProductDs, mockCategoryDs, &mock_interfaces.MockFileProvider{})
	dto := dtos.UpdateProductDTO{ID: "pid", Name: "Coca-Cola", CategoryID: "catid", Price: 5.99, Active: true}
	res, err := c.Update(dto)
	require.NoError(t, err)
	require.Equal(t, "pid", res.ID)
}

func TestProductController_Update_Error(t *testing.T) {
	mockProductDs := &mockProductDataSource{
		updateFunc: func(dao daos.ProductDAO) error { return errors.New("fail") },
		findByIDFunc: func(id string) (daos.ProductDAO, error) {
			return daos.ProductDAO{ID: id, Name: "Coca-Cola", CategoryID: "catid", Price: 5.99, Active: true}, nil
		},
	}
	mockCategoryDs := &mockCategoryDataSource{
		findByIDFunc: func(id string) (daos.CategoryDAO, error) {
			return daos.CategoryDAO{ID: id, Name: "Bebidas", Active: true}, nil
		},
	}
	c := NewProductController(mockProductDs, mockCategoryDs, &mock_interfaces.MockFileProvider{})
	dto := dtos.UpdateProductDTO{ID: "pid", Name: "Coca-Cola", CategoryID: "catid", Price: 5.99, Active: true}
	_, err := c.Update(dto)
	require.Error(t, err)
}

func TestProductController_Delete_Success(t *testing.T) {
	mockProductDs := &mockProductDataSource{
		deleteFunc: func(id string) error { return nil },
		findByIDFunc: func(id string) (daos.ProductDAO, error) {
			return daos.ProductDAO{ID: id, Name: "Coca-Cola", CategoryID: "catid", Price: 5.99, Active: true}, nil
		},
	}
	mockCategoryDs := &mockCategoryDataSource{
		findByIDFunc: func(id string) (daos.CategoryDAO, error) {
			return daos.CategoryDAO{ID: id, Name: "Bebidas", Active: true}, nil
		},
	}
	c := NewProductController(mockProductDs, mockCategoryDs, &mock_interfaces.MockFileProvider{})
	require.NoError(t, c.Delete("pid"))
}

func TestProductController_Delete_Error(t *testing.T) {
	mockProductDs := &mockProductDataSource{
		deleteFunc: func(id string) error { return errors.New("fail") },
		findByIDFunc: func(id string) (daos.ProductDAO, error) {
			return daos.ProductDAO{ID: id, Name: "Coca-Cola", CategoryID: "catid", Price: 5.99, Active: true}, nil
		},
	}
	mockCategoryDs := &mockCategoryDataSource{
		findByIDFunc: func(id string) (daos.CategoryDAO, error) {
			return daos.CategoryDAO{ID: id, Name: "Bebidas", Active: true}, nil
		},
	}
	c := NewProductController(mockProductDs, mockCategoryDs, &mock_interfaces.MockFileProvider{})

	require.Error(t, c.Delete("pid"))
}
func TestProductController_UploadImage_Success(t *testing.T) {
	mockProductDs := &mockProductDataSource{
		addProductImageFunc: func(img daos.ProductImageDAO) error { return nil },
		deleteFunc:          func(id string) error { return nil },
		findByIDFunc: func(id string) (daos.ProductDAO, error) {
			return daos.ProductDAO{ID: id, Name: "Coca-Cola", CategoryID: "catid", Price: 5.99, Active: true}, nil
		},
	}
	mockCategoryDs := &mockCategoryDataSource{
		findByIDFunc: func(id string) (daos.CategoryDAO, error) {
			return daos.CategoryDAO{ID: id, Name: "Bebidas", Active: true}, nil
		},
	}
	c := NewProductController(mockProductDs, mockCategoryDs, &mock_interfaces.MockFileProvider{})
	uploadDTO := dtos.UploadProductImageDTO{ProductID: "pid", FileName: "img.jpg", FileContent: []byte("fake")}
	require.NoError(t, c.UploadImage(uploadDTO))
}
func TestProductController_UploadImage_Error(t *testing.T) {
	mockProductDs := &mockProductDataSource{
		addProductImageFunc: func(img daos.ProductImageDAO) error { return errors.New("upload error") },
		deleteFunc:          func(id string) error { return nil },
		findByIDFunc: func(id string) (daos.ProductDAO, error) {
			return daos.ProductDAO{ID: id, Name: "Coca-Cola", CategoryID: "catid", Price: 5.99, Active: true}, nil
		},
	}
	mockCategoryDs := &mockCategoryDataSource{
		findByIDFunc: func(id string) (daos.CategoryDAO, error) {
			return daos.CategoryDAO{ID: id, Name: "Bebidas", Active: true}, nil
		},
	}
	c := NewProductController(mockProductDs, mockCategoryDs, &mock_interfaces.MockFileProvider{})
	uploadDTO := dtos.UploadProductImageDTO{ProductID: "pid", FileName: "img.jpg", FileContent: []byte("fake")}
	err := c.UploadImage(uploadDTO)
	require.Error(t, err)
}

type mockProductGateway struct {
	findAllImagesProductByIdFunc func(productID string) ([]daos.ProductImageDAO, error)
}

func (m *mockProductGateway) FindAllImagesProductById(productID string) ([]daos.ProductImageDAO, error) {
	if m.findAllImagesProductByIdFunc != nil {
		return m.findAllImagesProductByIdFunc(productID)
	}
	return nil, nil
}

func TestProductController_DeleteImage_Success(t *testing.T) {
	mockProductDs := &mockProductDataSource{
		deleteImageFunc: func(imageFileName string) error { return nil },
		deleteFunc:      func(id string) error { return nil },
		findByIDFunc: func(id string) (daos.ProductDAO, error) {
			return daos.ProductDAO{ID: id, Name: "Coca-Cola", CategoryID: "catid", Price: 5.99, Active: true}, nil
		},
		findAllImagesProductByIdFunc: func(productID string) ([]daos.ProductImageDAO, error) {
			return []daos.ProductImageDAO{
				{ID: "imgid", ProductID: productID, FileName: "img.jpg", IsDefault: true},
				{ID: "imgid2", ProductID: productID, FileName: "img2.jpg", IsDefault: false},
			}, nil
		},
	}
	mockCategoryDs := &mockCategoryDataSource{
		findByIDFunc: func(id string) (daos.CategoryDAO, error) {
			return daos.CategoryDAO{ID: id, Name: "Bebidas", Active: true}, nil
		},
	}
	c := NewProductController(mockProductDs, mockCategoryDs, &mock_interfaces.MockFileProvider{})
	err := c.DeleteImage("pid", "img.jpg")
	require.NoError(t, err)
}

func TestProductController_DeleteImage_Error(t *testing.T) {
	mockProductDs := &mockProductDataSource{
		deleteImageFunc: func(imageFileName string) error { return errors.New("delete error") },
		deleteFunc:      func(id string) error { return nil },
		findByIDFunc: func(id string) (daos.ProductDAO, error) {
			return daos.ProductDAO{ID: id, Name: "Coca-Cola", CategoryID: "catid", Price: 5.99, Active: true}, nil
		},
	}
	mockCategoryDs := &mockCategoryDataSource{
		findByIDFunc: func(id string) (daos.CategoryDAO, error) {
			return daos.CategoryDAO{ID: id, Name: "Bebidas", Active: true}, nil
		},
	}
	c := NewProductController(mockProductDs, mockCategoryDs, &mock_interfaces.MockFileProvider{})
	err := c.DeleteImage("pid", "img.jpg")
	require.Error(t, err)
}

func TestProductController_FindAllImagesProductById_Success(t *testing.T) {
	mockProductDs := &mockProductDataSource{
		deleteImageFunc: func(imageFileName string) error { return nil },
		deleteFunc:      func(id string) error { return nil },
		findByIDFunc: func(id string) (daos.ProductDAO, error) {
			return daos.ProductDAO{ID: id, Name: "Coca-Cola", CategoryID: "catid", Price: 5.99, Active: true}, nil
		},
		findAllImagesProductByIdFunc: func(productID string) ([]daos.ProductImageDAO, error) {
			return []daos.ProductImageDAO{
				{ID: "imgid", ProductID: productID, FileName: "img.jpg", IsDefault: true},
				{ID: "imgid2", ProductID: productID, FileName: "img2.jpg", IsDefault: false},
			}, nil
		},
	}
	mockCategoryDs := &mockCategoryDataSource{
		findByIDFunc: func(id string) (daos.CategoryDAO, error) {
			return daos.CategoryDAO{ID: id, Name: "Bebidas", Active: true}, nil
		},
	}
	c := NewProductController(mockProductDs, mockCategoryDs, &mock_interfaces.MockFileProvider{})
	res, err := c.FindAllImagesProductById("pid")
	require.NoError(t, err)
	require.Len(t, res, 2)
	require.Equal(t, "imgid", res[0].ID)
	require.Equal(t, "img.jpg", res[0].FileName)
	require.True(t, res[0].IsDefault)
	require.Equal(t, "imgid2", res[1].ID)
	require.Equal(t, "img2.jpg", res[1].FileName)
	require.False(t, res[1].IsDefault)
}

func TestProductController_FindAllImagesProductById_Error(t *testing.T) {
	mockProductDs := &mockProductDataSource{
		deleteImageFunc: func(imageFileName string) error { return nil },
		deleteFunc:      func(id string) error { return nil },
		findByIDFunc: func(id string) (daos.ProductDAO, error) {
			return daos.ProductDAO{ID: id, Name: "Coca-Cola", CategoryID: "catid", Price: 5.99, Active: true}, nil
		},
		findAllImagesProductByIdFunc: func(productID string) ([]daos.ProductImageDAO, error) {
			return nil, errors.New("mock error")
		},
	}
	mockCategoryDs := &mockCategoryDataSource{
		findByIDFunc: func(id string) (daos.CategoryDAO, error) {
			return daos.CategoryDAO{ID: id, Name: "Bebidas", Active: true}, nil
		},
	}
	c := NewProductController(mockProductDs, mockCategoryDs, &mock_interfaces.MockFileProvider{})
	res, err := c.FindAllImagesProductById("pid")
	require.Error(t, err)
	require.Nil(t, res)
}
