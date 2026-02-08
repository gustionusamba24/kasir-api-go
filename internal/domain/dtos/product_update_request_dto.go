package dtos

type ProductUpdateRequestDto struct {
	Name       string  `json:"name" validate:"required,min=3,max=100"`
	Price      float64 `json:"price" validate:"required,gt=0"`
	Stock      int     `json:"stock" validate:"required,gte=0"`
	Active     *bool   `json:"active" validate:"omitempty"`
	CategoryID *int    `json:"category_id" validate:"omitempty,gt=0"`
}
