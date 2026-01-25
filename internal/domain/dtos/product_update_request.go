package dtos

type ProductUpdateRequest struct {
	Name       string
	Price      float64
	Stock      int
	CategoryID *int
}
