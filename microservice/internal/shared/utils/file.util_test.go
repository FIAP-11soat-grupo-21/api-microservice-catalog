package utils

import (
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileIsImage_ValidTypes(t *testing.T) {
	validTypes := []string{"image/jpeg", "image/jpg", "image/png", "image/gif", "image/webp"}
	for _, ct := range validTypes {
		h := make(map[string][]string)
		h["Content-Type"] = []string{ct}
		fh := multipart.FileHeader{Header: h}
		require.True(t, FileIsImage(fh), ct)
	}
}

func TestFileIsImage_InvalidType(t *testing.T) {
	h := make(map[string][]string)
	h["Content-Type"] = []string{"application/pdf"}
	fh := multipart.FileHeader{Header: h}
	require.False(t, FileIsImage(fh))
}
