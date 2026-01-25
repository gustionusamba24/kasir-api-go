package services

import (
	"context"

	"github.com/gustionusamba24/kasir-api-go/internal/domain/dtos"
)

type CategoryService interface {
	// GetAll retrieves all categories
	GetAll(ctx context.Context) ([]dtos.CategoryDto, error)

	// GetByID retrieves a category by ID
	GetByID(ctx context.Context, id int) (*dtos.CategoryDto, error)

	// Create creates a new category
	Create(ctx context.Context, dto *dtos.CategoryCreateRequestDto) (*dtos.CategoryDto, error)

	// Update updates an existing category
	Update(ctx context.Context, id int, dto *dtos.CategoryUpdateRequestDto) (*dtos.CategoryDto, error)

	// Delete deletes a category by ID
	Delete(ctx context.Context, id int) error
}
