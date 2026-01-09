package exceptions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProductNotFoundException_Error(t *testing.T) {
	req := require.New(t)
	req.Equal("Product not found", (&ProductNotFoundException{}).Error())
	req.Equal("Custom", (&ProductNotFoundException{Message: "Custom"}).Error())
}

func TestInvalidProductDataException_Error(t *testing.T) {
	req := require.New(t)
	req.Equal("Invalid product data", (&InvalidProductDataException{}).Error())
	req.Equal("Custom", (&InvalidProductDataException{Message: "Custom"}).Error())
}

func TestInvalidProductImageException_Error(t *testing.T) {
	req := require.New(t)
	req.Equal("Invalid product image", (&InvalidProductImageException{}).Error())
	req.Equal("Custom", (&InvalidProductImageException{Message: "Custom"}).Error())
}

func TestImageNotFoundException_Error(t *testing.T) {
	req := require.New(t)
	req.Equal("Image not found", (&ImageNotFoundException{}).Error())
	req.Equal("Custom", (&ImageNotFoundException{Message: "Custom"}).Error())
}

func TestProductImagesNotFoundException_Error(t *testing.T) {
	req := require.New(t)
	req.Equal("No images found for this product.", (&ProductImagesNotFoundException{}).Error())
	req.Equal("Custom", (&ProductImagesNotFoundException{Message: "Custom"}).Error())
}

func TestProductImageCannotBeEmptyException_Error(t *testing.T) {
	req := require.New(t)
	req.Equal("Product image cannot be empty, at least one image is required", (&ProductImageCannotBeEmptyException{}).Error())
	req.Equal("Custom", (&ProductImageCannotBeEmptyException{Message: "Custom"}).Error())
}
