package mappers

import (
	"github.com/gustionusamba24/kasir-api-go/internal/domain/dtos"
	"github.com/gustionusamba24/kasir-api-go/internal/domain/entities"
)

type TransactionMapper struct{}

// ToDto converts Transaction entity to TransactionDto
func (m *TransactionMapper) ToDto(transaction *entities.Transaction) *dtos.TransactionDto {
	if transaction == nil {
		return nil
	}

	dto := &dtos.TransactionDto{
		ID:          transaction.ID,
		TotalAmount: transaction.TotalAmount,
		CreatedAt:   transaction.CreatedAt,
	}

	// Map details
	if transaction.Details != nil {
		dto.Details = make([]dtos.TransactionDetailDto, len(transaction.Details))
		for i, detail := range transaction.Details {
			dto.Details[i] = dtos.TransactionDetailDto{
				ID:            detail.ID,
				TransactionID: detail.TransactionID,
				ProductID:     detail.ProductID,
				ProductName:   detail.ProductName,
				Quantity:      detail.Quantity,
				Subtotal:      detail.Subtotal,
			}
		}
	}

	return dto
}

// ToDtoList converts slice of Transaction entities to slice of TransactionDto
func (m *TransactionMapper) ToDtoList(transactions []entities.Transaction) []dtos.TransactionDto {
	if transactions == nil {
		return nil
	}

	result := make([]dtos.TransactionDto, len(transactions))
	for i, transaction := range transactions {
		dto := m.ToDto(&transaction)
		if dto != nil {
			result[i] = *dto
		}
	}
	return result
}
