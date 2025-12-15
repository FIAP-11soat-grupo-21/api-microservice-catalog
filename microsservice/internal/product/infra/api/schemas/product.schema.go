package schemas

import (
	"mime/multipart"
	"tech_challenge/internal/product/application/dtos"
)

type CreateProductSchema struct {
	CategoryID  string  `json:"category_id" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price"`
	Active      bool    `json:"active"`
}

func (s *CreateProductSchema) ToDTO() dtos.CreateProductDTO {
	return dtos.CreateProductDTO{
		CategoryID:  s.CategoryID,
		Name:        s.Name,
		Description: s.Description,
		Price:       s.Price,
		Active:      s.Active,
	}
}

type UpdateProductRequestBodySchema struct {
	CategoryID  string  `json:"category_id" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price"`
	Active      bool    `json:"active"`
}

func (s *UpdateProductRequestBodySchema) ToDTO(productID string) dtos.UpdateProductDTO {
	return dtos.UpdateProductDTO{
		ID:          productID,
		CategoryID:  s.CategoryID,
		Name:        s.Name,
		Description: s.Description,
		Price:       s.Price,
		Active:      s.Active,
	}
}

type UploadImageRequestSchema struct {
	Image *multipart.FileHeader `form:"image" binding:"required"`
}

type ImageResponseSchema struct {
	FileName string `json:"file_name" example:"image.jpg"`
	Url      string `json:"url" example:"https://example.com/image.jpg"`
}

type ProductResponseSchema struct {
	ID          string                `json:"id" example:"76fbddb3-3e2f-4f5f-a4e1-30a0a2384eae"`
	Name        string                `json:"name" example:"X-Salada"`
	Description string                `json:"description" example:"Lanche com carne, queijo, alface e tomate"`
	Price       float64               `json:"price" example:"20.50"`
	Active      bool                  `json:"active" example:"true"`
	CategoryID  string                `json:"category_id" example:"2cb7f56d-89a1-4e60-b488-65dc4ffacbc6"`
	Images      []ImageResponseSchema `json:"images"`
}

func ToProductResponseSchema(product dtos.ProductResultDTO) ProductResponseSchema {
	images := make([]ImageResponseSchema, len(product.Images))

	for i, image := range product.Images {
		images[i] = ImageResponseSchema{
			FileName: image.FileName,
			Url:      image.Url,
		}
	}

	return ProductResponseSchema{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Active:      product.Active,
		CategoryID:  product.CategoryID,
		Images:      images,
	}
}

func ListProductsResponseSchema(products []dtos.ProductResultDTO) []ProductResponseSchema {
	response := make([]ProductResponseSchema, len(products))

	for i, product := range products {
		response[i] = ToProductResponseSchema(product)
	}

	return response
}

type ProductNotFoundErrorSchema struct {
	Error string `json:"error" example:"Product not found"`
}

type ProductAlreadyExistsErrorSchema struct {
	Error string `json:"error" example:"Product already exists"`
}

type InvalidProductDataErrorSchema struct {
	Error string `json:"error" example:"Invalid product data"`
}

type ErrorMessageSchema struct {
	Error string `json:"error" example:"Internal server error"`
}
