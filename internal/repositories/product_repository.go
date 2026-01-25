package repositories

import (
	"context"

	"github.com/gustionusamba24/kasir-api-go/internal/domain/entities"
)

type ProductRepository interface {
	FindAll(ctx context.Context) ([]entities.Product, error)
	FindByID(ctx context.Context, id int) (*entities.Product, error)
	FindByCategoryID(ctx context.Context, categoryID int) ([]entities.Product, error)
	Create(ctx context.Context, product *entities.Product) error
	Update(ctx context.Context, product *entities.Product) error
	Delete(ctx context.Context, id int) error
}