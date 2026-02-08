package repositories

import (
	"context"

	"github.com/gustionusamba24/kasir-api-go/internal/domain/entities"
)

type TransactionRepository interface {
	// Create creates a new transaction with details
	Create(ctx context.Context, transaction *entities.Transaction) error
	
	// FindByID retrieves a transaction by ID with its details
	FindByID(ctx context.Context, id int) (*entities.Transaction, error)
	
	// FindAll retrieves all transactions with their details
	FindAll(ctx context.Context) ([]entities.Transaction, error)
	
	// CreateDetail creates a transaction detail
	CreateDetail(ctx context.Context, detail *entities.TransactionDetail) error
}
