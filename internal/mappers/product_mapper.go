package mappers

import (
	"time"

	"github.com/gustionusamba24/kasir-api-go/internal/domain/dtos"
	"github.com/gustionusamba24/kasir-api-go/internal/domain/entities"
)

// ProductMapper handles mapping between Product entity and DTOs
type ProductMapper struct{}

// ToDto converts Product entity to ProductDto
// @Mapping(target = "id", source = "id")
// @Mapping(target = "name", source = "name")
// @Mapping(target = "price", source = "price")
// @Mapping(target = "stock", source = "stock")
// @Mapping(target = "categoryId", source = "categoryId")
// @Mapping(target = "createdAt", source = "createdAt")
// @Mapping(target = "updatedAt", source = "updatedAt")
func (m *ProductMapper) ToDto(product *entities.Product) *dtos.ProductDto {
	if product == nil {
		return nil
	}

	return &dtos.ProductDto{
		ID:         product.ID,
		Name:       product.Name,
		Price:      product.Price,
		Stock:      product.Stock,
		CategoryID: product.CategoryID,
		CreatedAt:  product.CreatedAt,
		UpdatedAt:  product.UpdatedAt,
	}
}

// ToDtoList converts slice of Product entities to slice of ProductDto
func (m *ProductMapper) ToDtoList(products []entities.Product) []dtos.ProductDto {
	if products == nil {
		return nil
	}

	dtos := make([]dtos.ProductDto, len(products))
	for i, product := range products {
		dto := m.ToDto(&product)
		if dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}

// ToCreateRequest converts ProductCreateRequestDto to ProductCreateRequest
// @Mapping(target = "name", source = "name")
// @Mapping(target = "price", source = "price")
// @Mapping(target = "stock", source = "stock")
// @Mapping(target = "categoryId", source = "categoryId")
func (m *ProductMapper) ToCreateRequest(dto *dtos.ProductCreateRequestDto) *dtos.ProductCreateRequest {
	if dto == nil {
		return nil
	}

	return &dtos.ProductCreateRequest{
		Name:       dto.Name,
		Price:      dto.Price,
		Stock:      dto.Stock,
		CategoryID: dto.CategoryID,
	}
}

// ToUpdateRequest converts ProductUpdateRequestDto to ProductUpdateRequest
// @Mapping(target = "name", source = "name")
// @Mapping(target = "price", source = "price")
// @Mapping(target = "stock", source = "stock")
// @Mapping(target = "categoryId", source = "categoryId")
func (m *ProductMapper) ToUpdateRequest(dto *dtos.ProductUpdateRequestDto) *dtos.ProductUpdateRequest {
	if dto == nil {
		return nil
	}

	return &dtos.ProductUpdateRequest{
		Name:       dto.Name,
		Price:      dto.Price,
		Stock:      dto.Stock,
		CategoryID: dto.CategoryID,
	}
}

// ToEntity converts ProductCreateRequest to Product entity
func (m *ProductMapper) ToEntity(request *dtos.ProductCreateRequest) *entities.Product {
	if request == nil {
		return nil
	}

	now := time.Now()
	return &entities.Product{
		Name:       request.Name,
		Price:      request.Price,
		Stock:      request.Stock,
		CategoryID: request.CategoryID,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}

// UpdateEntity updates existing Product entity with ProductUpdateRequest
func (m *ProductMapper) UpdateEntity(product *entities.Product, request *dtos.ProductUpdateRequest) {
	if product == nil || request == nil {
		return
	}

	product.Name = request.Name
	product.Price = request.Price
	product.Stock = request.Stock
	product.CategoryID = request.CategoryID
	product.UpdatedAt = time.Now()
}
