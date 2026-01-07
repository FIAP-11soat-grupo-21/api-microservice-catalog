package handlers

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

type mockFileProvider struct {
	getPresignedURLFunc func(fileName string) (string, error)
}

func (m *mockFileProvider) GetPresignedURL(fileName string) (string, error) {
	if m.getPresignedURLFunc != nil {
		return m.getPresignedURLFunc(fileName)
	}
	return "", nil
}

func (m *mockFileProvider) DeleteFile(fileName string) error {
	return nil
}
func (m *mockFileProvider) DeleteFiles(fileNames []string) error {
	return nil
}
func (m *mockFileProvider) UploadFile(fileName string, file []byte) error {
	return nil
}

func TestFileHandler_FindFile_Success(t *testing.T) {
	mockProvider := &mockFileProvider{
		getPresignedURLFunc: func(fileName string) (string, error) {
			return "http://localhost/uploads/" + fileName, nil
		},
	}
	h := NewFileHandler(mockProvider)
	url, err := h.FindFile("test.txt")
	require.NoError(t, err)
	require.Equal(t, "http://localhost/uploads/test.txt", url)
}

func TestFileHandler_FindFile_Error(t *testing.T) {
	mockProvider := &mockFileProvider{
		getPresignedURLFunc: func(fileName string) (string, error) {
			return "", errors.New("fail")
		},
	}
	h := NewFileHandler(mockProvider)
	url, err := h.FindFile("notfound.txt")
	require.Error(t, err)
	require.Equal(t, "", url)
}
