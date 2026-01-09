package exceptions

type ProductNotFoundException struct {
	Message string
}

type InvalidProductDataException struct {
	Message string
}
type InvalidProductImageException struct {
	Message string
}
type ImageNotFoundException struct {
	Message string
}
type ProductImagesNotFoundException struct {
	Message string
}
type ProductImageCannotBeEmptyException struct {
	Message string
}

func (e *ProductNotFoundException) Error() string {
	if e.Message == "" {
		return "Product not found"
	}
	return e.Message
}

func (e *InvalidProductDataException) Error() string {
	if e.Message == "" {
		return "Invalid product data"
	}

	return e.Message
}

func (e *InvalidProductImageException) Error() string {
	if e.Message == "" {
		return "Invalid product image"
	}

	return e.Message
}

func (e *ImageNotFoundException) Error() string {
	if e.Message == "" {
		return "Image not found"
	}

	return e.Message
}
func (e *ProductImagesNotFoundException) Error() string {
	if e.Message == "" {
		return "No images found for this product."
	}
	return e.Message
}
func (e *ProductImageCannotBeEmptyException) Error() string {
	if e.Message == "" {
		return "Product image cannot be empty, at least one image is required"
	}
	return e.Message
}
