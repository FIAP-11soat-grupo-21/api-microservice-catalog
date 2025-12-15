package http_errors

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"tech_challenge/internal/product/domain/exceptions"
)

func HandleDomainErrors(err error, ctx *gin.Context) bool {
	switch e := err.(type) {
	case *exceptions.ProductNotFoundException:
		ctx.JSON(http.StatusNotFound, gin.H{"error": e.Error()})
		return true

	case *exceptions.InvalidProductDataException:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return true

	case *exceptions.InvalidCategoryDataException:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return true

	case *exceptions.CategoryAlreadyExistsException:
		ctx.JSON(http.StatusConflict, gin.H{"error": e.Error()})
		return true

	case *exceptions.CategoryNotFoundException:
		ctx.JSON(http.StatusNotFound, gin.H{"error": e.Error()})
		return true

	case *exceptions.InvalidProductImageException:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return true

	case *exceptions.ImageNotFoundException:
		ctx.JSON(http.StatusNotFound, gin.H{"error": e.Error()})
		return true

	}

	return false
}
