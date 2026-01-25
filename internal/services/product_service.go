package services

import (
	"context"

	"github.com/gustionusamba24/kasir-api-go/internal/domain/dtos"
)

type ProductService interface {
	// GetAll retrieves all products
	GetAll(ctx context.Context) ([]dtos.ProductDto, error)

	// GetByID retrieves a product by ID
	GetByID(ctx context.Context, id int) (*dtos.ProductDto, error)

	// GetByCategoryID retrieves all products by category ID
	GetByCategoryID(ctx context.Context, categoryID int) ([]dtos.ProductDto, error)

	// Create creates a new product
	Create(ctx context.Context, dto *dtos.ProductCreateRequestDto) (*dtos.ProductDto, error)

	// Update updates an existing product
	Update(ctx context.Context, id int, dto *dtos.ProductUpdateRequestDto) (*dtos.ProductDto, error)

	// Delete deletes a product by ID
	Delete(ctx context.Context, id int) error
}
