package dtos

type ProductCreateRequest struct {
	Name       string
	Price      float64
	Stock      int
	Active     bool
	CategoryID *int
}
