package controllers

import (
	"net/http"

	"github.com/gustionusamba24/kasir-api-go/internal/services"
)

type ReportController struct {
	service services.ReportService
}

// NewReportController creates a new instance of ReportController
func NewReportController(service services.ReportService) *ReportController {
	return &ReportController{
		service: service,
	}
}

// GetTodayReport godoc
// @Summary      Get today's transaction report
// @Description  Retrieve today's transaction report including total revenue, transaction count, and best selling product
// @Tags         reports
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "success response with today's report data"
// @Failure      500  {object}  map[string]interface{}  "internal server error"
// @Router       /report/today [get]
func (c *ReportController) GetTodayReport(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	report, err := c.service.GetTodayReport(ctx)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    report,
	})
}

// GetDateRangeReport godoc
// @Summary      Get transaction report for date range
// @Description  Retrieve transaction report for a given date range including total revenue, transaction count, and best selling product
// @Tags         reports
// @Accept       json
// @Produce      json
// @Param        start_date  query  string  true  "Start date in YYYY-MM-DD format"
// @Param        end_date    query  string  true  "End date in YYYY-MM-DD format"
// @Success      200  {object}  map[string]interface{}  "success response with date range report data"
// @Failure      400  {object}  map[string]interface{}  "bad request - missing or invalid parameters"
// @Failure      500  {object}  map[string]interface{}  "internal server error"
// @Router       /report [get]
func (c *ReportController) GetDateRangeReport(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Get query parameters
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	// Validate required parameters
	if startDate == "" || endDate == "" {
		respondWithError(w, http.StatusBadRequest, "start_date and end_date are required")
		return
	}

	report, err := c.service.GetDateRangeReport(ctx, startDate, endDate)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    report,
	})
}
