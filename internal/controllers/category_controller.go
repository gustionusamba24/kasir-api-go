package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gustionusamba24/kasir-api-go/internal/domain/dtos"
	"github.com/gustionusamba24/kasir-api-go/internal/services"
)

type CategoryController struct {
	service services.CategoryService
}

// NewCategoryController creates a new instance of CategoryController
func NewCategoryController(service services.CategoryService) *CategoryController {
	return &CategoryController{
		service: service,
	}
}

// GetAll handles GET /categories - retrieve all categories
func (c *CategoryController) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	categories, err := c.service.GetAll(ctx)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    categories,
	})
}

// GetByID handles GET /categories/{id} - retrieve a category by ID
func (c *CategoryController) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Extract ID from URL path
	id, err := extractIDFromPath(r, "/categories/")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	category, err := c.service.GetByID(ctx, id)
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
		"data":    category,
	})
}

// Create handles POST /categories - create a new category
func (c *CategoryController) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var dto dtos.CategoryCreateRequestDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// TODO: Add validation here using validator library

	category, err := c.service.Create(ctx, &dto)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]interface{}{
		"success": true,
		"data":    category,
		"message": "Category created successfully",
	})
}

// Update handles PUT /categories/{id} - update an existing category
func (c *CategoryController) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Extract ID from URL path
	id, err := extractIDFromPath(r, "/categories/")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	var dto dtos.CategoryUpdateRequestDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// TODO: Add validation here using validator library

	category, err := c.service.Update(ctx, id, &dto)
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
		"data":    category,
		"message": "Category updated successfully",
	})
}

// Delete handles DELETE /categories/{id} - delete a category
func (c *CategoryController) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Extract ID from URL path
	id, err := extractIDFromPath(r, "/categories/")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid category ID")
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
		"message": "Category deleted successfully",
	})
}
