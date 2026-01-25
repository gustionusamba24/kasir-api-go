package mappers

import (
	"time"

	"github.com/gustionusamba24/kasir-api-go/internal/domain/dtos"
	"github.com/gustionusamba24/kasir-api-go/internal/domain/entities"
)

// CategoryMapper handles mapping between Category entity and DTOs
type CategoryMapper struct{}

// ToDto converts Category entity to CategoryDto
// @Mapping(target = "id", source = "id")
// @Mapping(target = "name", source = "name")
// @Mapping(target = "description", source = "description")
// @Mapping(target = "createdAt", source = "createdAt")
// @Mapping(target = "updatedAt", source = "updatedAt")

func (m *CategoryMapper) ToDto(category *entities.Category) *dtos.CategoryDto {
	if category == nil {
		return nil
	}

	return &dtos.CategoryDto{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}
}

// ToDtoList converts slice of Category entities to slice of CategoryDto
func (m *CategoryMapper) ToDtoList(categories []entities.Category) []dtos.CategoryDto {
	if categories == nil {
		return nil
	}

	dtos := make([]dtos.CategoryDto, len(categories))
	for i, category := range categories {
		dto := m.ToDto(&category)
		if dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}

// ToCreateRequest converts CategoryCreateRequestDto to CategoryCreateRequest
// @Mapping(target = "name", source = "name")
// @Mapping(target = "description", source = "description")
func (m *CategoryMapper) ToCreateRequest(dto *dtos.CategoryCreateRequestDto) *dtos.CategoryCreateRequest {
	if dto == nil {
		return nil
	}

	return &dtos.CategoryCreateRequest{
		Name:        dto.Name,
		Description: dto.Description,
	}
}

// ToUpdateRequest converts CategoryUpdateRequestDto to CategoryUpdateRequest
// @Mapping(target = "name", source = "name")
// @Mapping(target = "description", source = "description")
func (m *CategoryMapper) ToUpdateRequest(dto *dtos.CategoryUpdateRequestDto) *dtos.CategoryUpdateRequest {
	if dto == nil {
		return nil
	}

	return &dtos.CategoryUpdateRequest{
		Name:        dto.Name,
		Description: dto.Description,
	}
}

// ToEntity converts CategoryCreateRequest to Category entity
func (m *CategoryMapper) ToEntity(request *dtos.CategoryCreateRequest) *entities.Category {
	if request == nil {
		return nil
	}

	now := time.Now()
	return &entities.Category{
		Name:        request.Name,
		Description: request.Description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// UpdateEntity updates existing Category entity with CategoryUpdateRequest
func (m *CategoryMapper) UpdateEntity(category *entities.Category, request *dtos.CategoryUpdateRequest) {
	if category == nil || request == nil {
		return
	}

	category.Name = request.Name
	category.Description = request.Description
	category.UpdatedAt = time.Now()
}
