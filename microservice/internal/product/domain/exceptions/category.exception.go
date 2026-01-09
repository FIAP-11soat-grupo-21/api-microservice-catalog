package exceptions

type CategoryNotFoundException struct {
	Message string
}
type CategoryAlreadyExistsException struct {
	Message string
}
type CategoryHasProductsException struct {
	Message string
}

func (e *CategoryNotFoundException) Error() string {
	if e.Message == "" {
		return "Category not found"
	}
	return e.Message
}

func (e *CategoryAlreadyExistsException) Error() string {
	if e.Message == "" {
		return "Category already exists"
	}
	return e.Message
}

type InvalidCategoryDataException struct {
	Message string
}

func (e *InvalidCategoryDataException) Error() string {
	if e.Message == "" {
		return "Invalid category data"
	}
	return e.Message
}
func (e *CategoryHasProductsException) Error() string {
	if e.Message == "" {
		return "Cannot delete category because there are products linked to it."
	}
	return e.Message
}
