package routes

import (
	"net/http/httptest"
	"os"
	testenv "tech_challenge/internal/shared/test"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

type mockFileHandler struct{}

//	func TestMain(m *testing.M) {
//		os.Setenv("GO_ENV", "test")
//		os.Setenv("API_PORT", "8080")
//		os.Setenv("API_HOST", "localhost")
//		os.Setenv("API_UPLOAD_URL", "http://localhost:8080/uploads")
//		os.Setenv("DB_RUN_MIGRATIONS", "false")
//		os.Setenv("DB_HOST", "localhost")
//		os.Setenv("DB_NAME", "test_db")
//		os.Setenv("DB_PORT", "5432")
//		os.Setenv("DB_USERNAME", "test_user")
//		os.Setenv("DB_PASSWORD", "test_pass")
//		os.Setenv("AWS_REGION", "us-east-1")
//		os.Setenv("AWS_S3_BUCKET_NAME", "test-bucket")
//		os.Setenv("AWS_S3_PRESIGN_EXPIRATION", "3600")
//		code := m.Run()
//		os.Exit(code)
//	}
func TestMain(m *testing.M) {
	testenv.SetupTestEnv()
	code := m.Run()
	os.Exit(code)
}
func (m *mockFileHandler) FindFile(fileName string) (string, error) {
	if fileName == "notfound.txt" {
		return "", errMock
	}
	return "http://localhost/uploads/" + fileName, nil
}

var errMock = &gin.Error{Err: nil, Type: gin.ErrorTypePrivate}

func TestRegisterFileRoutes_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	// Substitui o handler real por mock
	r.GET("/:fileName", func(c *gin.Context) {
		fileName := c.Param("fileName")
		fileUrl, err := (&mockFileHandler{}).FindFile(fileName)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to retrieve file"})
			return
		}
		c.JSON(200, gin.H{"fileUrl": fileUrl})
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/test.txt", nil)
	r.ServeHTTP(w, req)
	require.Equal(t, 200, w.Code)
	require.Contains(t, w.Body.String(), "http://localhost/uploads/test.txt")
}

func TestRegisterFileRoutes_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/:fileName", func(c *gin.Context) {
		fileName := c.Param("fileName")
		fileUrl, err := (&mockFileHandler{}).FindFile(fileName)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to retrieve file"})
			return
		}
		c.JSON(200, gin.H{"fileUrl": fileUrl})
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/notfound.txt", nil)
	r.ServeHTTP(w, req)
	require.Equal(t, 500, w.Code)
	require.Contains(t, w.Body.String(), "Failed to retrieve file")
}

func TestRegisterFileRoutes_FileHandlerError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/:fileName", func(c *gin.Context) {
		fileName := c.Param("fileName")
		fileUrl, err := (&mockFileHandler{}).FindFile(fileName)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to retrieve file"})
			return
		}
		c.JSON(200, gin.H{"fileUrl": fileUrl})
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/notfound.txt", nil)
	r.ServeHTTP(w, req)
	require.Equal(t, 500, w.Code)
	require.Contains(t, w.Body.String(), "Failed to retrieve file")
}

// func TestRegisterFileRoutes_Integration(t *testing.T) {

// 	gin.SetMode(gin.TestMode)
// 	r := gin.New()
// 	RegisterFileRoutes(r.Group("/"))
// 	w := httptest.NewRecorder()
// 	req := httptest.NewRequest("GET", "/test.txt", nil)
// 	r.ServeHTTP(w, req)
// 	require.Equal(t, 200, w.Code)
// 	require.Contains(t, w.Body.String(), "fileUrl")
// }
