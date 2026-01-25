package repositories

import (
	"context"

	"github.com/gustionusamba24/kasir-api-go/internal/domain/entities"
)

type CategoryRepository interface {
	FindAll(ctx context.Context) ([]entities.Category, error)
	FindByID(ctx context.Context, id int) (*entities.Category, error)
	Create(ctx context.Context, category *entities.Category) error
	Update(ctx context.Context, category *entities.Category) error
	Delete(ctx context.Context, id int) error
}
