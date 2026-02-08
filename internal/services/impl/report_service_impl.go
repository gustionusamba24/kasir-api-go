package impl

import (
	"context"
	"fmt"

	"github.com/gustionusamba24/kasir-api-go/internal/domain/dtos"
	"github.com/gustionusamba24/kasir-api-go/internal/repositories"
	"github.com/gustionusamba24/kasir-api-go/internal/services"
)

type reportServiceImpl struct {
	transactionRepository repositories.TransactionRepository
}

func NewReportService(transactionRepository repositories.TransactionRepository) services.ReportService {
	return &reportServiceImpl{
		transactionRepository: transactionRepository,
	}
}

func (s *reportServiceImpl) GetTodayReport(ctx context.Context) (*dtos.TodayReportDto, error) {
	// Get today's total revenue
	totalRevenue, err := s.transactionRepository.GetTodayRevenue(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get today's revenue: %w", err)
	}

	// Get today's transaction count
	totalTransactions, err := s.transactionRepository.GetTodayTransactionCount(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get today's transaction count: %w", err)
	}

	// Get today's best selling product
	productName, qtySold, err := s.transactionRepository.GetTodayBestSellingProduct(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get today's best selling product: %w", err)
	}

	// Build report DTO
	report := &dtos.TodayReportDto{
		TotalRevenue:      totalRevenue,
		TotalTransactions: totalTransactions,
	}

	// Only add best selling product if there are transactions today
	if productName != "" {
		report.BestSellingProduct = &dtos.BestSellingProductDto{
			Name:    productName,
			QtySold: qtySold,
		}
	}

	return report, nil
}
