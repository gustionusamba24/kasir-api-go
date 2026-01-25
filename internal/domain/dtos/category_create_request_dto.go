package dtos

type CategoryCreateRequestDto struct {
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"max=500"`
}
