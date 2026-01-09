package gateways

import (
	"errors"
	"os"
	"tech_challenge/internal/product/daos"
	"tech_challenge/internal/product/domain/entities"
	value_objects "tech_challenge/internal/product/domain/value-objects"
	testenv "tech_challenge/internal/shared/test"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	testenv.SetupTestEnv()
	code := m.Run()
	os.Exit(code)
}

type mockProductDataSource struct {
	insertFunc                           func(dao daos.ProductDAO) error
	findAllFunc                          func() ([]daos.ProductDAO, error)
	findByIDFunc                         func(id string) (daos.ProductDAO, error)
	updateFunc                           func(dao daos.ProductDAO) error
	deleteFunc                           func(id string) error
	findAllByCategoryIDFunc              func(categoryID string) ([]daos.ProductDAO, error)
	addProductImageFunc                  func(img daos.ProductImageDAO) error
	setAllPreviousImagesAsNotDefaultFunc func(productID, exceptImageID string) error
	findAllImagesProductByIdFunc         func(productID string) ([]daos.ProductImageDAO, error)
	setImageAsDefaultFunc                func(productID, imageID string) error
	deleteImageFunc                      func(imageFileName string) error
}

func (m *mockProductDataSource) Insert(dao daos.ProductDAO) error {
	return m.insertFunc(dao)
}
func (m *mockProductDataSource) FindAll() ([]daos.ProductDAO, error) {
	return m.findAllFunc()
}
func (m *mockProductDataSource) FindByID(id string) (daos.ProductDAO, error) {
	return m.findByIDFunc(id)
}
func (m *mockProductDataSource) Update(dao daos.ProductDAO) error {
	return m.updateFunc(dao)
}
func (m *mockProductDataSource) Delete(id string) error {
	return m.deleteFunc(id)
}
func (m *mockProductDataSource) FindAllByCategoryID(categoryID string) ([]daos.ProductDAO, error) {
	return m.findAllByCategoryIDFunc(categoryID)
}
func (m *mockProductDataSource) AddProductImage(img daos.ProductImageDAO) error {
	return m.addProductImageFunc(img)
}
func (m *mockProductDataSource) SetAllPreviousImagesAsNotDefault(productID, exceptImageID string) error {
	return m.setAllPreviousImagesAsNotDefaultFunc(productID, exceptImageID)
}
func (m *mockProductDataSource) FindAllImagesProductById(productID string) ([]daos.ProductImageDAO, error) {
	return m.findAllImagesProductByIdFunc(productID)
}
func (m *mockProductDataSource) SetImageAsDefault(productID, imageID string) error {
	return m.setImageAsDefaultFunc(productID, imageID)
}
func (m *mockProductDataSource) DeleteImage(imageFileName string) error {
	return m.deleteImageFunc(imageFileName)
}

type mockFileProvider struct{}

func (m *mockFileProvider) UploadFile(fileName string, fileContent []byte) error { return nil }
func (m *mockFileProvider) DeleteFile(fileName string) error                     { return nil }
func (m *mockFileProvider) GetPresignedURL(fileName string) (string, error) {
	return "http://localhost/" + fileName, nil
}
func (m *mockFileProvider) DeleteFiles(fileNames []string) error { return nil }

func TestProductGateway_Insert(t *testing.T) {
	gw := NewProductGateway(&mockProductDataSource{
		insertFunc: func(dao daos.ProductDAO) error { return nil },
	}, &mockFileProvider{})
	name, _ := value_objects.NewName("Coca-Cola")
	price, _ := value_objects.NewPrice(5.99)
	prod, _ := entities.NewProduct("pid", "catid", name.Value(), "desc", price.Value(), true)
	require.NoError(t, gw.Insert(*prod))
}

func TestProductGateway_FindAll_Success(t *testing.T) {
	gw := NewProductGateway(&mockProductDataSource{
		findAllFunc: func() ([]daos.ProductDAO, error) {
			return []daos.ProductDAO{{ID: "pid", Name: "Coca-Cola", CategoryID: "catid", Price: 5.99, Active: true, Images: []daos.ProductImageDAO{}}}, nil
		},
	}, &mockFileProvider{})
	prods, err := gw.FindAll()
	require.NoError(t, err)
	require.Len(t, prods, 1)
	require.Equal(t, "pid", prods[0].ID)
}

func TestProductGateway_FindAll_Error(t *testing.T) {
	gw := NewProductGateway(&mockProductDataSource{
		findAllFunc: func() ([]daos.ProductDAO, error) { return nil, errors.New("fail") },
	}, &mockFileProvider{})
	prods, err := gw.FindAll()
	require.Error(t, err)
	require.Nil(t, prods)
}

func TestProductGateway_FindAll_ImagesMapping(t *testing.T) {
	createdAt := time.Now()
	gw := NewProductGateway(&mockProductDataSource{
		findAllFunc: func() ([]daos.ProductDAO, error) {
			return []daos.ProductDAO{{
				ID:         "pid",
				Name:       "Coca-Cola",
				CategoryID: "catid",
				Price:      5.99,
				Active:     true,
				Images: []daos.ProductImageDAO{{
					ID:        "imgid",
					ProductID: "pid",
					FileName:  "img.jpg",
					Url:       "http://localhost/img.jpg",
					CreatedAt: createdAt,
					IsDefault: true,
				}},
			}}, nil
		},
	}, &mockFileProvider{})
	prods, err := gw.FindAll()
	require.NoError(t, err)
	require.Len(t, prods, 1)
	require.Len(t, prods[0].Images, 1)
	img := prods[0].Images[0]
	require.Equal(t, "img.jpg", img.FileName)
	require.Equal(t, "http://localhost/img.jpg", img.Url)
	require.Equal(t, createdAt, img.CreatedAt)
	require.Equal(t, "imgid", img.ID)
	require.True(t, img.IsDefault)
}

func TestProductGateway_FindByID_Success(t *testing.T) {
	gw := NewProductGateway(&mockProductDataSource{
		findByIDFunc: func(id string) (daos.ProductDAO, error) {
			return daos.ProductDAO{ID: "pid", Name: "Coca-Cola", CategoryID: "catid", Price: 5.99, Active: true, Images: []daos.ProductImageDAO{}}, nil
		},
	}, &mockFileProvider{})
	prod, err := gw.FindByID("pid")
	require.NoError(t, err)
	require.Equal(t, "pid", prod.ID)
}

func TestProductGateway_FindByID_Error(t *testing.T) {
	gw := NewProductGateway(&mockProductDataSource{
		findByIDFunc: func(id string) (daos.ProductDAO, error) { return daos.ProductDAO{}, errors.New("fail") },
	}, &mockFileProvider{})
	prod, err := gw.FindByID("pid")
	require.Error(t, err)
	require.Equal(t, entities.Product{}, prod)
}

func TestProductGateway_FindByID_WithImages(t *testing.T) {
	createdAt := time.Now()
	gw := NewProductGateway(&mockProductDataSource{
		findByIDFunc: func(id string) (daos.ProductDAO, error) {
			return daos.ProductDAO{
				ID:         "pid",
				Name:       "Coca-Cola",
				CategoryID: "catid",
				Price:      5.99,
				Active:     true,
				Images: []daos.ProductImageDAO{{
					ID:        "imgid",
					ProductID: "pid",
					FileName:  "img.jpg",
					Url:       "http://localhost/img.jpg",
					CreatedAt: createdAt,
					IsDefault: true,
				}},
			}, nil
		},
	}, &mockFileProvider{})
	prod, err := gw.FindByID("pid")
	require.NoError(t, err)
	require.Equal(t, "pid", prod.ID)
	require.Len(t, prod.Images, 1)
	img := prod.Images[0]
	require.Equal(t, "img.jpg", img.FileName)
	require.Equal(t, "http://localhost/img.jpg", img.Url)
	require.Equal(t, createdAt, img.CreatedAt)
	require.Equal(t, "imgid", img.ID)
	require.True(t, img.IsDefault)
}

func TestProductGateway_Update(t *testing.T) {
	gw := NewProductGateway(&mockProductDataSource{
		updateFunc: func(dao daos.ProductDAO) error { return nil },
	}, &mockFileProvider{})
	name, _ := value_objects.NewName("Coca-Cola")
	price, _ := value_objects.NewPrice(5.99)
	prod, _ := entities.NewProduct("pid", "catid", name.Value(), "desc", price.Value(), true)
	require.NoError(t, gw.Update(*prod))
}

func TestProductGateway_Delete(t *testing.T) {
	gw := NewProductGateway(&mockProductDataSource{
		deleteFunc: func(id string) error { return nil },
	}, &mockFileProvider{})
	require.NoError(t, gw.Delete("pid"))
}

func TestProductGateway_UploadImage(t *testing.T) {
	gw := NewProductGateway(&mockProductDataSource{}, &mockFileProvider{})
	url, err := gw.UploadImage("img.jpg", []byte("data"))
	require.NoError(t, err)
	require.Contains(t, url, "img.jpg")
}

func TestProductGateway_UploadImage_Error(t *testing.T) {
	gw := NewProductGateway(&mockProductDataSource{}, &mockFileProviderErrorUpload{})
	url, err := gw.UploadImage("img.jpg", []byte("data"))
	require.Error(t, err)
	require.Equal(t, "", url)
}

// Mock para simular erro no UploadFile

type mockFileProviderErrorUpload struct{}

func (m *mockFileProviderErrorUpload) UploadFile(fileName string, fileContent []byte) error {
	return errors.New("upload fail")
}
func (m *mockFileProviderErrorUpload) DeleteFile(fileName string) error { return nil }
func (m *mockFileProviderErrorUpload) GetPresignedURL(fileName string) (string, error) {
	return "", nil
}
func (m *mockFileProviderErrorUpload) DeleteFiles(fileNames []string) error { return nil }

func TestProductGateway_DeleteImage(t *testing.T) {
	gw := NewProductGateway(&mockProductDataSource{}, &mockFileProvider{})
	require.NoError(t, gw.DeleteImage("img.jpg"))
}

func TestProductGateway_GetImageUrl(t *testing.T) {
	gw := NewProductGateway(&mockProductDataSource{}, &mockFileProvider{})
	url := gw.GetImageUrl("img.jpg")
	require.Contains(t, url, "img.jpg")
}

func TestProductGateway_GetImageUrl_Error(t *testing.T) {
	gw := NewProductGateway(&mockProductDataSource{}, &mockFileProviderError{})
	url := gw.GetImageUrl("img.jpg")
	require.Equal(t, "", url)
}

// Mock para simular erro no GetPresignedURL

type mockFileProviderError struct{}

func (m *mockFileProviderError) UploadFile(fileName string, fileContent []byte) error { return nil }
func (m *mockFileProviderError) DeleteFile(fileName string) error                     { return nil }
func (m *mockFileProviderError) GetPresignedURL(fileName string) (string, error) {
	return "", errors.New("fail")
}
func (m *mockFileProviderError) DeleteFiles(fileNames []string) error { return nil }

func TestProductGateway_AddProductImage(t *testing.T) {
	gw := NewProductGateway(&mockProductDataSource{
		addProductImageFunc: func(img daos.ProductImageDAO) error { return nil },
	}, &mockFileProvider{})
	img := daos.ProductImageDAO{ID: "imgid", ProductID: "pid", FileName: "img.jpg"}
	require.NoError(t, gw.AddProductImage(img))
}

func TestProductGateway_AddAndSetDefaultImage(t *testing.T) {
	gw := NewProductGateway(&mockProductDataSource{
		addProductImageFunc:                  func(img daos.ProductImageDAO) error { return nil },
		setAllPreviousImagesAsNotDefaultFunc: func(productID, exceptImageID string) error { return nil },
	}, &mockFileProvider{})
	name, _ := value_objects.NewName("Coca-Cola")
	price, _ := value_objects.NewPrice(5.99)
	prod, _ := entities.NewProduct("pid", "catid", name.Value(), "desc", price.Value(), true)
	img := &value_objects.Image{ID: "imgid", FileName: "img.jpg"}
	prod.Images = append(prod.Images, img)
	require.NoError(t, gw.AddAndSetDefaultImage(*prod, "url"))
}

func TestProductGateway_AddAndSetDefaultImage_NoImages(t *testing.T) {
	gw := NewProductGateway(&mockProductDataSource{}, &mockFileProvider{})
	name, _ := value_objects.NewName("Coca-Cola")
	price, _ := value_objects.NewPrice(5.99)
	prod, _ := entities.NewProduct("pid", "catid", name.Value(), "desc", price.Value(), true)
	prod.Images = nil // sem imagens
	err := gw.AddAndSetDefaultImage(*prod, "url")
	require.Error(t, err)
	require.Contains(t, err.Error(), "Produto não possui imagens para atualizar")
}

func TestProductGateway_AddAndSetDefaultImage_AddProductImageError(t *testing.T) {
	gw := NewProductGateway(&mockProductDataSource{
		addProductImageFunc:                  func(img daos.ProductImageDAO) error { return errors.New("add image error") },
		setAllPreviousImagesAsNotDefaultFunc: func(productID, exceptImageID string) error { return nil },
	}, &mockFileProvider{})
	name, _ := value_objects.NewName("Coca-Cola")
	price, _ := value_objects.NewPrice(5.99)
	prod, _ := entities.NewProduct("pid", "catid", name.Value(), "desc", price.Value(), true)
	img := &value_objects.Image{ID: "imgid", FileName: "img.jpg"}
	prod.Images = append(prod.Images, img)
	err := gw.AddAndSetDefaultImage(*prod, "url")
	require.Error(t, err)
	require.Contains(t, err.Error(), "add image error")
}

func TestProductGateway_FindAllImagesProductById(t *testing.T) {
	gw := NewProductGateway(&mockProductDataSource{
		findAllImagesProductByIdFunc: func(productID string) ([]daos.ProductImageDAO, error) {
			return []daos.ProductImageDAO{{ID: "imgid", ProductID: productID, FileName: "img.jpg"}}, nil
		},
	}, &mockFileProvider{})
	prod, err := gw.FindAllImagesProductById("pid")
	require.NoError(t, err)
	require.Equal(t, "pid", prod.ID)
	require.Len(t, prod.Images, 1)
}

func TestProductGateway_FindAllImagesProductById_Error(t *testing.T) {
	gw := NewProductGateway(&mockProductDataSource{
		findAllImagesProductByIdFunc: func(productID string) ([]daos.ProductImageDAO, error) {
			return nil, errors.New("find images error")
		},
	}, &mockFileProvider{})
	prod, err := gw.FindAllImagesProductById("pid")
	require.Error(t, err)
	require.Equal(t, entities.Product{}, prod)
}

func TestProductGateway_SetLastImageAsDefault(t *testing.T) {
	gw := NewProductGateway(&mockProductDataSource{
		findAllImagesProductByIdFunc: func(productID string) ([]daos.ProductImageDAO, error) {
			return []daos.ProductImageDAO{{ID: "imgid", ProductID: productID, FileName: "img.jpg", CreatedAt: time.Now()}}, nil
		},
		setImageAsDefaultFunc: func(productID, imageID string) error { return nil },
	}, &mockFileProvider{})
	require.NoError(t, gw.SetLastImageAsDefault("pid", "except.jpg"))
}

func TestProductGateway_SetLastImageAsDefault_ContinueSkipExceptImage(t *testing.T) {
	createdAt1 := time.Now().Add(-time.Hour)
	createdAt2 := time.Now()
	gw := NewProductGateway(&mockProductDataSource{
		findAllImagesProductByIdFunc: func(productID string) ([]daos.ProductImageDAO, error) {
			return []daos.ProductImageDAO{
				{ID: "imgid1", ProductID: productID, FileName: "except.jpg", CreatedAt: createdAt1},
				{ID: "imgid2", ProductID: productID, FileName: "img2.jpg", CreatedAt: createdAt2},
			}, nil
		},
		setImageAsDefaultFunc: func(productID, imageID string) error {
			require.Equal(t, "imgid2", imageID)
			return nil
		},
	}, &mockFileProvider{})
	err := gw.SetLastImageAsDefault("pid", "except.jpg")
	require.NoError(t, err)
}

func TestProductGateway_SetLastImageAsDefault_FindAllImagesError(t *testing.T) {
	gw := NewProductGateway(&mockProductDataSource{
		findAllImagesProductByIdFunc: func(productID string) ([]daos.ProductImageDAO, error) {
			return nil, errors.New("find images error")
		},
	}, &mockFileProvider{})
	err := gw.SetLastImageAsDefault("pid", "except.jpg")
	require.Error(t, err)
	require.EqualError(t, err, "find images error")
}

func TestProductGateway_SetLastImageAsDefault_LastImageNil(t *testing.T) {
	gw := NewProductGateway(&mockProductDataSource{
		findAllImagesProductByIdFunc: func(productID string) ([]daos.ProductImageDAO, error) {
			// Só retorna a imagem que será ignorada pelo continue
			return []daos.ProductImageDAO{
				{ID: "imgid1", ProductID: productID, FileName: "except.jpg", CreatedAt: time.Now()},
			}, nil
		},
	}, &mockFileProvider{})
	err := gw.SetLastImageAsDefault("pid", "except.jpg")
	require.NoError(t, err)
}

func TestProductGateway_SetLastImageAsDefault_NoImagesLeft(t *testing.T) {
	gw := NewProductGateway(&mockProductDataSource{
		findAllImagesProductByIdFunc: func(productID string) ([]daos.ProductImageDAO, error) {
			// Todas as imagens serão ignoradas pelo continue
			return []daos.ProductImageDAO{
				{ID: "imgid1", ProductID: productID, FileName: "except.jpg", CreatedAt: time.Now()},
			}, nil
		},
	}, &mockFileProvider{})
	err := gw.SetLastImageAsDefault("pid", "except.jpg")
	require.NoError(t, err)
}

func TestProductGateway_SetLastImageAsDefault_ReturnNilWhenNoImages(t *testing.T) {
	gw := NewProductGateway(&mockProductDataSource{
		findAllImagesProductByIdFunc: func(productID string) ([]daos.ProductImageDAO, error) {
			// Retorna slice vazio para simular nenhum registro
			return []daos.ProductImageDAO{}, nil
		},
	}, &mockFileProvider{})
	err := gw.SetLastImageAsDefault("pid", "except.jpg")
	require.NoError(t, err)
}

func TestProductGateway_DeleteProductImage(t *testing.T) {
	gw := NewProductGateway(&mockProductDataSource{
		deleteImageFunc: func(imageFileName string) error { return nil },
	}, &mockFileProvider{})
	require.NoError(t, gw.DeleteProductImage("img.jpg"))
}

func TestProductGateway_DeleteFiles(t *testing.T) {
	gw := NewProductGateway(&mockProductDataSource{}, &mockFileProvider{})
	img := &value_objects.Image{FileName: "img.jpg"}
	img2 := &value_objects.Image{FileName: value_objects.DEFAULT_IMAGE_FILE_NAME}
	require.NoError(t, gw.DeleteFiles([]*value_objects.Image{img, img2}))
}

func TestProductGateway_FindAllByCategoryID_Success(t *testing.T) {
	createdAt := time.Now()
	gw := NewProductGateway(&mockProductDataSource{
		findAllByCategoryIDFunc: func(categoryID string) ([]daos.ProductDAO, error) {
			return []daos.ProductDAO{{
				ID:         "pid",
				Name:       "Coca-Cola",
				CategoryID: categoryID,
				Price:      5.99,
				Active:     true,
				Images: []daos.ProductImageDAO{{
					ID:        "imgid",
					ProductID: "pid",
					FileName:  "img.jpg",
					Url:       "http://localhost/img.jpg",
					CreatedAt: createdAt,
					IsDefault: true,
				}},
			}}, nil
		},
	}, &mockFileProvider{})
	prods, err := gw.FindAllByCategoryID("catid")
	require.NoError(t, err)
	require.Len(t, prods, 1)
	require.Equal(t, "pid", prods[0].ID)
	require.Equal(t, "catid", prods[0].CategoryID)
	require.Len(t, prods[0].Images, 1)
	img := prods[0].Images[0]
	require.Equal(t, "img.jpg", img.FileName)
	require.Equal(t, "http://localhost/img.jpg", img.Url)
	require.Equal(t, createdAt, img.CreatedAt)
	require.Equal(t, "imgid", img.ID)
	require.True(t, img.IsDefault)
}

func TestProductGateway_FindAllByCategoryID_Error(t *testing.T) {
	gw := NewProductGateway(&mockProductDataSource{
		findAllByCategoryIDFunc: func(categoryID string) ([]daos.ProductDAO, error) {
			return nil, errors.New("fail")
		},
	}, &mockFileProvider{})
	prods, err := gw.FindAllByCategoryID("catid")
	require.Error(t, err)
	require.Nil(t, prods)
}

func TestProductGateway_FindAllByCategoryID_Error_Entity(t *testing.T) {
	gw := NewProductGateway(&mockProductDataSource{
		findAllByCategoryIDFunc: func(categoryID string) ([]daos.ProductDAO, error) {
			// Retorna um ProductDAO inválido para forçar erro na conversão para entidade
			return []daos.ProductDAO{{ID: "", Name: "", CategoryID: categoryID, Price: 5.99, Active: true}}, nil
		},
	}, &mockFileProvider{})
	prods, err := gw.FindAllByCategoryID("catid")
	require.Error(t, err)
	require.Nil(t, prods)
}

func TestProductGateway_FindAll_Error_Entity(t *testing.T) {
	gw := NewProductGateway(&mockProductDataSource{
		findAllFunc: func() ([]daos.ProductDAO, error) {
			// Retorna um ProductDAO inválido para forçar erro na conversão para entidade
			return []daos.ProductDAO{{ID: "", Name: "", CategoryID: "catid", Price: 5.99, Active: true}}, nil
		},
	}, &mockFileProvider{})
	prods, err := gw.FindAll()
	require.Error(t, err)
	require.Nil(t, prods)
}

func TestProductGateway_FindByID_Error_Entity(t *testing.T) {
	gw := NewProductGateway(&mockProductDataSource{
		findByIDFunc: func(id string) (daos.ProductDAO, error) {
			// Retorna um ProductDAO inválido para forçar erro na conversão para entidade
			return daos.ProductDAO{ID: "", Name: "", CategoryID: "catid", Price: 5.99, Active: true}, nil
		},
	}, &mockFileProvider{})
	prod, err := gw.FindByID("pid")
	require.Error(t, err)
	require.Equal(t, entities.Product{}, prod)
}
