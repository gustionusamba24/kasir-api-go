package impl

import (
	"context"
	"fmt"

	"github.com/gustionusamba24/kasir-api-go/internal/domain/dtos"
	"github.com/gustionusamba24/kasir-api-go/internal/mappers"
	"github.com/gustionusamba24/kasir-api-go/internal/repositories"
	"github.com/gustionusamba24/kasir-api-go/internal/services"
)

type categoryServiceImpl struct {
	repository repositories.CategoryRepository
	mapper     *mappers.CategoryMapper
}

// NewCategoryService creates a new instance of CategoryService
func NewCategoryService(repository repositories.CategoryRepository) services.CategoryService {
	return &categoryServiceImpl{
		repository: repository,
		mapper:     &mappers.CategoryMapper{},
	}
}

// GetAll retrieves all categories
func (s *categoryServiceImpl) GetAll(ctx context.Context) ([]dtos.CategoryDto, error) {
	categories, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all categories: %w", err)
	}

	return s.mapper.ToDtoList(categories), nil
}

// GetByID retrieves a category by ID
func (s *categoryServiceImpl) GetByID(ctx context.Context, id int) (*dtos.CategoryDto, error) {
	category, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get category by id %d: %w", id, err)
	}

	if category == nil {
		return nil, fmt.Errorf("category with id %d not found", id)
	}

	return s.mapper.ToDto(category), nil
}

// Create creates a new category
func (s *categoryServiceImpl) Create(ctx context.Context, dto *dtos.CategoryCreateRequestDto) (*dtos.CategoryDto, error) {
	if dto == nil {
		return nil, fmt.Errorf("create request dto cannot be nil")
	}

	// Convert DTO to request
	request := s.mapper.ToCreateRequest(dto)

	// Convert request to entity
	category := s.mapper.ToEntity(request)

	// Save to repository
	err := s.repository.Create(ctx, category)
	if err != nil {
		return nil, fmt.Errorf("failed to create category: %w", err)
	}

	// Return created category as DTO
	return s.mapper.ToDto(category), nil
}

// Update updates an existing category
func (s *categoryServiceImpl) Update(ctx context.Context, id int, dto *dtos.CategoryUpdateRequestDto) (*dtos.CategoryDto, error) {
	if dto == nil {
		return nil, fmt.Errorf("update request dto cannot be nil")
	}

	// Check if category exists
	existingCategory, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find category by id %d: %w", id, err)
	}

	if existingCategory == nil {
		return nil, fmt.Errorf("category with id %d not found", id)
	}

	// Convert DTO to request
	request := s.mapper.ToUpdateRequest(dto)

	// Update entity with request data
	s.mapper.UpdateEntity(existingCategory, request)
	existingCategory.ID = id // Ensure ID is preserved

	// Save updated entity
	err = s.repository.Update(ctx, existingCategory)
	if err != nil {
		return nil, fmt.Errorf("failed to update category: %w", err)
	}

	// Return updated category as DTO
	return s.mapper.ToDto(existingCategory), nil
}

// Delete deletes a category by ID
func (s *categoryServiceImpl) Delete(ctx context.Context, id int) error {
	// Check if category exists
	existingCategory, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to find category by id %d: %w", id, err)
	}

	if existingCategory == nil {
		return fmt.Errorf("category with id %d not found", id)
	}

	// Delete category
	err = s.repository.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete category: %w", err)
	}

	return nil
}
