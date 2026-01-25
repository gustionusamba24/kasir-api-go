package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gustionusamba24/kasir-api-go/internal/domain/dtos"
	"github.com/gustionusamba24/kasir-api-go/internal/services"
)

type ProductController struct {
	service services.ProductService
}

// NewProductController creates a new instance of ProductController
func NewProductController(service services.ProductService) *ProductController {
	return &ProductController{
		service: service,
	}
}

// GetAll handles GET /products - retrieve all products
func (c *ProductController) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Check if filtering by category
	categoryIDStr := r.URL.Query().Get("category_id")
	if categoryIDStr != "" {
		categoryID, err := strconv.Atoi(categoryIDStr)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid category ID")
			return
		}

		products, err := c.service.GetByCategoryID(ctx, categoryID)
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
			"data":    products,
		})
		return
	}

	// Get all products
	products, err := c.service.GetAll(ctx)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    products,
	})
}

// GetByID handles GET /products/{id} - retrieve a product by ID
func (c *ProductController) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Extract ID from URL path
	id, err := extractIDFromPath(r, "/products/")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product, err := c.service.GetByID(ctx, id)
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
		"data":    product,
	})
}

// Create handles POST /products - create a new product
func (c *ProductController) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var dto dtos.ProductCreateRequestDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// TODO: Add validation here using validator library

	product, err := c.service.Create(ctx, &dto)
	if err != nil {
		if isNotFoundError(err) {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]interface{}{
		"success": true,
		"data":    product,
		"message": "Product created successfully",
	})
}

// Update handles PUT /products/{id} - update an existing product
func (c *ProductController) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Extract ID from URL path
	id, err := extractIDFromPath(r, "/products/")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var dto dtos.ProductUpdateRequestDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// TODO: Add validation here using validator library

	product, err := c.service.Update(ctx, id, &dto)
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
		"data":    product,
		"message": "Product updated successfully",
	})
}

// Delete handles DELETE /products/{id} - delete a product
func (c *ProductController) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Extract ID from URL path
	id, err := extractIDFromPath(r, "/products/")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	err = c.service.Delete(ctx, id)
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
		"message": "Product deleted successfully",
	})
}
