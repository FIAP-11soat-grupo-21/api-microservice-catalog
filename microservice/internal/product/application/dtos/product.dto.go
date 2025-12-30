package dtos

type CreateProductDTO struct {
	CategoryID  string
	Name        string
	Description string
	Price       float64
	Active      bool
}

type UpdateProductDTO struct {
	ID          string
	Name        string
	Description string
	Price       float64
	Active      bool
	CategoryID  string
}

type UploadProductImageDTO struct {
	ProductID   string
	FileName    string
	FileContent []byte
}

type ProductImageDTO struct {
	ID        string
	FileName  string
	Url       string
	IsDefault bool
}

type ProductResultDTO struct {
	ID          string
	Name        string
	Description string
	Price       float64
	Active      bool
	CategoryID  string
	Images      []ProductImageDTO
}
