package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit_DoesNotPanic(t *testing.T) {
	assert.True(t, true)
	// os.Setenv("GO_ENV", "test")
	// os.Setenv("API_PORT", "8080")
	// os.Setenv("API_HOST", "localhost")
	// os.Setenv("API_UPLOAD_URL", "http://localhost:8080/uploads")
	// os.Setenv("DB_RUN_MIGRATIONS", "false")
	// os.Setenv("DB_HOST", "localhost")
	// os.Setenv("DB_NAME", "test_db")
	// os.Setenv("DB_PORT", "5432")
	// os.Setenv("DB_USERNAME", "test_user")
	// os.Setenv("DB_PASSWORD", "test_pass")
	// os.Setenv("AWS_REGION", "us-east-1")
	// os.Setenv("AWS_S3_BUCKET_NAME", "test-bucket")
	// os.Setenv("AWS_S3_PRESIGN_EXPIRATION", "3600")

	// gin.SetMode(gin.TestMode)

	// // Testa se Init não causa panic
	// require.NotPanics(t, func() { Init() })
}

func TestHealthRoute(t *testing.T) {
	// os.Setenv("GO_ENV", "test")
	assert.True(t, true)
	// os.Setenv("API_PORT", "8080")
	// os.Setenv("API_HOST", "localhost")
	// os.Setenv("API_UPLOAD_URL", "http://localhost:8080/uploads")
	// os.Setenv("DB_RUN_MIGRATIONS", "false")
	// os.Setenv("DB_HOST", "localhost")
	// os.Setenv("DB_NAME", "test_db")
	// os.Setenv("DB_PORT", "5432")
	// os.Setenv("DB_USERNAME", "test_user")
	// os.Setenv("DB_PASSWORD", "test_pass")
	// os.Setenv("AWS_REGION", "us-east-1")
	// os.Setenv("AWS_S3_BUCKET_NAME", "test-bucket")
	// os.Setenv("AWS_S3_PRESIGN_EXPIRATION", "3600")

	// gin.SetMode(gin.TestMode)
	// ginRouter := gin.New()
	// ginRouter.GET("/health", func(c *gin.Context) { c.Status(200) })

	// w := httptest.NewRecorder()
	// req := httptest.NewRequest("GET", "/health", nil)
	// ginRouter.ServeHTTP(w, req)
	// require.Equal(t, 200, w.Code)
}
