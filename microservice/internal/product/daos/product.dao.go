package daos

type ProductDAO struct {
	ID          string
	CategoryID  string
	Name        string
	Description string
	Price       float64
	Images      []ProductImageDAO
	Active      bool
}
