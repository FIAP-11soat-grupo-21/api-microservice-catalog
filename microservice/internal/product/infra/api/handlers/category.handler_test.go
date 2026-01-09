package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"

	"tech_challenge/internal/product/daos"
	testmocks "tech_challenge/internal/shared/test"
)

func setupCategoryTestEnv(mockCategoryDs *testmocks.MockCategoryDataSource) (*gin.Engine, *httptest.ResponseRecorder, *CategoryHandler) {
	gin.SetMode(gin.TestMode)
	h := setupCategoryHandlerWithFakeGateway(mockCategoryDs)
	r := gin.New()
	w := httptest.NewRecorder()
	return r, w, h
}

func TestFindAllCategories_Success(t *testing.T) {
	mockCategoryDs := &testmocks.MockCategoryDataSource{
		FindAllFunc: func() ([]daos.CategoryDAO, error) {
			return []daos.CategoryDAO{{ID: "1", Name: "Bebidas", Active: true}}, nil
		},
	}
	r, w, h := setupCategoryTestEnv(mockCategoryDs)

	r.GET("/categories", h.FindAllCategories)

	req := httptest.NewRequest(http.MethodGet, "/categories", nil)
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp []map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Len(t, resp, 1)
	require.Equal(t, "Bebidas", resp[0]["name"])
}

func TestFindAllCategories_Error(t *testing.T) {
	mockCategoryDs := &testmocks.MockCategoryDataSource{
		FindAllFunc: func() ([]daos.CategoryDAO, error) {
			return nil, errors.New("mock error")
		},
	}

	r, w, h := setupCategoryTestEnv(mockCategoryDs)

	r.Use(func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": c.Errors[0].Error()})
		}
	})
	r.GET("/categories", h.FindAllCategories)

	req := httptest.NewRequest(http.MethodGet, "/categories", nil)
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestFindCategoryByID_Success(t *testing.T) {
	mockCategoryDs := &testmocks.MockCategoryDataSource{
		FindByIDFunc: func(id string) (daos.CategoryDAO, error) {
			return daos.CategoryDAO{ID: id, Name: "Bebidas", Active: true}, nil
		},
	}
	r, w, h := setupCategoryTestEnv(mockCategoryDs)

	r.GET("/categories/:id", h.FindCategoryByID)

	req := httptest.NewRequest(http.MethodGet, "/categories/1", nil)
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Equal(t, "Bebidas", resp["name"])
}

func TestFindCategoryByID_Error(t *testing.T) {
	mockCategoryDs := &testmocks.MockCategoryDataSource{
		FindByIDFunc: func(id string) (daos.CategoryDAO, error) {
			return daos.CategoryDAO{}, errors.New("not found")
		},
	}
	r, w, h := setupCategoryTestEnv(mockCategoryDs)

	r.Use(func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": c.Errors[0].Error()})
		}
	})
	r.GET("/categories/:id", h.FindCategoryByID)

	req := httptest.NewRequest(http.MethodGet, "/categories/1", nil)
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
}

func TestCreateCategory_Success(t *testing.T) {
	mockCategoryDs := &testmocks.MockCategoryDataSource{
		InsertFunc: func(dao daos.CategoryDAO) error { return nil },
	}
	r, w, h := setupCategoryTestEnv(mockCategoryDs)

	r.POST("/categories", h.CreateCategory)

	body := `{"name":"Bebidas","active":true}`
	req := httptest.NewRequest(http.MethodPost, "/categories", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusCreated, w.Code)

	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Equal(t, "Bebidas", resp["name"])
}

func TestCreateCategory_BadRequest(t *testing.T) {
	mockCategoryDs := &testmocks.MockCategoryDataSource{}
	r, w, h := setupCategoryTestEnv(mockCategoryDs)

	r.POST("/categories", h.CreateCategory)

	body := `{"name":1}` // inválido
	req := httptest.NewRequest(http.MethodPost, "/categories", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)

	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Contains(t, resp["error"], "json")
}

func TestCreateCategory_Error(t *testing.T) {
	mockCategoryDs := &testmocks.MockCategoryDataSource{
		InsertFunc: func(dao daos.CategoryDAO) error { return errors.New("create error") },
	}
	r, w, h := setupCategoryTestEnv(mockCategoryDs)

	r.Use(func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": c.Errors[0].Error()})
		}
	})
	r.POST("/categories", h.CreateCategory)

	body := `{"name":"Bebidas","active":true}`
	req := httptest.NewRequest(http.MethodPost, "/categories", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestUpdateCategory_Success(t *testing.T) {
	mockCategoryDs := &testmocks.MockCategoryDataSource{
		UpdateFunc: func(dao daos.CategoryDAO) error { return nil },
		FindByIDFunc: func(id string) (daos.CategoryDAO, error) {
			return daos.CategoryDAO{ID: id, Name: "Bebidas", Active: true}, nil
		},
	}
	r, w, h := setupCategoryTestEnv(mockCategoryDs)

	r.PUT("/categories/:id", h.UpdateCategory)

	body := `{"name":"Bebidas","active":true}`
	req := httptest.NewRequest(http.MethodPut, "/categories/1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Equal(t, "Bebidas", resp["name"])
}

func TestUpdateCategory_BadRequest(t *testing.T) {
	mockCategoryDs := &testmocks.MockCategoryDataSource{}
	r, w, h := setupCategoryTestEnv(mockCategoryDs)

	r.PUT("/categories/:id", h.UpdateCategory)

	body := `{"name":1}` // inválido
	req := httptest.NewRequest(http.MethodPut, "/categories/1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)

	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Contains(t, resp["error"], "json")
}

func TestUpdateCategory_Error(t *testing.T) {
	mockCategoryDs := &testmocks.MockCategoryDataSource{
		UpdateFunc: func(dao daos.CategoryDAO) error { return errors.New("update error") },
		FindByIDFunc: func(id string) (daos.CategoryDAO, error) {
			return daos.CategoryDAO{ID: id, Name: "Bebidas", Active: true}, nil
		},
	}
	r, w, h := setupCategoryTestEnv(mockCategoryDs)

	r.Use(func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": c.Errors[0].Error()})
		}
	})
	r.PUT("/categories/:id", h.UpdateCategory)

	body := `{"name":"Bebidas","active":true}`
	req := httptest.NewRequest(http.MethodPut, "/categories/1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestDeleteCategory_Success(t *testing.T) {
	mockCategoryDs := &testmocks.MockCategoryDataSource{
		DeleteFunc: func(id string) error { return nil },
		FindByIDFunc: func(id string) (daos.CategoryDAO, error) {
			return daos.CategoryDAO{ID: id, Name: "Bebidas", Active: true}, nil
		},
	}
	r, w, h := setupCategoryTestEnv(mockCategoryDs)

	r.DELETE("/categories/:id", h.DeleteCategory)

	req := httptest.NewRequest(http.MethodDelete, "/categories/1", nil)
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNoContent, w.Code)
}

func TestDeleteCategory_Error(t *testing.T) {
	mockCategoryDs := &testmocks.MockCategoryDataSource{
		DeleteFunc: func(id string) error { return errors.New("delete error") },
		FindByIDFunc: func(id string) (daos.CategoryDAO, error) {
			return daos.CategoryDAO{ID: id, Name: "Bebidas", Active: true}, nil
		},
	}
	r, w, h := setupCategoryTestEnv(mockCategoryDs)

	r.Use(func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": c.Errors[0].Error()})
		}
	})
	r.DELETE("/categories/:id", h.DeleteCategory)

	req := httptest.NewRequest(http.MethodDelete, "/categories/1", nil)
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestNewCategoryHandler(t *testing.T) {
	h := NewCategoryHandler()
	require.NotNil(t, h)
	require.NotNil(t, h.categoryController)
}
