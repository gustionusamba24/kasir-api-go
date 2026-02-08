package dtos

type TransactionCreateRequestDto struct {
	Items []CheckoutItemDto `json:"items" validate:"required,min=1,dive"`
}

type CheckoutItemDto struct {
	ProductID int `json:"product_id" validate:"required,gt=0"`
	Quantity  int `json:"quantity" validate:"required,gt=0"`
}
