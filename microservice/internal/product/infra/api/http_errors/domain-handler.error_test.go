package http_errors

import (
	"errors"
	"net/http"
	"tech_challenge/internal/product/domain/exceptions"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestHandleDomainErrors(t *testing.T) {
	gin.SetMode(gin.TestMode)
	cases := []struct {
		err          error
		expectedCode int
	}{
		{&exceptions.ProductNotFoundException{}, http.StatusNotFound},
		{&exceptions.InvalidProductDataException{}, http.StatusBadRequest},
		{&exceptions.InvalidCategoryDataException{}, http.StatusBadRequest},
		{&exceptions.CategoryAlreadyExistsException{}, http.StatusConflict},
		{&exceptions.CategoryNotFoundException{}, http.StatusNotFound},
		{&exceptions.InvalidProductImageException{}, http.StatusBadRequest},
		{&exceptions.ImageNotFoundException{}, http.StatusNotFound},
		{&exceptions.CategoryHasProductsException{}, http.StatusBadRequest},
	}

	for _, c := range cases {
		w := &testResponseWriter{}
		ctx := gin.CreateTestContextOnly(w, gin.New())
		result := HandleDomainErrors(c.err, ctx)
		require.True(t, result)
		require.Equal(t, c.expectedCode, w.status)
	}

	// Test for error not handled
	w := &testResponseWriter{}
	ctx := gin.CreateTestContextOnly(w, gin.New())
	result := HandleDomainErrors(errors.New("other error"), ctx)
	require.False(t, result)
}

// Helper to capture status code

type testResponseWriter struct {
	status int
}

func (w *testResponseWriter) Header() http.Header         { return http.Header{} }
func (w *testResponseWriter) Write(b []byte) (int, error) { return len(b), nil }
func (w *testResponseWriter) WriteHeader(statusCode int)  { w.status = statusCode }
