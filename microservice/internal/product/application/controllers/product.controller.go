package controllers

import (
	"fmt"
	"tech_challenge/internal/product/application/dtos"
	"tech_challenge/internal/product/application/gateways"
	"tech_challenge/internal/product/application/presenters"

	"tech_challenge/internal/product/interfaces"
	use_cases "tech_challenge/internal/product/use_cases/product"
	shared_interfaces "tech_challenge/internal/shared/interfaces"
)

type ProductController struct {
	productGateway  gateways.ProductGateway
	categoryGateway gateways.CategoryGateway
}

func NewProductController(
	productDataSource interfaces.IProductDataSource,
	categoryDataSource interfaces.ICategoryDataSource,
	fileService shared_interfaces.IFileProvider,
) *ProductController {
	return &ProductController{
		productGateway:  *gateways.NewProductGateway(productDataSource, fileService),
		categoryGateway: gateways.NewCategoryGateway(categoryDataSource),
	}
}

func (c *ProductController) Create(productDTO dtos.CreateProductDTO) (dtos.ProductResultDTO, error) {
	createProductUseCase := use_cases.NewCreateProductUseCase(c.productGateway, c.categoryGateway)

	product, err := createProductUseCase.Execute(productDTO)

	if err != nil {
		return dtos.ProductResultDTO{}, err
	}

	return presenters.ProductFromDomainToResultDTO(product), nil
}

func (c *ProductController) FindByID(productID string) (dtos.ProductResultDTO, error) {
	findProductUseCase := use_cases.NewFindProductByIDUseCase(c.productGateway)

	product, err := findProductUseCase.Execute(productID)

	if err != nil {
		return dtos.ProductResultDTO{}, err
	}

	return presenters.ProductFromDomainToResultDTO(product), nil
}

func (c *ProductController) FindAll(categoryID *string) ([]dtos.ProductResultDTO, error) {
	findAllProductsUseCase := use_cases.NewFindAllProductsUseCase(c.productGateway, c.categoryGateway)

	products, err := findAllProductsUseCase.Execute(categoryID)

	if err != nil {
		return nil, err
	}

	return presenters.ListProductDomainToResultDTO(products), nil
}

func (c *ProductController) Update(productDTO dtos.UpdateProductDTO) (dtos.ProductResultDTO, error) {
	updateProductUseCase := use_cases.NewUpdateProductUseCase(c.productGateway)

	product, err := updateProductUseCase.Execute(productDTO)

	if err != nil {
		return dtos.ProductResultDTO{}, err
	}

	return presenters.ProductFromDomainToResultDTO(product), nil
}

func (c *ProductController) UploadImage(uploadDTO dtos.UploadProductImageDTO) error {
	fmt.Println("[UploadImage] Iniciando upload para produto:", uploadDTO.ProductID, "Arquivo:", uploadDTO.FileName)
	uploadProductImageUseCase := use_cases.NewUploadProductImageUseCase(c.productGateway)
	err := uploadProductImageUseCase.Execute(uploadDTO)
	if err != nil {
		fmt.Printf("[UploadImage] Erro ao executar use case: %v (tipo: %T)\n", err, err)
	} else {
		fmt.Println("[UploadImage] Upload realizado com sucesso!")
	}
	return err
}

func (c *ProductController) DeleteImage(productID string, imageFileName string) error {
	deleteProductImageUseCase := use_cases.NewDeleteProductImageUseCase(c.productGateway)

	return deleteProductImageUseCase.Execute(productID, imageFileName)
}

func (c *ProductController) Delete(productID string) error {
	deleteProductUseCase := use_cases.NewDeleteProductUseCase(c.productGateway)

	return deleteProductUseCase.Execute(productID)
}

func (c *ProductController) FindAllImagesProductById(productId string) ([]dtos.ProductImageDTO, error) {
	product, err := c.productGateway.FindAllImagesProductById(productId)
	if err != nil {
		return nil, err
	}
	return presenters.ProductImagesFromDomainToResultDTO(product.Images), nil
}
