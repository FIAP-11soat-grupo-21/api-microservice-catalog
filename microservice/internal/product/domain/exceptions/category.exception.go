package exceptions

type InvalidCategoryDataException struct {
	Message string
}

func (e *InvalidCategoryDataException) Error() string {
	if e.Message == "" {
		return "Invalid category data"
	}

	return e.Message
}

type CategoryNotFoundException struct {
	Message string
}

func (e *CategoryNotFoundException) Error() string {
	if e.Message == "" {
		return "Category not found"
	}

	return e.Message
}

type CategoryAlreadyExistsException struct {
	Message string
}

func (e *CategoryAlreadyExistsException) Error() string {
	if e.Message == "" {
		return "Category already exists"
	}

	return e.Message
}
