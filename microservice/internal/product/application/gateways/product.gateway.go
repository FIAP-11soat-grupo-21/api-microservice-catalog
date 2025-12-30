package gateways

import (
	"fmt"
	"tech_challenge/internal/product/daos"
	"tech_challenge/internal/product/domain/entities"
	value_objects "tech_challenge/internal/product/domain/value-objects"
	"tech_challenge/internal/product/interfaces"
	shared_interfaces "tech_challenge/internal/shared/interfaces"
)

type ProductGateway struct {
	dataSource  interfaces.IProductDataSource
	fileService shared_interfaces.IFileProvider
}

func NewProductGateway(dataSource interfaces.IProductDataSource, fileService shared_interfaces.IFileProvider) *ProductGateway {
	return &ProductGateway{
		dataSource:  dataSource,
		fileService: fileService,
	}
}

func (g *ProductGateway) Insert(product entities.Product) error {
	productImages := make([]daos.ProductImageDAO, len(product.Images))
	for i, img := range product.Images {
		productImages[i] = daos.ProductImageDAO{
			ID:        img.ID,
			ProductID: product.ID,
			FileName:  img.FileName,
			Url:       img.Url,
			IsDefault: img.IsDefault,
			CreatedAt: img.CreatedAt,
		}
	}

	return g.dataSource.Insert(daos.ProductDAO{
		ID:          product.ID,
		Name:        product.Name.Value(),
		Description: product.Description,
		Price:       product.Price.Value(),
		CategoryID:  product.CategoryID,
		Images:      productImages,
		Active:      product.Active,
	})
}

func (g *ProductGateway) FindAll() ([]entities.Product, error) {
	productsDAO, err := g.dataSource.FindAll()
	if err != nil {
		return nil, err
	}
	products := make([]entities.Product, len(productsDAO))
	for i, p := range productsDAO {
		productImages := make([]*value_objects.Image, len(p.Images))
		for j, img := range p.Images {
			productImages[j] = &value_objects.Image{
				FileName:  img.FileName,
				Url:       img.Url,
				CreatedAt: img.CreatedAt,
				ID:        img.ID,
				IsDefault: img.IsDefault,
			}
		}
		product, err := entities.NewProductWithImages(
			p.ID,
			p.CategoryID,
			p.Name,
			p.Description,
			p.Price,
			p.Active,
			[]struct{ FileName, Url string }{},
		)
		if err != nil {
			return nil, err
		}
		product.Images = productImages
		products[i] = *product
	}
	return products, nil
}

func (g *ProductGateway) FindAllByCategoryID(categoryID string) ([]entities.Product, error) {
	productsDAO, err := g.dataSource.FindAllByCategoryID(categoryID)
	if err != nil {
		return nil, err
	}
	products := make([]entities.Product, len(productsDAO))
	for i, p := range productsDAO {
		productImages := make([]*value_objects.Image, len(p.Images))
		for j, img := range p.Images {
			productImages[j] = &value_objects.Image{
				FileName:  img.FileName,
				Url:       img.Url,
				CreatedAt: img.CreatedAt,
				ID:        img.ID,
				IsDefault: img.IsDefault,
			}
		}
		product, err := entities.NewProductWithImages(
			p.ID,
			p.CategoryID,
			p.Name,
			p.Description,
			p.Price,
			p.Active,
			[]struct{ FileName, Url string }{},
		)
		if err != nil {
			return nil, err
		}
		product.Images = productImages
		products[i] = *product
	}
	return products, nil
}

func (g *ProductGateway) FindByID(id string) (entities.Product, error) {
	productDAO, err := g.dataSource.FindByID(id)
	if err != nil {
		return entities.Product{}, err
	}
	productImages := make([]*value_objects.Image, 0, len(productDAO.Images))
	if len(productDAO.Images) > 0 {
		img := productDAO.Images[0]
		productImages = append(productImages, &value_objects.Image{
			FileName:  img.FileName,
			Url:       img.Url,
			CreatedAt: img.CreatedAt,
			ID:        img.ID,
			IsDefault: img.IsDefault,
		})
	}
	product, err := entities.NewProductWithImages(
		productDAO.ID,
		productDAO.CategoryID,
		productDAO.Name,
		productDAO.Description,
		productDAO.Price,
		productDAO.Active,
		[]struct{ FileName, Url string }{},
	)
	if err != nil {
		return entities.Product{}, err
	}
	product.Images = productImages
	return *product, nil
}

func (g *ProductGateway) Update(product entities.Product) error {
	productImages := make([]daos.ProductImageDAO, len(product.Images))
	for i, img := range product.Images {
		productImages[i] = daos.ProductImageDAO{
			ID:        img.ID,
			ProductID: product.ID,
			FileName:  img.FileName,
			Url:       img.Url,
			IsDefault: img.IsDefault,
			CreatedAt: img.CreatedAt,
		}
	}

	return g.dataSource.Update(daos.ProductDAO{
		ID:          product.ID,
		Name:        product.Name.Value(),
		Description: product.Description,
		Price:       product.Price.Value(),
		CategoryID:  product.CategoryID,
		Images:      productImages,
		Active:      product.Active,
	})
}

func (g *ProductGateway) Delete(id string) error {
	return g.dataSource.Delete(id)
}

func (g *ProductGateway) UploadImage(fileName string, fileContent []byte) (string, error) {
	err := g.fileService.UploadFile(fileName, fileContent)
	if err != nil {
		return "", err
	}
	return g.GetImageUrl(fileName), nil
}

func (g *ProductGateway) DeleteImage(fileName string) error {
	return g.fileService.DeleteFile(fileName)
}

func (g *ProductGateway) GetImageUrl(fileName string) string {
	url, err := g.fileService.GetPresignedURL(fileName)
	if err != nil {
		// fallback local ou log de erro
		return ""
	}
	return url
}

func (g *ProductGateway) AddProductImage(img daos.ProductImageDAO) error {
	return g.dataSource.AddProductImage(img)
}

func (g *ProductGateway) UpdateProductImage(img daos.ProductImageDAO) error {
	return g.dataSource.UpdateProductImage(img)
}

func (g *ProductGateway) SetProductImageAsDefault(product entities.Product, img *value_objects.Image) error {
	imgDAO := daos.ProductImageDAO{
		ID:        img.ID,
		ProductID: product.ID,
		FileName:  img.FileName,
		Url:       img.Url,
		IsDefault: true,
		CreatedAt: img.CreatedAt,
	}
	// Adiciona a nova imagem como default
	if err := g.dataSource.AddProductImage(imgDAO); err != nil {
		return err
	}
	// Atualiza todas as outras imagens para is_default = false
	return g.dataSource.SetPreviousImagesAsNotDefault(product.ID, img.ID)
}

func (g *ProductGateway) AddAndSetDefaultImage(product entities.Product, url string) error {
	if len(product.Images) == 0 {
		return fmt.Errorf("Produto não possui imagens para atualizar")
	}
	img := product.Images[len(product.Images)-1]
	img.Url = url
	img.IsDefault = true
	imgDAO := daos.ProductImageDAO{
		ID:        img.ID,
		ProductID: product.ID,
		FileName:  img.FileName,
		Url:       img.Url,
		IsDefault: img.IsDefault,
		CreatedAt: img.CreatedAt,
	}
	if err := g.dataSource.AddProductImage(imgDAO); err != nil {
		return err
	}
	return g.dataSource.SetPreviousImagesAsNotDefault(product.ID, img.ID)
}
