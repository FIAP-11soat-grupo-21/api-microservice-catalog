package handlers

import (
	"net/http"
	"tech_challenge/internal/product/application/controllers"
	"tech_challenge/internal/product/factories"
	"tech_challenge/internal/product/infra/api/schemas"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryController controllers.CategoryController
}

func NewCategoryHandler() *CategoryHandler {
	categoryDataSource := factories.NewCategoryDataSource()
	categoryController := controllers.NewCategoryController(categoryDataSource)

	return &CategoryHandler{
		categoryController: *categoryController,
	}
}

// @Summary List all categories
// @Tags Categories
// @Produce json
// @Success 200 {array} schemas.CategoryResponseSchema
// @Failure 500 {object} schemas.ErrorMessageSchema
// @Router /categories/ [get]
func (h *CategoryHandler) FindAllCategories(ctx *gin.Context) {
	categories, err := h.categoryController.FindAll()

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, schemas.ListToCategoryResponseSchema(categories))
}

// @Summary Get a Category by ID
// @Tags Categories
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} schemas.CategoryResponseSchema
// @Failure 400 {object} schemas.InvalidCategoryDataErrorSchema
// @Failure 404 {object} schemas.CategoryNotFoundErrorSchema
// @Router /categories/{id} [get]
func (h *CategoryHandler) FindCategoryByID(ctx *gin.Context) {
	categoryId := ctx.Param("id")

	category, err := h.categoryController.FindByID(categoryId)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, schemas.ToCategoryResponseSchema(category))
}

// @Summary CreateCategory a new category
// @Tags Categories
// @Accept json
// @Produce json
// @Param category body schemas.CreateCategorySchema true "Category to create"
// @Success 201 {object} schemas.CategoryResponseSchema
// @Failure 400 {object} schemas.InvalidCategoryDataErrorSchema
// @Failure 500 {object} schemas.ErrorMessageSchema
// @Router /categories/ [post]
func (h *CategoryHandler) CreateCategory(ctx *gin.Context) {
	var categoryRequestBody schemas.CreateCategorySchema

	if err := ctx.ShouldBindJSON(&categoryRequestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := h.categoryController.Create(categoryRequestBody.ToDTO())

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, schemas.ToCategoryResponseSchema(category))
}

// @Summary UpdateCategory a Category by ID
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Param category body schemas.UpdateCategoryRequestBodySchema true "Updated Category data"
// @Success 200 {object} schemas.CategoryResponseSchema
// @Failure 400 {object} schemas.InvalidCategoryDataErrorSchema
// @Failure 404 {object} schemas.CategoryNotFoundErrorSchema
// @Failure 500 {object} schemas.ErrorMessageSchema
// @Router /categories/{id} [put]
func (h *CategoryHandler) UpdateCategory(ctx *gin.Context) {
	categoryId := ctx.Param("id")

	var updateCategoryRequestBody schemas.UpdateCategoryRequestBodySchema

	if err := ctx.ShouldBindJSON(&updateCategoryRequestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := h.categoryController.Update(updateCategoryRequestBody.ToDTO(categoryId))

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, schemas.ToCategoryResponseSchema(category))

}

// @Summary DeleteCategory a Category by ID
// @Tags Categories
// @Produce json
// @Param id path string true "Category Order ID"
// @Success 204 {object} nil
// @Failure 400 {object} schemas.InvalidCategoryDataErrorSchema
// @Failure 500 {object} schemas.ErrorMessageSchema
// @Router /categories/{id} [delete]
func (h *CategoryHandler) DeleteCategory(ctx *gin.Context) {
	categoryId := ctx.Param("id")

	if err := h.categoryController.Delete(categoryId); err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
