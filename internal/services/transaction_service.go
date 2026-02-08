package services

import (
	"context"

	"github.com/gustionusamba24/kasir-api-go/internal/domain/dtos"
)

type TransactionService interface {
	// Checkout creates a new transaction from checkout request
	Checkout(ctx context.Context, dto *dtos.TransactionCreateRequestDto) (*dtos.TransactionDto, error)

	// GetByID retrieves a transaction by ID
	GetByID(ctx context.Context, id int) (*dtos.TransactionDto, error)

	// GetAll retrieves all transactions
	GetAll(ctx context.Context) ([]dtos.TransactionDto, error)
}
