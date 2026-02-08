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
	
	// GetTodayRevenue returns the total revenue from today's transactions
	GetTodayRevenue(ctx context.Context) (int, error)
	
	// GetTodayTransactionCount returns the count of transactions made today
	GetTodayTransactionCount(ctx context.Context) (int, error)
	
	// GetTodayBestSellingProduct returns the product name and quantity sold for today's best selling product
	GetTodayBestSellingProduct(ctx context.Context) (productName string, qtySold int, err error)
	
	// GetDateRangeRevenue returns the total revenue from transactions within a date range
	GetDateRangeRevenue(ctx context.Context, startDate, endDate string) (int, error)
	
	// GetDateRangeTransactionCount returns the count of transactions within a date range
	GetDateRangeTransactionCount(ctx context.Context, startDate, endDate string) (int, error)
	
	// GetDateRangeBestSellingProduct returns the product name and quantity sold for best selling product within a date range
	GetDateRangeBestSellingProduct(ctx context.Context, startDate, endDate string) (productName string, qtySold int, err error)
}
