package gateways

import (
	"tech_challenge/internal/product/daos"
	"tech_challenge/internal/product/domain/entities"
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
			FileName: img.FileName,
			Url:      img.Url,
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
		// Seleciona imagem principal: default ou mais recente
		var mainImage struct{ FileName, Url string }
		if len(p.Images) > 0 {
			var foundDefault bool
			var latestImage daos.ProductImageDAO
			for _, img := range p.Images {
				if img.IsDefault {
					mainImage = struct{ FileName, Url string }{img.FileName, img.Url}
					foundDefault = true
					break
				}
				if latestImage.CreatedAt.IsZero() || img.CreatedAt.After(latestImage.CreatedAt) {
					latestImage = img
				}
			}
			if !foundDefault {
				mainImage = struct{ FileName, Url string }{latestImage.FileName, latestImage.Url}
			}
		}

		product, err := entities.NewProductWithImages(
			p.ID,
			p.CategoryID,
			p.Name,
			p.Description,
			p.Price,
			p.Active,
			[]struct{ FileName, Url string }{mainImage},
		)

		if err != nil {
			return nil, err
		}

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
		// Seleciona imagem principal: default ou mais recente
		var mainImage struct{ FileName, Url string }
		if len(p.Images) > 0 {
			var foundDefault bool
			var latestImage daos.ProductImageDAO
			for _, img := range p.Images {
				if img.IsDefault {
					mainImage = struct{ FileName, Url string }{img.FileName, img.Url}
					foundDefault = true
					break
				}
				if latestImage.CreatedAt.IsZero() || img.CreatedAt.After(latestImage.CreatedAt) {
					latestImage = img
				}
			}
			if !foundDefault {
				mainImage = struct{ FileName, Url string }{latestImage.FileName, latestImage.Url}
			}
		}

		product, err := entities.NewProductWithImages(
			p.ID,
			p.CategoryID,
			p.Name,
			p.Description,
			p.Price,
			p.Active,
			[]struct{ FileName, Url string }{mainImage},
		)

		if err != nil {
			return nil, err
		}

		products[i] = *product
	}

	return products, nil
}

func (g *ProductGateway) FindByID(id string) (entities.Product, error) {
	productDAO, err := g.dataSource.FindByID(id)

	if err != nil {
		return entities.Product{}, err
	}

	// Seleciona imagem principal: default ou mais recente
	var mainImage struct{ FileName, Url string }
	if len(productDAO.Images) > 0 {
		var foundDefault bool
		var latestImage daos.ProductImageDAO
		for _, img := range productDAO.Images {
			if img.IsDefault {
				mainImage = struct{ FileName, Url string }{img.FileName, img.Url}
				foundDefault = true
				break
			}
			if latestImage.CreatedAt.IsZero() || img.CreatedAt.After(latestImage.CreatedAt) {
				latestImage = img
			}
		}
		if !foundDefault {
			mainImage = struct{ FileName, Url string }{latestImage.FileName, latestImage.Url}
		}
	}

	product, err := entities.NewProductWithImages(
		productDAO.ID,
		productDAO.CategoryID,
		productDAO.Name,
		productDAO.Description,
		productDAO.Price,
		productDAO.Active,
		[]struct{ FileName, Url string }{mainImage},
	)

	if err != nil {
		return entities.Product{}, err
	}

	return *product, nil
}

func (g *ProductGateway) Update(product entities.Product) error {
	productImages := make([]daos.ProductImageDAO, len(product.Images))

	for i, img := range product.Images {
		productImages[i] = daos.ProductImageDAO{
			FileName: img.FileName,
			Url:      img.Url,
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

func (g *ProductGateway) UploadImage(fileName string, fileContent []byte) error {
	return g.fileService.UploadFile(fileName, fileContent)
}

func (g *ProductGateway) DeleteImage(fileName string) error {
	return g.fileService.DeleteFile(fileName)
}
