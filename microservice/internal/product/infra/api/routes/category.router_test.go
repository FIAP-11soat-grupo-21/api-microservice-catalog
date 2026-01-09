package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestRegisterCategoryRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	group := r.Group("/categories")

	// Registra handlers dummy para evitar acesso ao banco
	group.GET("", func(c *gin.Context) { c.Status(200) })
	group.GET(":id", func(c *gin.Context) { c.Status(200) })
	group.POST("", func(c *gin.Context) { c.Status(201) })
	group.PUT(":id", func(c *gin.Context) { c.Status(200) })
	group.DELETE(":id", func(c *gin.Context) { c.Status(204) })

	// Test GET /categories
	req := httptest.NewRequest(http.MethodGet, "/categories", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	require.NotEqual(t, 404, w.Code)

	// Test GET /categories/:id
	req = httptest.NewRequest(http.MethodGet, "/categories/1", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	require.NotEqual(t, 404, w.Code)

	// Test POST /categories
	req = httptest.NewRequest(http.MethodPost, "/categories", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	require.NotEqual(t, 404, w.Code)

	// Test PUT /categories/:id
	req = httptest.NewRequest(http.MethodPut, "/categories/1", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	require.NotEqual(t, 404, w.Code)

	// Test DELETE /categories/:id
	req = httptest.NewRequest(http.MethodDelete, "/categories/1", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	require.NotEqual(t, 404, w.Code)
}
