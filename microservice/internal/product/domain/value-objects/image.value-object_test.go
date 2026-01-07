package value_objects

import (
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	_ = godotenv.Load("../../../.env.local.example")
	os.Setenv("GO_ENV", "test")
	os.Setenv("API_PORT", "8080")
	os.Setenv("API_HOST", "localhost")
	os.Setenv("API_UPLOAD_URL", "http://localhost:8080/uploads")
	os.Setenv("DB_RUN_MIGRATIONS", "false")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_NAME", "test_db")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USERNAME", "test_user")
	os.Setenv("DB_PASSWORD", "test_pass")
	os.Setenv("AWS_REGION", "us-east-1")
	os.Setenv("AWS_S3_BUCKET_NAME", "test-bucket")
	os.Setenv("AWS_S3_PRESIGN_EXPIRATION", "3600")
	code := m.Run()
	os.Exit(code)
}
func TestNewImage_InvalidEmpty(t *testing.T) {
	_, err := NewImage("")
	require.Error(t, err)
}

func TestNewImage_InvalidLong(t *testing.T) {
	longName := strings.Repeat("a", 300) + ".jpg"
	_, err := NewImage(longName)
	require.Error(t, err)
}

func TestNewImage_Valid(t *testing.T) {
	img, err := NewImage("produto.jpg")
	require.NoError(t, err)
	require.Contains(t, img.FileName, "produto_")
	require.Contains(t, img.Url, img.FileName)
	require.True(t, img.IsDefault)
}

func TestNewImageWithFileNameAndUrl_Invalid(t *testing.T) {
	_, err := NewImageWithFileNameAndUrl("", "url", true)
	require.Error(t, err)
	_, err = NewImageWithFileNameAndUrl("file.jpg", "", true)
	require.Error(t, err)
}

func TestNewImageWithFileNameAndUrl_Valid(t *testing.T) {
	img, err := NewImageWithFileNameAndUrl("file.jpg", "http://host/file.jpg", false)
	require.NoError(t, err)
	require.Equal(t, "file.jpg", img.FileName)
	require.Equal(t, "http://host/file.jpg", img.Url)
	require.False(t, img.IsDefault)
}

func TestImage_Value(t *testing.T) {
	img := Image{
		FileName:  "file.jpg",
		Url:       "http://host/file.jpg",
		IsDefault: false,
	}
	val := img.Value()
	require.Equal(t, img.FileName, val.FileName)
	require.Equal(t, img.Url, val.Url)
}
