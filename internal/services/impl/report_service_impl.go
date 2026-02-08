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

func (s *reportServiceImpl) GetDateRangeReport(ctx context.Context, startDate, endDate string) (*dtos.DateRangeReportDto, error) {
	// Get date range total revenue
	totalRevenue, err := s.transactionRepository.GetDateRangeRevenue(ctx, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to get date range revenue: %w", err)
	}

	// Get date range transaction count
	totalTransactions, err := s.transactionRepository.GetDateRangeTransactionCount(ctx, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to get date range transaction count: %w", err)
	}

	// Get date range best selling product
	productName, qtySold, err := s.transactionRepository.GetDateRangeBestSellingProduct(ctx, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to get date range best selling product: %w", err)
	}

	// Build report DTO
	report := &dtos.DateRangeReportDto{
		StartDate:         startDate,
		EndDate:           endDate,
		TotalRevenue:      totalRevenue,
		TotalTransactions: totalTransactions,
	}

	// Only add best selling product if there are transactions in the date range
	if productName != "" {
		report.BestSellingProduct = &dtos.BestSellingProductDto{
			Name:    productName,
			QtySold: qtySold,
		}
	}

	return report, nil
}
