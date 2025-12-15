package exceptions

type ProductNotFoundException struct {
	Message string
}

type InvalidProductDataException struct {
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

type InvalidProductImageException struct {
	Message string
}

func (e *InvalidProductImageException) Error() string {
	if e.Message == "" {
		return "Invalid product image"
	}

	return e.Message
}

type ImageNotFoundException struct {
	Message string
}

func (e *ImageNotFoundException) Error() string {
	if e.Message == "" {
		return "Image not found"
	}

	return e.Message
}
