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

// GetAll godoc
// @Summary      Get all categories
// @Description  Retrieve a list of all categories
// @Tags         categories
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "success response with categories data"
// @Failure      500  {object}  map[string]interface{}  "internal server error"
// @Router       /categories [get]
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

// GetByID godoc
// @Summary      Get a category by ID
// @Description  Retrieve a single category by its ID
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Category ID"
// @Success      200  {object}  map[string]interface{}  "success response with category data"
// @Failure      400  {object}  map[string]interface{}  "invalid category ID"
// @Failure      404  {object}  map[string]interface{}  "category not found"
// @Failure      500  {object}  map[string]interface{}  "internal server error"
// @Router       /categories/{id} [get]
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

// Create godoc
// @Summary      Create a new category
// @Description  Create a new category with the provided data
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        category  body      dtos.CategoryCreateRequestDto  true  "Category data"
// @Success      201       {object}  map[string]interface{}  "success response with created category"
// @Failure      400       {object}  map[string]interface{}  "invalid request payload"
// @Failure      500       {object}  map[string]interface{}  "internal server error"
// @Router       /categories [post]
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

// Update godoc
// @Summary      Update a category
// @Description  Update an existing category by its ID
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        id        path      int  true  "Category ID"
// @Param        category  body      dtos.CategoryUpdateRequestDto  true  "Category data"
// @Success      200       {object}  map[string]interface{}  "success response with updated category"
// @Failure      400       {object}  map[string]interface{}  "invalid request"
// @Failure      404       {object}  map[string]interface{}  "category not found"
// @Failure      500       {object}  map[string]interface{}  "internal server error"
// @Router       /categories/{id} [put]
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

// Delete godoc
// @Summary      Delete a category
// @Description  Delete a category by its ID
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Category ID"
// @Success      200  {object}  map[string]interface{}  "success response"
// @Failure      400  {object}  map[string]interface{}  "invalid category ID"
// @Failure      404  {object}  map[string]interface{}  "category not found"
// @Failure      500  {object}  map[string]interface{}  "internal server error"
// @Router       /categories/{id} [delete]
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
