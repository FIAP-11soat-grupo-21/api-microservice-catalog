package controllers

import (
	"errors"
	"testing"
	"time"

	"tech_challenge/internal/product/application/dtos"
	"tech_challenge/internal/product/daos"
	mock_interfaces "tech_challenge/internal/product/interfaces/mocks"
	testmocks "tech_challenge/internal/shared/test"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func setupProductControllerTest(t *testing.T) (*testmocks.MockCategoryDataSource, *testmocks.MockProductDataSource, *mock_interfaces.MockIFileProvider, *gomock.Controller) {
	ctrl := gomock.NewController(t)
	mockCategoryDs := &testmocks.MockCategoryDataSource{
		FindByIDFunc: func(id string) (daos.CategoryDAO, error) {
			return daos.CategoryDAO{ID: id, Name: "Bebidas", Active: true}, nil
		},
	}
	mockProductDs := &testmocks.MockProductDataSource{}
	mockFileProvider := mock_interfaces.NewMockIFileProvider(ctrl)
	return mockCategoryDs, mockProductDs, mockFileProvider, ctrl
}

func TestProductController_Create_Error(t *testing.T) {
	mockCategoryDs, mockProductDs, mockFileProvider, ctrl := setupProductControllerTest(t)
	defer ctrl.Finish()
	mockProductDs.InsertFunc = func(dao daos.ProductDAO) error { return nil }
	c := NewProductController(mockProductDs, mockCategoryDs, mockFileProvider)
	productDTO := dtos.CreateProductDTO{
		CategoryID:  "cat1",
		Name:        "Produto Teste",
		Description: "Descrição",
		Price:       10.0,
		Active:      true,
	}
	res, err := c.Create(productDTO)
	require.NoError(t, err)
	require.Equal(t, "cat1", res.CategoryID)
	require.Equal(t, "Produto Teste", res.Name)
	require.Equal(t, "Descrição", res.Description)
	require.Equal(t, 10.0, res.Price)
	require.True(t, res.Active)
}

func TestProductController_Create_ReturnsError(t *testing.T) {
	mockCategoryDs, mockProductDs, mockFileProvider, ctrl := setupProductControllerTest(t)
	defer ctrl.Finish()
	mockProductDs.InsertFunc = func(dao daos.ProductDAO) error { return errors.New("insert error") }
	c := NewProductController(mockProductDs, mockCategoryDs, mockFileProvider)
	productDTO := dtos.CreateProductDTO{
		CategoryID:  "cat1",
		Name:        "Produto Teste",
		Description: "Descrição",
		Price:       10.0,
		Active:      true,
	}
	res, err := c.Create(productDTO)
	require.Error(t, err)
	require.Equal(t, dtos.ProductResultDTO{}, res)
}

func TestProductController_FindByID_Success(t *testing.T) {
	mockCategoryDs, mockProductDs, mockFileProvider, ctrl := setupProductControllerTest(t)
	defer ctrl.Finish()
	mockProductDs.FindByIDFunc = func(id string) (daos.ProductDAO, error) {
		return daos.ProductDAO{ID: id, Name: "Produto Teste", Description: "desc", Price: 10.0, CategoryID: "cat1", Active: true}, nil
	}
	c := NewProductController(mockProductDs, mockCategoryDs, mockFileProvider)
	res, err := c.FindByID("pid")
	require.NoError(t, err)
	require.Equal(t, "pid", res.ID)
	require.Equal(t, "Produto Teste", res.Name)
}

func TestProductController_FindByID_Error(t *testing.T) {
	mockCategoryDs, mockProductDs, mockFileProvider, ctrl := setupProductControllerTest(t)
	defer ctrl.Finish()
	mockProductDs.FindByIDFunc = func(id string) (daos.ProductDAO, error) {
		return daos.ProductDAO{}, errors.New("not found")
	}
	c := NewProductController(mockProductDs, mockCategoryDs, mockFileProvider)
	res, err := c.FindByID("pid")
	require.Error(t, err)
	require.Equal(t, dtos.ProductResultDTO{}, res)
}

func TestProductController_FindAll_Success(t *testing.T) {
	mockCategoryDs, mockProductDs, mockFileProvider, ctrl := setupProductControllerTest(t)
	defer ctrl.Finish()
	mockProductDs.FindAllFunc = func() ([]daos.ProductDAO, error) {
		return []daos.ProductDAO{
			{ID: "pid", Name: "Produto Teste", Description: "desc", Price: 10.0, CategoryID: "cat1", Active: true},
		}, nil
	}
	c := NewProductController(mockProductDs, mockCategoryDs, mockFileProvider)
	res, err := c.FindAll(nil)
	require.NoError(t, err)
	require.Len(t, res, 1)
	require.Equal(t, "pid", res[0].ID)
}

func TestProductController_FindAll_Error(t *testing.T) {
	mockCategoryDs, mockProductDs, mockFileProvider, ctrl := setupProductControllerTest(t)
	defer ctrl.Finish()
	mockProductDs.FindAllFunc = func() ([]daos.ProductDAO, error) {
		return nil, errors.New("find all error")
	}
	c := NewProductController(mockProductDs, mockCategoryDs, mockFileProvider)
	res, err := c.FindAll(nil)
	require.Error(t, err)
	require.Nil(t, res)
}

func TestProductController_Update_Success(t *testing.T) {
	mockCategoryDs, mockProductDs, mockFileProvider, ctrl := setupProductControllerTest(t)
	defer ctrl.Finish()
	mockProductDs.UpdateFunc = func(dao daos.ProductDAO) error { return nil }
	mockProductDs.FindByIDFunc = func(id string) (daos.ProductDAO, error) {
		return daos.ProductDAO{ID: id, Name: "Produto Atualizado", Description: "desc", Price: 20.0, CategoryID: "cat1", Active: true}, nil
	}
	c := NewProductController(mockProductDs, mockCategoryDs, mockFileProvider)
	updateDTO := dtos.UpdateProductDTO{
		ID:          "pid",
		CategoryID:  "cat1",
		Name:        "Produto Atualizado",
		Description: "desc",
		Price:       20.0,
		Active:      true,
	}
	res, err := c.Update(updateDTO)
	require.NoError(t, err)
	require.Equal(t, "pid", res.ID)
	require.Equal(t, "Produto Atualizado", res.Name)
}

func TestProductController_Update_Error(t *testing.T) {
	mockCategoryDs, mockProductDs, mockFileProvider, ctrl := setupProductControllerTest(t)
	defer ctrl.Finish()
	mockProductDs.UpdateFunc = func(dao daos.ProductDAO) error { return errors.New("update error") }
	c := NewProductController(mockProductDs, mockCategoryDs, mockFileProvider)
	updateDTO := dtos.UpdateProductDTO{
		ID:          "pid",
		CategoryID:  "cat1",
		Name:        "Produto Atualizado",
		Description: "desc",
		Price:       20.0,
		Active:      true,
	}
	res, err := c.Update(updateDTO)
	require.Error(t, err)
	require.Equal(t, dtos.ProductResultDTO{}, res)
}

func TestProductController_UploadImage_Success(t *testing.T) {
	mockCategoryDs, mockProductDs, mockFileProvider, ctrl := setupProductControllerTest(t)
	defer ctrl.Finish()
	mockProductDs.FindByIDFunc = func(id string) (daos.ProductDAO, error) {
		return daos.ProductDAO{ID: id, Name: "Produto Teste", Description: "desc", Price: 10.0, CategoryID: "cat1", Active: true}, nil
	}
	mockProductDs.UploadImageFunc = func(uploadDTO dtos.UploadProductImageDTO) error { return nil }
	mockFileProvider.EXPECT().UploadFile(gomock.Any(), gomock.Any()).Return(nil)
	mockFileProvider.EXPECT().GetPresignedURL(gomock.Any()).Return("http://localhost:8080/uploads/test-bucket/img.jpg", nil).AnyTimes()
	c := NewProductController(mockProductDs, mockCategoryDs, mockFileProvider)
	uploadDTO := dtos.UploadProductImageDTO{
		ProductID:   "pid",
		FileName:    "img.jpg",
		FileContent: []byte("filedata"),
	}
	err := c.UploadImage(uploadDTO)
	require.NoError(t, err)
}

func TestProductController_UploadImage_Error(t *testing.T) {
	mockCategoryDs, mockProductDs, mockFileProvider, ctrl := setupProductControllerTest(t)
	defer ctrl.Finish()
	mockProductDs.FindByIDFunc = func(id string) (daos.ProductDAO, error) {
		return daos.ProductDAO{ID: id, Name: "Produto Teste", Description: "desc", Price: 10.0, CategoryID: "cat1", Active: true}, nil
	}
	mockProductDs.UploadImageFunc = func(uploadDTO dtos.UploadProductImageDTO) error { return errors.New("upload error") }
	mockFileProvider.EXPECT().UploadFile(gomock.Any(), gomock.Any()).Return(errors.New("upload error"))
	c := NewProductController(mockProductDs, mockCategoryDs, mockFileProvider)
	uploadDTO := dtos.UploadProductImageDTO{
		ProductID:   "pid",
		FileName:    "img.jpg",
		FileContent: []byte("filedata"),
	}
	err := c.UploadImage(uploadDTO)
	require.Error(t, err)
}

func TestProductController_DeleteImage_Success(t *testing.T) {
	mockCategoryDs, mockProductDs, mockFileProvider, ctrl := setupProductControllerTest(t)
	defer ctrl.Finish()
	mockProductDs.FindByIDFunc = func(id string) (daos.ProductDAO, error) {
		return daos.ProductDAO{ID: id, Name: "Produto Teste", Description: "desc", Price: 10.0, CategoryID: "cat1", Active: true}, nil
	}
	mockProductDs.FindAllImagesProductByIdFunc = func(productID string) ([]daos.ProductImageDAO, error) {
		return []daos.ProductImageDAO{
			{ID: "imgid1", ProductID: productID, FileName: "img.jpg", CreatedAt: time.Now()},
			{ID: "imgid2", ProductID: productID, FileName: "img2.jpg", CreatedAt: time.Now()},
		}, nil
	}
	mockProductDs.DeleteImageFunc = func(imageFileName string) error { return nil }
	mockFileProvider.EXPECT().DeleteFile(gomock.Any()).Return(nil).AnyTimes()
	c := NewProductController(mockProductDs, mockCategoryDs, mockFileProvider)
	err := c.DeleteImage("pid", "img.jpg")
	require.NoError(t, err)
}

func TestProductController_DeleteImage_Error(t *testing.T) {
	mockCategoryDs, mockProductDs, mockFileProvider, ctrl := setupProductControllerTest(t)
	defer ctrl.Finish()
	mockProductDs.DeleteImageFunc = func(imageFileName string) error { return errors.New("delete image error") }
	mockFileProvider.EXPECT().DeleteFiles(gomock.Any()).Return(nil).AnyTimes()
	c := NewProductController(mockProductDs, mockCategoryDs, mockFileProvider)
	err := c.DeleteImage("pid", "img.jpg")
	require.Error(t, err)
}

func TestProductController_Delete_Success(t *testing.T) {
	mockCategoryDs, mockProductDs, mockFileProvider, ctrl := setupProductControllerTest(t)
	defer ctrl.Finish()
	mockProductDs.FindByIDFunc = func(id string) (daos.ProductDAO, error) {
		return daos.ProductDAO{ID: id, Name: "Produto Teste", Description: "desc", Price: 10.0, CategoryID: "cat1", Active: true}, nil
	}
	mockProductDs.DeleteFunc = func(id string) error { return nil }
	mockFileProvider.EXPECT().DeleteFiles(gomock.Any()).Return(nil).AnyTimes()
	mockFileProvider.EXPECT().DeleteFile(gomock.Any()).Return(nil).AnyTimes()
	c := NewProductController(mockProductDs, mockCategoryDs, mockFileProvider)
	err := c.Delete("pid")
	require.NoError(t, err)
}

func TestProductController_Delete_Error(t *testing.T) {
	mockCategoryDs, mockProductDs, mockFileProvider, ctrl := setupProductControllerTest(t)
	defer ctrl.Finish()
	mockProductDs.DeleteFunc = func(id string) error { return errors.New("delete error") }
	mockFileProvider.EXPECT().DeleteFiles(gomock.Any()).Return(nil).AnyTimes()
	mockFileProvider.EXPECT().DeleteFile(gomock.Any()).Return(nil).AnyTimes()
	c := NewProductController(mockProductDs, mockCategoryDs, mockFileProvider)
	err := c.Delete("pid")
	require.Error(t, err)
}

func TestProductController_FindAllImagesProductById_Success(t *testing.T) {
	mockCategoryDs, mockProductDs, mockFileProvider, ctrl := setupProductControllerTest(t)
	defer ctrl.Finish()
	mockProductDs.FindAllImagesProductByIdFunc = func(productID string) ([]daos.ProductImageDAO, error) {
		return []daos.ProductImageDAO{
			{ID: "imgid1", ProductID: productID, FileName: "img.jpg", CreatedAt: time.Now()},
			{ID: "imgid2", ProductID: productID, FileName: "img2.jpg", CreatedAt: time.Now()},
		}, nil
	}
	c := NewProductController(mockProductDs, mockCategoryDs, mockFileProvider)
	res, err := c.FindAllImagesProductById("pid")
	require.NoError(t, err)
	require.Len(t, res, 2)
	require.Equal(t, "img.jpg", res[0].FileName)
}

func TestProductController_FindAllImagesProductById_Error(t *testing.T) {
	mockCategoryDs, mockProductDs, mockFileProvider, ctrl := setupProductControllerTest(t)
	defer ctrl.Finish()
	mockProductDs.FindAllImagesProductByIdFunc = func(productID string) ([]daos.ProductImageDAO, error) {
		return nil, errors.New("find images error")
	}
	c := NewProductController(mockProductDs, mockCategoryDs, mockFileProvider)
	res, err := c.FindAllImagesProductById("pid")
	require.Error(t, err)
	require.Nil(t, res)
}
