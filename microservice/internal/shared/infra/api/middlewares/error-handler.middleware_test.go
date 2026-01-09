package middlewares

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestErrorHandlerMiddleware_InternalServerError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(ErrorHandlerMiddleware())
	r.GET("/fail", func(c *gin.Context) {
		_ = c.Error(errors.New("unexpected error"))
	})
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/fail", nil)
	r.ServeHTTP(w, req)
	require.Equal(t, http.StatusInternalServerError, w.Code)
	require.Contains(t, w.Body.String(), "Internal server error")
}

func TestErrorHandlerMiddleware_NoError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(ErrorHandlerMiddleware())
	r.GET("/ok", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "ok"})
	})
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/ok", nil)
	r.ServeHTTP(w, req)
	require.Equal(t, 200, w.Code)
	require.Contains(t, w.Body.String(), "ok")
}
