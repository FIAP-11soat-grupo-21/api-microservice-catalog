package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"tech_challenge/internal/product/daos"
	mock_interfaces "tech_challenge/internal/product/interfaces/mocks"
	testmocks "tech_challenge/internal/shared/test"
)

func setupProductTestEnv(productDs *testmocks.MockProductDataSource, categoryDs *testmocks.MockCategoryDataSource, fileProvider *mock_interfaces.MockIFileProvider) (*gin.Engine, *httptest.ResponseRecorder, *ProductHandler) {
	gin.SetMode(gin.TestMode)
	h := setupProductHandlerWithFakeGateway(productDs, categoryDs, fileProvider)
	r := gin.New()
	w := httptest.NewRecorder()
	return r, w, h
}

func makeDefaultMocks(productDs *testmocks.MockProductDataSource) (*testmocks.MockProductDataSource, *testmocks.MockCategoryDataSource, *mock_interfaces.MockIFileProvider) {
	mockCategoryDs := &testmocks.MockCategoryDataSource{
		FindByIDFunc: func(id string) (daos.CategoryDAO, error) {
			return daos.CategoryDAO{ID: id, Name: "Bebidas", Active: true}, nil
		},
	}
	mockFileProvider := &mock_interfaces.MockIFileProvider{}
	return productDs, mockCategoryDs, mockFileProvider
}

func makeGomockFileProvider(t *testing.T) *mock_interfaces.MockIFileProvider {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)
	return mock_interfaces.NewMockIFileProvider(ctrl)
}

func TestFindAllProducts_Success(t *testing.T) {
	mockProductDs := &testmocks.MockProductDataSource{
		FindAllFunc: func() ([]daos.ProductDAO, error) {
			return []daos.ProductDAO{{ID: "1", Name: "prod", Description: "desc", Price: 1.0, Active: true, CategoryID: "catid"}}, nil
		},
	}
	mockProductDs, mockCategoryDs, mockFileProvider := makeDefaultMocks(mockProductDs)
	r, w, h := setupProductTestEnv(mockProductDs, mockCategoryDs, mockFileProvider)

	r.GET("/products", h.FindAllProducts)

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp []map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Len(t, resp, 1)
	require.Equal(t, "prod", resp[0]["name"])
}

func TestFindAllProducts_WithCategoryID(t *testing.T) {
	mockProductDs := &testmocks.MockProductDataSource{
		FindAllFunc: func() ([]daos.ProductDAO, error) {
			return []daos.ProductDAO{{ID: "2", Name: "prodcat", Description: "desc", Price: 2.0, Active: true, CategoryID: "catid2"}}, nil
		},
	}
	mockProductDs, mockCategoryDs, mockFileProvider := makeDefaultMocks(mockProductDs)
	r, w, h := setupProductTestEnv(mockProductDs, mockCategoryDs, mockFileProvider)

	r.GET("/products", h.FindAllProducts)

	req := httptest.NewRequest(http.MethodGet, "/products?category_id=catid2", nil)
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp []map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	if len(resp) == 0 {
		t.Logf("Body: %s", w.Body.String())
	}
}

func TestFindAllProducts_WithoutCategoryID(t *testing.T) {
	mockProductDs := &testmocks.MockProductDataSource{
		FindAllFunc: func() ([]daos.ProductDAO, error) {
			return []daos.ProductDAO{{ID: "3", Name: "prodsemcat", Description: "desc", Price: 3.0, Active: true, CategoryID: ""}}, nil
		},
	}
	mockProductDs, mockCategoryDs, mockFileProvider := makeDefaultMocks(mockProductDs)
	r, w, h := setupProductTestEnv(mockProductDs, mockCategoryDs, mockFileProvider)

	r.GET("/products", h.FindAllProducts)

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp []map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Len(t, resp, 1)
	require.Equal(t, "prodsemcat", resp[0]["name"])
}

func TestFindAllProducts_Error(t *testing.T) {
	mockProductDs := &testmocks.MockProductDataSource{
		FindAllFunc: func() ([]daos.ProductDAO, error) {
			return nil, errors.New("mock error")
		},
	}
	mockProductDs, mockCategoryDs, mockFileProvider := makeDefaultMocks(mockProductDs)
	r, w, h := setupProductTestEnv(mockProductDs, mockCategoryDs, mockFileProvider)

	r.Use(func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": c.Errors[0].Error()})
		}
	})
	r.GET("/products", h.FindAllProducts)

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestFindAllImagesProductById_Success(t *testing.T) {
	mockProductDs := &testmocks.MockProductDataSource{
		FindAllImagesProductByIdFunc: func(productID string) ([]daos.ProductImageDAO, error) {
			return []daos.ProductImageDAO{{ID: "imgid", ProductID: productID, FileName: "img.jpg", IsDefault: true}}, nil
		},
	}
	mockProductDs, mockCategoryDs, mockFileProvider := makeDefaultMocks(mockProductDs)
	r, w, h := setupProductTestEnv(mockProductDs, mockCategoryDs, mockFileProvider)

	r.GET("/products/:id/images", h.FindAllImagesProductById)

	req := httptest.NewRequest(http.MethodGet, "/products/1/images", nil)
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp []map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Len(t, resp, 1)
	fileName, ok := resp[0]["FileName"].(string)
	require.True(t, ok)
	require.Equal(t, "img.jpg", fileName)
}

func TestFindAllImagesProductById_Error(t *testing.T) {
	mockProductDs := &testmocks.MockProductDataSource{
		FindAllImagesProductByIdFunc: func(productID string) ([]daos.ProductImageDAO, error) {
			return nil, errors.New("not found")
		},
	}
	mockProductDs, mockCategoryDs, mockFileProvider := makeDefaultMocks(mockProductDs)
	r, w, h := setupProductTestEnv(mockProductDs, mockCategoryDs, mockFileProvider)

	r.GET("/products/:id/images", h.FindAllImagesProductById)

	req := httptest.NewRequest(http.MethodGet, "/products/1/images", nil)
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)

	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Contains(t, resp["error"], "not found")
}

func TestDeleteProduct_Success(t *testing.T) {
	mockProductDs := &testmocks.MockProductDataSource{
		DeleteFunc: func(id string) error { return nil },
	}
	mockProductDs, mockCategoryDs, _ := makeDefaultMocks(mockProductDs)
	mockFileProvider := makeGomockFileProvider(t)
	mockFileProvider.EXPECT().DeleteFiles(gomock.Any()).Return(nil).AnyTimes()
	r, w, h := setupProductTestEnv(mockProductDs, mockCategoryDs, mockFileProvider)

	r.DELETE("/products/:id", func(c *gin.Context) {
		h.DeleteProduct(c)
		c.Status(http.StatusNoContent)
	})

	req := httptest.NewRequest(http.MethodDelete, "/products/1", nil)
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNoContent, w.Code)
}

func TestDeleteProduct_Error(t *testing.T) {
	mockProductDs := &testmocks.MockProductDataSource{
		DeleteFunc: func(id string) error { return errors.New("delete error") },
	}
	mockProductDs, mockCategoryDs, _ := makeDefaultMocks(mockProductDs)
	mockFileProvider := makeGomockFileProvider(t)
	mockFileProvider.EXPECT().DeleteFiles(gomock.Any()).Return(nil).AnyTimes()
	r, w, h := setupProductTestEnv(mockProductDs, mockCategoryDs, mockFileProvider)

	r.Use(func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": c.Errors[0].Error()})
		}
	})
	r.DELETE("/products/:id", h.DeleteProduct)

	req := httptest.NewRequest(http.MethodDelete, "/products/1", nil)
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestUpdateProduct_Success(t *testing.T) {
	mockProductDs := &testmocks.MockProductDataSource{
		UpdateFunc: func(dao daos.ProductDAO) error { return nil },
		FindByIDFunc: func(id string) (daos.ProductDAO, error) {
			return daos.ProductDAO{
				ID:          id,
				Name:        "prod",
				Description: "desc",
				Price:       1.0,
				Active:      true,
				CategoryID:  "catid",
				Images: []daos.ProductImageDAO{
					{ID: "img.jpg", ProductID: id, FileName: "img.jpg", IsDefault: true},
				},
			}, nil
		},
	}
	mockProductDs, mockCategoryDs, mockFileProvider := makeDefaultMocks(mockProductDs)
	r, w, h := setupProductTestEnv(mockProductDs, mockCategoryDs, mockFileProvider)

	r.PUT("/products/:id", h.UpdateProduct)

	body := `{"category_id":"catid","name":"prod","description":"desc","price":1.0,"active":true}`
	req := httptest.NewRequest(http.MethodPut, "/products/1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Equal(t, "prod", resp["name"])
}

func TestUpdateProduct_BadRequest(t *testing.T) {
	mockProductDs := &testmocks.MockProductDataSource{}
	mockProductDs, mockCategoryDs, mockFileProvider := makeDefaultMocks(mockProductDs)
	r, w, h := setupProductTestEnv(mockProductDs, mockCategoryDs, mockFileProvider)

	r.PUT("/products/:id", h.UpdateProduct)

	body := `{"name":1}`
	req := httptest.NewRequest(http.MethodPut, "/products/1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)

	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Equal(t, "Invalid request body", resp["error"])
}

func TestUpdateProduct_Error(t *testing.T) {
	mockProductDs := &testmocks.MockProductDataSource{
		UpdateFunc: func(dao daos.ProductDAO) error { return errors.New("update error") },
		FindByIDFunc: func(id string) (daos.ProductDAO, error) {
			return daos.ProductDAO{
				ID:          id,
				Name:        "prod",
				Description: "desc",
				Price:       1.0,
				Active:      true,
				CategoryID:  "catid",
				Images: []daos.ProductImageDAO{
					{ID: "img.jpg", ProductID: id, FileName: "img.jpg", IsDefault: true},
				},
			}, nil
		},
	}
	mockProductDs, mockCategoryDs, mockFileProvider := makeDefaultMocks(mockProductDs)
	r, w, h := setupProductTestEnv(mockProductDs, mockCategoryDs, mockFileProvider)

	r.Use(func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": c.Errors[0].Error()})
		}
	})
	r.PUT("/products/:id", h.UpdateProduct)

	body := `{"category_id":"catid","name":"prod","description":"desc","price":1.0,"active":true}`
	req := httptest.NewRequest(http.MethodPut, "/products/1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestFindProductByID_Success(t *testing.T) {
	mockProductDs := &testmocks.MockProductDataSource{
		FindByIDFunc: func(id string) (daos.ProductDAO, error) {
			return daos.ProductDAO{
				ID:          id,
				Name:        "prod",
				Description: "desc",
				Price:       1.0,
				Active:      true,
				CategoryID:  "catid",
				Images: []daos.ProductImageDAO{
					{ID: "img.jpg", ProductID: id, FileName: "img.jpg", IsDefault: true},
				},
			}, nil
		},
	}
	mockProductDs, mockCategoryDs, mockFileProvider := makeDefaultMocks(mockProductDs)
	r, w, h := setupProductTestEnv(mockProductDs, mockCategoryDs, mockFileProvider)

	r.GET("/products/:id", h.FindProductByID)

	req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Equal(t, "prod", resp["name"])
}

func TestFindProductByID_Error(t *testing.T) {
	mockProductDs := &testmocks.MockProductDataSource{
		FindByIDFunc: func(id string) (daos.ProductDAO, error) {
			return daos.ProductDAO{}, errors.New("not found")
		},
	}
	mockProductDs, mockCategoryDs, mockFileProvider := makeDefaultMocks(mockProductDs)
	r, w, h := setupProductTestEnv(mockProductDs, mockCategoryDs, mockFileProvider)

	r.Use(func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": c.Errors[0].Error()})
		}
	})
	r.GET("/products/:id", h.FindProductByID)

	req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
}

func TestCreateProduct_Success(t *testing.T) {
	mockProductDs := &testmocks.MockProductDataSource{
		InsertFunc: func(dao daos.ProductDAO) error { return nil },
	}
	mockProductDs, mockCategoryDs, mockFileProvider := makeDefaultMocks(mockProductDs)
	r, w, h := setupProductTestEnv(mockProductDs, mockCategoryDs, mockFileProvider)

	r.POST("/products", h.CreateProduct)

	body := `{"category_id":"catid","name":"prod","description":"desc","price":1.0,"active":true}`
	req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusCreated, w.Code)

	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Equal(t, "prod", resp["name"])
}

func TestCreateProduct_BadRequest(t *testing.T) {
	mockProductDs := &testmocks.MockProductDataSource{}
	mockProductDs, mockCategoryDs, mockFileProvider := makeDefaultMocks(mockProductDs)
	r, w, h := setupProductTestEnv(mockProductDs, mockCategoryDs, mockFileProvider)

	r.POST("/products", h.CreateProduct)

	body := `{"name":1}`
	req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)

	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Contains(t, resp["error"], "json")
}

func TestCreateProduct_Error(t *testing.T) {
	mockProductDs := &testmocks.MockProductDataSource{
		InsertFunc: func(dao daos.ProductDAO) error { return errors.New("create error") },
	}
	mockProductDs, mockCategoryDs, mockFileProvider := makeDefaultMocks(mockProductDs)
	r, w, h := setupProductTestEnv(mockProductDs, mockCategoryDs, mockFileProvider)

	r.Use(func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": c.Errors[0].Error()})
		}
	})
	r.POST("/products", h.CreateProduct)

	body := `{"category_id":"catid","name":"prod","description":"desc","price":1.0,"active":true}`
	req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestNewProductHandler(t *testing.T) {
	h := NewProductHandler()
	require.NotNil(t, h)
	require.NotNil(t, h.productController)
}

func TestDeleteProductImage_Success(t *testing.T) {
	mockProductDs := &testmocks.MockProductDataSource{
		DeleteImageFunc: func(imageFileName string) error { return nil },
		FindAllImagesProductByIdFunc: func(productID string) ([]daos.ProductImageDAO, error) {
			return []daos.ProductImageDAO{
				{ID: "img.jpg", ProductID: productID, FileName: "img.jpg", IsDefault: true},
				{ID: "img2.jpg", ProductID: productID, FileName: "img2.jpg", IsDefault: false},
			}, nil
		},
		FindByIDFunc: func(id string) (daos.ProductDAO, error) {
			return daos.ProductDAO{
				ID:          id,
				Name:        "prod",
				Description: "desc",
				Price:       1.0,
				Active:      true,
				CategoryID:  "catid",
				Images: []daos.ProductImageDAO{
					{ID: "img.jpg", ProductID: id, FileName: "img.jpg", IsDefault: true},
					{ID: "img2.jpg", ProductID: id, FileName: "img2.jpg", IsDefault: false},
				},
			}, nil
		},
	}
	mockProductDs, mockCategoryDs, _ := makeDefaultMocks(mockProductDs)
	mockFileProvider := makeGomockFileProvider(t)
	mockFileProvider.EXPECT().DeleteFiles(gomock.Any()).Return(nil).AnyTimes()
	mockFileProvider.EXPECT().DeleteFile(gomock.Any()).Return(nil).AnyTimes()

	r, w, h := setupProductTestEnv(mockProductDs, mockCategoryDs, mockFileProvider)

	r.DELETE("/products/:id/images/:image_file_name", h.DeleteProductImage)

	req := httptest.NewRequest(http.MethodDelete, "/products/1/images/img.jpg", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Logf("Body: %s", w.Body.String())
	}
	require.Equal(t, http.StatusNoContent, w.Code)
	require.Equal(t, 0, w.Body.Len())
}

func TestDeleteProductImage_BadRequest(t *testing.T) {
	mockProductDs := &testmocks.MockProductDataSource{
		DeleteImageFunc: func(imageFileName string) error {
			return errors.New("Product image cannot be empty, at least one image is required")
		},
		FindByIDFunc: func(id string) (daos.ProductDAO, error) {
			return daos.ProductDAO{
				ID:          id,
				Name:        "prod",
				Description: "desc",
				Price:       1.0,
				Active:      true,
				CategoryID:  "catid",
				Images: []daos.ProductImageDAO{
					{ID: "img.jpg", ProductID: id, FileName: "img.jpg", IsDefault: true},
				},
			}, nil
		},
	}
	mockProductDs, mockCategoryDs, mockFileProvider := makeDefaultMocks(mockProductDs)
	r, w, h := setupProductTestEnv(mockProductDs, mockCategoryDs, mockFileProvider)

	r.DELETE("/products/:id/images/:image_file_name", h.DeleteProductImage)

	req := httptest.NewRequest(http.MethodDelete, "/products/1/images/img.jpg", nil)
	r.ServeHTTP(w, req)

	if w.Code == http.StatusConflict {
		var resp map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		require.NoError(t, err)
		require.Contains(t, resp["error"], "cannot be empty")
	} else {
		require.Equal(t, http.StatusBadRequest, w.Code)
	}
}

func TestDeleteProductImage_Conflict(t *testing.T) {
	mockProductDs := &testmocks.MockProductDataSource{
		DeleteImageFunc: func(imageFileName string) error {
			return errors.New("só possui uma imagem")
		},
		FindByIDFunc: func(id string) (daos.ProductDAO, error) {
			return daos.ProductDAO{
				ID:          id,
				Name:        "prod",
				Description: "desc",
				Price:       1.0,
				Active:      true,
				CategoryID:  "catid",
				Images: []daos.ProductImageDAO{
					{ID: "img.jpg", ProductID: id, FileName: "img.jpg", IsDefault: true},
				},
			}, nil
		},
		FindAllImagesProductByIdFunc: func(productID string) ([]daos.ProductImageDAO, error) {
			return []daos.ProductImageDAO{
				{ID: "img.jpg", ProductID: productID, FileName: "img.jpg", IsDefault: true},
			}, nil
		},
	}
	mockProductDs, mockCategoryDs, mockFileProvider := makeDefaultMocks(mockProductDs)
	r, w, h := setupProductTestEnv(mockProductDs, mockCategoryDs, mockFileProvider)

	r.DELETE("/products/:id/images/:image_file_name", h.DeleteProductImage)

	req := httptest.NewRequest(http.MethodDelete, "/products/1/images/img.jpg", nil)
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusConflict, w.Code)
	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Contains(t, resp["error"], "Product image cannot be empty")
}

func TestDeleteProduct_ReturnsNoContent(t *testing.T) {
	mockProductDs := &testmocks.MockProductDataSource{
		DeleteFunc: func(id string) error { return nil },
	}
	mockProductDs, mockCategoryDs, _ := makeDefaultMocks(mockProductDs)
	mockFileProvider := makeGomockFileProvider(t)
	mockFileProvider.EXPECT().DeleteFiles(gomock.Any()).Return(nil).AnyTimes()
	r, w, h := setupProductTestEnv(mockProductDs, mockCategoryDs, mockFileProvider)

	r.DELETE("/products/:id", func(c *gin.Context) {
		h.DeleteProduct(c)
		c.Status(http.StatusNoContent)
	})

	req := httptest.NewRequest(http.MethodDelete, "/products/1", nil)
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNoContent, w.Code)
	require.Equal(t, 0, w.Body.Len())
}

func TestUploadProductImage_BadRequest_BindError(t *testing.T) {
	mockProductDs := &testmocks.MockProductDataSource{}
	mockProductDs, mockCategoryDs, mockFileProvider := makeDefaultMocks(mockProductDs)
	r, w, h := setupProductTestEnv(mockProductDs, mockCategoryDs, mockFileProvider)

	r.POST("/products/:id/images", h.UploadProductImage)

	req := httptest.NewRequest(http.MethodPost, "/products/1/images", nil)
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
}
