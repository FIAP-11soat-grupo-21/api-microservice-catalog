package controllers

import (
	"tech_challenge/internal/product/application/dtos"
	"tech_challenge/internal/product/application/gateways"
	"tech_challenge/internal/product/application/presenters"
	"tech_challenge/internal/product/interfaces"
	use_cases "tech_challenge/internal/product/use_cases/product"
	shared_interfaces "tech_challenge/internal/shared/interfaces"
)

type ProductController struct {
	gateway gateways.ProductGateway
}

func NewProductController(dataSource interfaces.IProductDataSource, fileService shared_interfaces.IFileProvider) *ProductController {
	return &ProductController{
		gateway: *gateways.NewProductGateway(dataSource, fileService),
	}
}

func (c *ProductController) Create(productDTO dtos.CreateProductDTO) (dtos.ProductResultDTO, error) {
	createProductUseCase := use_cases.NewCreateProductUseCase(c.gateway)

	product, err := createProductUseCase.Execute(productDTO)

	if err != nil {
		return dtos.ProductResultDTO{}, err
	}

	return presenters.ProductFromDomainToResultDTO(product), nil
}

func (c *ProductController) FindByID(productID string) (dtos.ProductResultDTO, error) {
	findProductUseCase := use_cases.NewFindProductByIDUseCase(c.gateway)

	product, err := findProductUseCase.Execute(productID)

	if err != nil {
		return dtos.ProductResultDTO{}, err
	}

	return presenters.ProductFromDomainToResultDTO(product), nil
}

func (c *ProductController) FindAll(categoryID *string) ([]dtos.ProductResultDTO, error) {
	findAllProductsUseCase := use_cases.NewFindAllProductsUseCase(c.gateway)

	products, err := findAllProductsUseCase.Execute(categoryID)

	if err != nil {
		return nil, err
	}

	return presenters.ListProductDomainToResultDTO(products), nil
}

func (c *ProductController) Update(productDTO dtos.UpdateProductDTO) (dtos.ProductResultDTO, error) {
	updateProductUseCase := use_cases.NewUpdateProductUseCase(c.gateway)

	product, err := updateProductUseCase.Execute(productDTO)

	if err != nil {
		return dtos.ProductResultDTO{}, err
	}

	return presenters.ProductFromDomainToResultDTO(product), nil
}

func (c *ProductController) UploadImage(uploadDTO dtos.UploadProductImageDTO) error {
	uploadProductImageUseCase := use_cases.NewUploadProductImageUseCase(c.gateway)

	return uploadProductImageUseCase.Execute(uploadDTO)
}

func (c *ProductController) DeleteImage(productID string, imageFileName string) error {
	deleteProductImageUseCase := use_cases.NewDeleteProductImageUseCase(c.gateway)

	return deleteProductImageUseCase.Execute(productID, imageFileName)
}

func (c *ProductController) Delete(productID string) error {
	deleteProductUseCase := use_cases.NewDeleteProductUseCase(c.gateway)

	return deleteProductUseCase.Execute(productID)
}
