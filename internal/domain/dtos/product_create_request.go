package dtos

type ProductCreateRequest struct {
	Name       string
	Price      float64
	Stock      int
	CategoryID *int
}
