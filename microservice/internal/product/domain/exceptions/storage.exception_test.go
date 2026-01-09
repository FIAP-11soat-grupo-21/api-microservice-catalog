package exceptions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeleteImagesStorageException_Error(t *testing.T) {
	req := require.New(t)
	req.Equal("Failed to delete file(s) from storage", (&DeleteImagesStorageException{}).Error())
	req.Equal("Custom", (&DeleteImagesStorageException{Message: "Custom"}).Error())
}

func TestBucketNotFoundException_Error(t *testing.T) {
	req := require.New(t)
	req.Equal("Bucket S3 não existe ou é inválido", (&BucketNotFoundException{}).Error())
	req.Equal("Custom", (&BucketNotFoundException{Message: "Custom"}).Error())
}
