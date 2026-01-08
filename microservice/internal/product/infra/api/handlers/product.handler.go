package handlers

import (
	"io"
	"net/http"
	"strings"
	"tech_challenge/internal/product/application/controllers"
	"tech_challenge/internal/product/application/dtos"
	"tech_challenge/internal/product/infra/api/schemas"
	"tech_challenge/internal/product/infra/database/data_sources"
	shared_factories "tech_challenge/internal/shared/factories"
	"tech_challenge/internal/shared/infra/database"
	"tech_challenge/internal/shared/utils"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productController controllers.ProductController
}

func NewProductHandler() *ProductHandler {
	productDataSource := data_sources.NewProductDataSource(database.GetDB())
	categoryDataSource := data_sources.NewGormCategoryDataSource(database.GetDB())
	fileProvider := shared_factories.NewFileProvider()

	productController := controllers.NewProductController(productDataSource, categoryDataSource, fileProvider)

	return &ProductHandler{
		productController: *productController,
	}
}

// @Summary Create a new product
// @Tags Products
// @Accept json
// @Produce json
// @Param product body schemas.CreateProductSchema true "Product to create"
// @Success 201 {object} schemas.ProductResponseSchema
// @Failure 400 {object} schemas.InvalidProductDataErrorSchema
// @Failure 500 {object} schemas.ErrorMessageSchema
// @Router /products/ [post]
func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	var productRequestBody schemas.CreateProductSchema

	if err := ctx.ShouldBindJSON(&productRequestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productCreated, err := h.productController.Create(productRequestBody.ToDTO())
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, schemas.ToProductResponseSchema(productCreated))
}

// @Summary List all products
// @Tags Products
// @Produce json
// @Param category_id query string false "Filter by category ID"
// @Success 200 {array} schemas.ProductResponseSchema
// @Failure 500 {object} schemas.ErrorMessageSchema
// @Router /products/ [get]
func (h *ProductHandler) FindAllProducts(ctx *gin.Context) {
	categoryIdQuery := ctx.Query("category_id")

	var categoryId *string

	if categoryIdQuery != "" {
		categoryId = &categoryIdQuery
	} else {
		categoryId = nil
	}

	products, err := h.productController.FindAll(categoryId)

	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, schemas.ListProductsResponseSchema(products))
}

// @Summary Get a product by ID
// @Tags Products
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} schemas.ProductResponseSchema
// @Failure 400 {object} schemas.InvalidProductDataErrorSchema
// @Failure 404 {object} schemas.ProductNotFoundErrorSchema
// @Router /products/{id} [get]
func (h *ProductHandler) FindProductByID(ctx *gin.Context) {
	productId := ctx.Param("id")

	product, err := h.productController.FindByID(productId)

	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, schemas.ToProductResponseSchema(product))
}

// @Summary Update a product by ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param product body schemas.UpdateProductRequestBodySchema true "Updated product data"
// @Success 200 {object} schemas.ProductResponseSchema
// @Failure 400 {object} schemas.InvalidProductDataErrorSchema
// @Failure 404 {object} schemas.ProductNotFoundErrorSchema
// @Failure 500 {object} schemas.ErrorMessageSchema
// @Router /products/{id} [put]
func (h *ProductHandler) UpdateProduct(ctx *gin.Context) {
	productId := ctx.Param("id")

	var productBodyRequest schemas.UpdateProductRequestBodySchema

	if err := ctx.ShouldBindJSON(&productBodyRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	product, err := h.productController.Update(productBodyRequest.ToDTO(productId))

	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, schemas.ToProductResponseSchema(product))
}

// @Summary Add image to a product
// @Tags         Products
// @Accept       multipart/form-data
// @Produce      json
// @Param 		 id path string true "Product ID"
// @Param        image formData  file true "Image file"
// @Success 	 204   {object}  nil
// @Failure      400   {object}  schemas.InvalidProductDataErrorSchema
// @Failure      404   {object}  schemas.ProductNotFoundErrorSchema
// @Failure      500   {object}  schemas.ErrorMessageSchema
// @Router       /products/{id}/images [patch]
func (h *ProductHandler) UploadProductImage(ctx *gin.Context) {
	productId := ctx.Param("id")

	var fileUploaded schemas.UploadImageRequestSchema

	if err := ctx.ShouldBind(&fileUploaded); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid file upload",
		})
		return
	}

	file, err := fileUploaded.Image.Open()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to open file",
		})
		return
	}

	defer file.Close()

	if !utils.FileIsImage(*fileUploaded.Image) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid file type. Only images are allowed.",
		})
		return
	}

	fileName := fileUploaded.Image.Filename
	fileContent, err := io.ReadAll(file)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read file content",
		})
		return
	}

	err = h.productController.UploadImage(dtos.UploadProductImageDTO{
		ProductID:   productId,
		FileName:    fileName,
		FileContent: fileContent,
	})

	if err != nil {
		// Retorna erro 404 se for bucket inexistente ou inválido
		if strings.Contains(err.Error(), "NoSuchBucket") || strings.Contains(err.Error(), "InvalidBucketName") {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		// Sempre retorna a mensagem real do erro
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// @Summary Delete an image from a product
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param 		 id 			 path 	   string 					true "Product ID"
// @Param        image_file_name path      string                   true  "Nome do arquivo da imagem"
// @Success 	 204   {object}  nil
// @Failure      400   {object}  schemas.InvalidProductDataErrorSchema
// @Failure      404   {object}  schemas.ProductNotFoundErrorSchema
// @Failure      500   {object}  schemas.ErrorMessageSchema
// @Router       /products/{id}/images/{image_file_name} [delete]
func (h *ProductHandler) DeleteProductImage(ctx *gin.Context) {
	productId := ctx.Param("id")
	imageFileName := ctx.Param("image_file_name")

	err := h.productController.DeleteImage(productId, imageFileName)

	if err != nil {
		// Se for erro de imagem não pode ser removida por ser a última, retorna 409 (conflito)
		if strings.Contains(err.Error(), "cannot be empty") || strings.Contains(err.Error(), "só possui uma imagem") {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		// Retorna a mensagem de erro específica para o client
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// @Summary Delete a product by ID
// @Tags Products
// @Produce json
// @Param id path string true "Product ID"
// @Success 204 {object} nil
// @Failure 400 {object} schemas.InvalidProductDataErrorSchema
// @Failure 500 {object} schemas.ErrorMessageSchema
// @Router /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(ctx *gin.Context) {
	productId := ctx.Param("id")

	err := h.productController.Delete(productId)

	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// @Summary List all images of a product
// @Tags Products
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {array} schemas.ProductImageResponseSchema
// @Failure 404 {object} schemas.ErrorMessageSchema
// @Router /products/{id}/images [get]
func (h *ProductHandler) FindAllImagesProductById(ctx *gin.Context) {
	productId := ctx.Param("id")
	images, err := h.productController.FindAllImagesProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	// Retorna apenas o array de imagens
	ctx.JSON(http.StatusOK, images)
}
