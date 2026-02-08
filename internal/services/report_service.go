package services

import (
	"context"

	"github.com/gustionusamba24/kasir-api-go/internal/domain/dtos"
)

type ReportService interface {
	// GetTodayReport retrieves today's transaction report
	GetTodayReport(ctx context.Context) (*dtos.TodayReportDto, error)

	// GetDateRangeReport retrieves transaction report for a specific date range
	GetDateRangeReport(ctx context.Context, startDate, endDate string) (*dtos.DateRangeReportDto, error)
}
