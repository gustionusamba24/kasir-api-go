package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gustionusamba24/kasir-api-go/internal/domain/dtos"
	"github.com/gustionusamba24/kasir-api-go/internal/services"
)

type TransactionController struct {
	service services.TransactionService
}

// NewTransactionController creates a new instance of TransactionController
func NewTransactionController(service services.TransactionService) *TransactionController {
	return &TransactionController{
		service: service,
	}
}

// Checkout godoc
// @Summary      Create a new transaction (checkout)
// @Description  Create a new transaction with multiple products
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param        request  body      dtos.TransactionCreateRequestDto  true  "Checkout request with items"
// @Success      201      {object}  map[string]interface{}  "success response with transaction data"
// @Failure      400      {object}  map[string]interface{}  "invalid request"
// @Failure      404      {object}  map[string]interface{}  "product not found"
// @Failure      500      {object}  map[string]interface{}  "internal server error"
// @Router       /transactions/checkout [post]
func (c *TransactionController) Checkout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var dto dtos.TransactionCreateRequestDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	transaction, err := c.service.Checkout(ctx, &dto)
	if err != nil {
		if isNotFoundError(err) {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]interface{}{
		"success": true,
		"message": "Transaction created successfully",
		"data":    transaction,
	})
}

// GetAll godoc
// @Summary      Get all transactions
// @Description  Retrieve a list of all transactions with their details
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "success response with transactions data"
// @Failure      500  {object}  map[string]interface{}  "internal server error"
// @Router       /transactions [get]
func (c *TransactionController) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	transactions, err := c.service.GetAll(ctx)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    transactions,
	})
}

// GetByID godoc
// @Summary      Get a transaction by ID
// @Description  Retrieve a single transaction by its ID with details
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Transaction ID"
// @Success      200  {object}  map[string]interface{}  "success response with transaction data"
// @Failure      400  {object}  map[string]interface{}  "invalid transaction ID"
// @Failure      404  {object}  map[string]interface{}  "transaction not found"
// @Failure      500  {object}  map[string]interface{}  "internal server error"
// @Router       /transactions/{id} [get]
func (c *TransactionController) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Extract ID from URL path
	id, err := extractIDFromPath(r, "/transactions/")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid transaction ID")
		return
	}

	transaction, err := c.service.GetByID(ctx, id)
	if err != nil {
		if isNotFoundError(err) {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    transaction,
	})
}
