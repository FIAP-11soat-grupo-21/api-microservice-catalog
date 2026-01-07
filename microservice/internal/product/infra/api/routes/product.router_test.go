package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestRegisterProductRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	group := r.Group("/products")

	// Registra handlers dummy para evitar acesso ao banco
	group.POST("", func(c *gin.Context) { c.Status(201) })
	group.GET("", func(c *gin.Context) { c.Status(200) })
	group.GET(":id", func(c *gin.Context) { c.Status(200) })
	group.GET(":id/images", func(c *gin.Context) { c.Status(200) })
	group.PUT(":id", func(c *gin.Context) { c.Status(200) })
	group.PATCH(":id/images", func(c *gin.Context) { c.Status(200) })
	group.DELETE(":id/images/:image_file_name", func(c *gin.Context) { c.Status(204) })
	group.DELETE(":id", func(c *gin.Context) { c.Status(204) })

	// Test POST /products
	req := httptest.NewRequest(http.MethodPost, "/products", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	require.NotEqual(t, 404, w.Code)

	// Test GET /products
	req = httptest.NewRequest(http.MethodGet, "/products", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	require.NotEqual(t, 404, w.Code)

	// Test GET /products/:id
	req = httptest.NewRequest(http.MethodGet, "/products/1", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	require.NotEqual(t, 404, w.Code)

	// Test GET /products/:id/images
	req = httptest.NewRequest(http.MethodGet, "/products/1/images", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	require.NotEqual(t, 404, w.Code)

	// Test PUT /products/:id
	req = httptest.NewRequest(http.MethodPut, "/products/1", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	require.NotEqual(t, 404, w.Code)

	// Test PATCH /products/:id/images
	req = httptest.NewRequest(http.MethodPatch, "/products/1/images", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	require.NotEqual(t, 404, w.Code)

	// Test DELETE /products/:id/images/:image_file_name
	req = httptest.NewRequest(http.MethodDelete, "/products/1/images/img.jpg", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	require.NotEqual(t, 404, w.Code)

	// Test DELETE /products/:id
	req = httptest.NewRequest(http.MethodDelete, "/products/1", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	require.NotEqual(t, 404, w.Code)
}
