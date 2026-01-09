package file_service

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLocalFileProvider_UploadAndDeleteFile(t *testing.T) {
	provider := NewLocalFileProvider()
	fileName := "test_file.txt"
	fileContent := []byte("conteudo de teste")
	filePath := filepath.Join(provider.basePath, fileName)

	// Garante que o arquivo não existe antes do teste
	_ = os.Remove(filePath)

	// Upload deve funcionar
	err := provider.UploadFile(fileName, fileContent)
	require.NoError(t, err)
	// Upload de novo deve dar erro de existente
	err = provider.UploadFile(fileName, fileContent)
	require.ErrorIs(t, err, os.ErrExist)

	// Delete deve funcionar
	err = provider.DeleteFile(fileName)
	require.NoError(t, err)
	// Delete de novo deve dar erro de não existente
	err = provider.DeleteFile(fileName)
	require.ErrorIs(t, err, os.ErrNotExist)
}

func TestLocalFileProvider_fileExists(t *testing.T) {
	provider := NewLocalFileProvider()
	fileName := "test_exists.txt"
	filePath := filepath.Join(provider.basePath, fileName)
	_ = os.Remove(filePath)

	// Não existe
	require.False(t, provider.fileExists(fileName))
	// Cria o arquivo
	err := os.WriteFile(filePath, []byte("abc"), 0644)
	require.NoError(t, err)
	require.True(t, provider.fileExists(fileName))
	_ = os.Remove(filePath)
}
