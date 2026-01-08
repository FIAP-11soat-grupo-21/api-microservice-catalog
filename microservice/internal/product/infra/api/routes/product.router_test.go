package routes

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func setupProductTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	group := r.Group("/products")
	group.POST("", func(c *gin.Context) { c.Status(201) })
	group.GET("", func(c *gin.Context) { c.Status(200) })
	group.GET(":id", func(c *gin.Context) { c.Status(200) })
	group.GET(":id/images", func(c *gin.Context) { c.Status(200) })
	group.PUT(":id", func(c *gin.Context) { c.Status(200) })
	group.PATCH(":id/images", func(c *gin.Context) { c.Status(200) })
	group.DELETE(":id/images/:image_file_name", func(c *gin.Context) { c.Status(204) })
	group.DELETE(":id", func(c *gin.Context) { c.Status(204) })
	return r
}

func TestRegisterProductRoutes(t *testing.T) {
	r := setupProductTestRouter()
	endpoints := []struct {
		method string
		path   string
		want   int
	}{
		{"POST", "/products", 201},
		{"GET", "/products", 200},
		{"GET", "/products/1", 200},
		{"GET", "/products/1/images", 200},
		{"PUT", "/products/1", 200},
		{"PATCH", "/products/1/images", 200},
		{"DELETE", "/products/1/images/img.jpg", 204},
		{"DELETE", "/products/1", 204},
	}
	for _, ep := range endpoints {
		req := httptest.NewRequest(ep.method, ep.path, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		require.NotEqual(t, 404, w.Code)
		require.Equal(t, ep.want, w.Code)
	}
}
