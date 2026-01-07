package exceptions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCategoryNotFoundException_Error(t *testing.T) {
	req := require.New(t)
	req.Equal("Category not found", (&CategoryNotFoundException{}).Error())
	req.Equal("Custom", (&CategoryNotFoundException{Message: "Custom"}).Error())
}

func TestCategoryAlreadyExistsException_Error(t *testing.T) {
	req := require.New(t)
	req.Equal("Category already exists", (&CategoryAlreadyExistsException{}).Error())
	req.Equal("Custom", (&CategoryAlreadyExistsException{Message: "Custom"}).Error())
}

func TestCategoryHasProductsException_Error(t *testing.T) {
	req := require.New(t)
	req.Equal("Cannot delete category because there are products linked to it.", (&CategoryHasProductsException{}).Error())
	req.Equal("Custom", (&CategoryHasProductsException{Message: "Custom"}).Error())
}

func TestInvalidCategoryDataException_Error(t *testing.T) {
	req := require.New(t)
	req.Equal("Invalid category data", (&InvalidCategoryDataException{}).Error())
	req.Equal("Custom", (&InvalidCategoryDataException{Message: "Custom"}).Error())
}
