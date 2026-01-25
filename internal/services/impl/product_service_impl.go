package impl

import (
	"context"
	"fmt"

	"github.com/gustionusamba24/kasir-api-go/internal/domain/dtos"
	"github.com/gustionusamba24/kasir-api-go/internal/mappers"
	"github.com/gustionusamba24/kasir-api-go/internal/repositories"
	"github.com/gustionusamba24/kasir-api-go/internal/services"
)

type productServiceImpl struct {
	repository         repositories.ProductRepository
	categoryRepository repositories.CategoryRepository
	mapper             *mappers.ProductMapper
}

// NewProductService creates a new instance of ProductService
func NewProductService(
	repository repositories.ProductRepository,
	categoryRepository repositories.CategoryRepository,
) services.ProductService {
	return &productServiceImpl{
		repository:         repository,
		categoryRepository: categoryRepository,
		mapper:             &mappers.ProductMapper{},
	}
}

// GetAll retrieves all products
func (s *productServiceImpl) GetAll(ctx context.Context) ([]dtos.ProductDto, error) {
	products, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all products: %w", err)
	}

	return s.mapper.ToDtoList(products), nil
}

// GetByID retrieves a product by ID
func (s *productServiceImpl) GetByID(ctx context.Context, id int) (*dtos.ProductDto, error) {
	product, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get product by id %d: %w", id, err)
	}

	if product == nil {
		return nil, fmt.Errorf("product with id %d not found", id)
	}

	return s.mapper.ToDto(product), nil
}

// GetByCategoryID retrieves all products by category ID
func (s *productServiceImpl) GetByCategoryID(ctx context.Context, categoryID int) ([]dtos.ProductDto, error) {
	// Validate category exists
	category, err := s.categoryRepository.FindByID(ctx, categoryID)
	if err != nil {
		return nil, fmt.Errorf("failed to find category by id %d: %w", categoryID, err)
	}

	if category == nil {
		return nil, fmt.Errorf("category with id %d not found", categoryID)
	}

	products, err := s.repository.FindByCategoryID(ctx, categoryID)
	if err != nil {
		return nil, fmt.Errorf("failed to get products by category id %d: %w", categoryID, err)
	}

	return s.mapper.ToDtoList(products), nil
}

// Create creates a new product
func (s *productServiceImpl) Create(ctx context.Context, dto *dtos.ProductCreateRequestDto) (*dtos.ProductDto, error) {
	if dto == nil {
		return nil, fmt.Errorf("create request dto cannot be nil")
	}

	// Validate category exists if provided
	if dto.CategoryID != nil {
		category, err := s.categoryRepository.FindByID(ctx, *dto.CategoryID)
		if err != nil {
			return nil, fmt.Errorf("failed to find category by id %d: %w", *dto.CategoryID, err)
		}
		if category == nil {
			return nil, fmt.Errorf("category with id %d not found", *dto.CategoryID)
		}
	}

	// Convert DTO to request
	request := s.mapper.ToCreateRequest(dto)

	// Convert request to entity
	product := s.mapper.ToEntity(request)

	// Save to repository
	err := s.repository.Create(ctx, product)
	if err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	// Return created product as DTO
	return s.mapper.ToDto(product), nil
}

// Update updates an existing product
func (s *productServiceImpl) Update(ctx context.Context, id int, dto *dtos.ProductUpdateRequestDto) (*dtos.ProductDto, error) {
	if dto == nil {
		return nil, fmt.Errorf("update request dto cannot be nil")
	}

	// Check if product exists
	existingProduct, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find product by id %d: %w", id, err)
	}

	if existingProduct == nil {
		return nil, fmt.Errorf("product with id %d not found", id)
	}

	// Validate category exists if provided
	if dto.CategoryID != nil {
		category, err := s.categoryRepository.FindByID(ctx, *dto.CategoryID)
		if err != nil {
			return nil, fmt.Errorf("failed to find category by id %d: %w", *dto.CategoryID, err)
		}
		if category == nil {
			return nil, fmt.Errorf("category with id %d not found", *dto.CategoryID)
		}
	}

	// Convert DTO to request
	request := s.mapper.ToUpdateRequest(dto)

	// Update entity with request data
	s.mapper.UpdateEntity(existingProduct, request)
	existingProduct.ID = id // Ensure ID is preserved

	// Save updated entity
	err = s.repository.Update(ctx, existingProduct)
	if err != nil {
		return nil, fmt.Errorf("failed to update product: %w", err)
	}

	// Return updated product as DTO
	return s.mapper.ToDto(existingProduct), nil
}

// Delete deletes a product by ID
func (s *productServiceImpl) Delete(ctx context.Context, id int) error {
	// Check if product exists
	existingProduct, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to find product by id %d: %w", id, err)
	}

	if existingProduct == nil {
		return fmt.Errorf("product with id %d not found", id)
	}

	// Delete product
	err = s.repository.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	return nil
}
