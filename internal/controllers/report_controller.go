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
