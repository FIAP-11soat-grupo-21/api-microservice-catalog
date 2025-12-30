package entities

import (
	// "fmt"
	"slices"
	"tech_challenge/internal/product/domain/exceptions"
	value_objects "tech_challenge/internal/product/domain/value-objects"
)

type Product struct {
	ID          string
	CategoryID  string
	Name        value_objects.Name
	Description string
	Price       value_objects.Price
	Images      []*value_objects.Image
	Active      bool
}

func NewProduct(id, categoryID, name, description string, price float64, active bool) (*Product, error) {
	productName, err := value_objects.NewName(name)
	if err != nil {
		return nil, err
	}

	productPrice, err := value_objects.NewPrice(price)
	if err != nil {
		return nil, err
	}

	defaultImage, err := value_objects.NewImageDefault()
	if err != nil {
		return nil, err
	}

	defaultImagePtr := &defaultImage

	return &Product{
		ID:          id,
		CategoryID:  categoryID,
		Name:        productName,
		Description: description,
		Price:       productPrice,
		Images:      []*value_objects.Image{defaultImagePtr},
		Active:      active,
	}, nil
}

func NewProductWithImages(
	id,
	categoryID,
	name,
	description string,
	price float64,
	active bool,
	images []struct{ FileName, Url string },
) (*Product, error) {
	productName, err := value_objects.NewName(name)

	if err != nil {
		return nil, err
	}

	productPrice, err := value_objects.NewPrice(price)

	if err != nil {
		return nil, err
	}

	productImages := make([]*value_objects.Image, len(images))

	for i, img := range images {
		image, err := value_objects.NewImageWithFileNameAndUrl(img.FileName, img.Url, false)

		if err != nil {
			return nil, err
		}

		productImages[i] = &image
	}

	return &Product{
		ID:          id,
		CategoryID:  categoryID,
		Name:        productName,
		Description: description,
		Price:       productPrice,
		Images:      productImages,
		Active:      active,
	}, nil
}

func (c *Product) SetName(name string) error {
	newName, err := value_objects.NewName(name)
	if err != nil {
		return err
	}

	c.Name = newName
	return nil
}

func (c *Product) SetPrice(price float64) error {
	newPrice, err := value_objects.NewPrice(price)
	if err != nil {
		return err
	}

	c.Price = newPrice
	return nil
}

func (c *Product) SetDescription(description string) error {
	c.Description = description
	return nil
}

func (p *Product) Activate() error {
	p.Active = true
	return nil
}

func (p *Product) Deactivate() error {
	p.Active = false
	return nil
}

func (c *Product) SetCategory(categoryId string) error {
	c.CategoryID = categoryId
	return nil
}

func (c *Product) AddImage(originalFileName string) (*string, error) {
	img, err := value_objects.NewImage(originalFileName)
	if err != nil {
		return nil, err
	}
	c.Images = append(c.Images, &img)
	return &img.FileName, nil
}

func (c *Product) RemoveImage(fileName string) error {
	isLastImage := len(c.Images) == 1

	if isLastImage {
		return &exceptions.InvalidProductImageException{
			Message: "Product image cannot be empty, at least one image is required",
		}
	}

	for i, img := range c.Images {
		if img.FileName == fileName {
			c.Images = slices.Delete(c.Images, i, i+1)
			return nil
		}
	}

	return &exceptions.ImageNotFoundException{}
}

func (c *Product) ImageIsDefault(imageFileName string) bool {
	for _, img := range c.Images {
		if img.FileName == imageFileName && img.IsDefault {
			return true
		}
	}
	return false
}

func (c *Product) IsEmpty() bool {
	return c.ID == ""
}

func (p *Product) SetPreviousImagesAsNotDefault() {
	for i := 0; i < len(p.Images)-1; i++ {
		p.Images[i].IsDefault = false
	}
}
