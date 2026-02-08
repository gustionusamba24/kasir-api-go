package impl

import (
	"context"
	"fmt"

	"github.com/gustionusamba24/kasir-api-go/internal/domain/dtos"
	"github.com/gustionusamba24/kasir-api-go/internal/domain/entities"
	"github.com/gustionusamba24/kasir-api-go/internal/mappers"
	"github.com/gustionusamba24/kasir-api-go/internal/repositories"
	"github.com/gustionusamba24/kasir-api-go/internal/services"
)

type transactionServiceImpl struct {
	transactionRepository repositories.TransactionRepository
	productRepository     repositories.ProductRepository
	mapper                *mappers.TransactionMapper
}

func NewTransactionService(
	transactionRepository repositories.TransactionRepository,
	productRepository repositories.ProductRepository,
) services.TransactionService {
	return &transactionServiceImpl{
		transactionRepository: transactionRepository,
		productRepository:     productRepository,
		mapper:                &mappers.TransactionMapper{},
	}
}

func (s *transactionServiceImpl) Checkout(ctx context.Context, dto *dtos.TransactionCreateRequestDto) (*dtos.TransactionDto, error) {
	if dto == nil {
		return nil, fmt.Errorf("checkout request cannot be nil")
	}

	if len(dto.Items) == 0 {
		return nil, fmt.Errorf("checkout items cannot be empty")
	}

	// Build transaction with details
	var transaction entities.Transaction
	var details []entities.TransactionDetail
	totalAmount := 0

	for _, item := range dto.Items {
		// Get product to validate and calculate subtotal
		product, err := s.productRepository.FindByID(ctx, item.ProductID)
		if err != nil {
			return nil, fmt.Errorf("failed to find product with id %d: %w", item.ProductID, err)
		}
		if product == nil {
			return nil, fmt.Errorf("product with id %d not found", item.ProductID)
		}

		// Check stock availability
		if product.Stock < item.Quantity {
			return nil, fmt.Errorf("insufficient stock for product %s (available: %d, requested: %d)", 
				product.Name, product.Stock, item.Quantity)
		}

		// Check if product is active
		if !product.Active {
			return nil, fmt.Errorf("product %s is not active", product.Name)
		}

		// Calculate subtotal (price * quantity)
		subtotal := int(product.Price * float64(item.Quantity))
		totalAmount += subtotal

		// Create transaction detail
		detail := entities.TransactionDetail{
			ProductID:   item.ProductID,
			ProductName: product.Name,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		}
		details = append(details, detail)

		// Update product stock
		product.Stock -= item.Quantity
		err = s.productRepository.Update(ctx, product)
		if err != nil {
			return nil, fmt.Errorf("failed to update product stock: %w", err)
		}
	}

	transaction.TotalAmount = totalAmount
	transaction.Details = details

	// Create transaction with details in database
	err := s.transactionRepository.Create(ctx, &transaction)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %w", err)
	}

	// Return transaction DTO
	return s.mapper.ToDto(&transaction), nil
}

func (s *transactionServiceImpl) GetByID(ctx context.Context, id int) (*dtos.TransactionDto, error) {
	transaction, err := s.transactionRepository.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction by id %d: %w", id, err)
	}

	if transaction == nil {
		return nil, fmt.Errorf("transaction with id %d not found", id)
	}

	return s.mapper.ToDto(transaction), nil
}

func (s *transactionServiceImpl) GetAll(ctx context.Context) ([]dtos.TransactionDto, error) {
	transactions, err := s.transactionRepository.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all transactions: %w", err)
	}

	return s.mapper.ToDtoList(transactions), nil
}
