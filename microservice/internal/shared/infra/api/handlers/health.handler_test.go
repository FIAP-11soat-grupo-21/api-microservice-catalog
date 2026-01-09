package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestHealthHandler_Health(t *testing.T) {
	gin.SetMode(gin.TestMode)
	h := NewHealthHandler()
	r := gin.New()
	r.GET("/health", h.Health)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/health", nil)
	r.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)
	require.Contains(t, w.Body.String(), "Serviço estável - OK")
	require.Contains(t, w.Body.String(), "timestamp")
}
