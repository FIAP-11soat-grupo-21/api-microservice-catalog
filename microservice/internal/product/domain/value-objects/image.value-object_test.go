package value_objects

import (
	"os"
	"strings"
	testenv "tech_challenge/internal/shared/test"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	testenv.SetupTestEnv()
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
