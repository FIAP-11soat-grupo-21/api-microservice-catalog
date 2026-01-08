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

func setupTestRouter() *gin.Engine {
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
	return r
}

func TestRegisterFileRoutes_Success(t *testing.T) {
	r := setupTestRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/test.txt", nil)
	r.ServeHTTP(w, req)
	require.Equal(t, 200, w.Code)
	require.Contains(t, w.Body.String(), "http://localhost/uploads/test.txt")
}

func TestRegisterFileRoutes_Error(t *testing.T) {
	r := setupTestRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/notfound.txt", nil)
	r.ServeHTTP(w, req)
	require.Equal(t, 500, w.Code)
	require.Contains(t, w.Body.String(), "Failed to retrieve file")
}

func TestRegisterFileRoutes_FileHandlerError(t *testing.T) {
	r := setupTestRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/notfound.txt", nil)
	r.ServeHTTP(w, req)
	require.Equal(t, 500, w.Code)
	require.Contains(t, w.Body.String(), "Failed to retrieve file")
}
