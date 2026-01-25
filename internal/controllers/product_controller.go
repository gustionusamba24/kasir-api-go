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

// GetAll godoc
// @Summary      Get all products
// @Description  Retrieve a list of all products, optionally filtered by category ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        category_id  query     int  false  "Filter by Category ID"
// @Success      200          {object}  map[string]interface{}  "success response with products data"
// @Failure      400          {object}  map[string]interface{}  "invalid category ID"
// @Failure      404          {object}  map[string]interface{}  "category not found"
// @Failure      500          {object}  map[string]interface{}  "internal server error"
// @Router       /products [get]
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

// GetByID godoc
// @Summary      Get a product by ID
// @Description  Retrieve a single product by its ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  map[string]interface{}  "success response with product data"
// @Failure      400  {object}  map[string]interface{}  "invalid product ID"
// @Failure      404  {object}  map[string]interface{}  "product not found"
// @Failure      500  {object}  map[string]interface{}  "internal server error"
// @Router       /products/{id} [get]
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

// Create godoc
// @Summary      Create a new product
// @Description  Create a new product with the provided data
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      dtos.ProductCreateRequestDto  true  "Product data"
// @Success      201      {object}  map[string]interface{}  "success response with created product"
// @Failure      400      {object}  map[string]interface{}  "invalid request payload"
// @Failure      404      {object}  map[string]interface{}  "category not found"
// @Failure      500      {object}  map[string]interface{}  "internal server error"
// @Router       /products [post]
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

// Update godoc
// @Summary      Update a product
// @Description  Update an existing product by its ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id       path      int  true  "Product ID"
// @Param        product  body      dtos.ProductUpdateRequestDto  true  "Product data"
// @Success      200      {object}  map[string]interface{}  "success response with updated product"
// @Failure      400      {object}  map[string]interface{}  "invalid request"
// @Failure      404      {object}  map[string]interface{}  "product not found"
// @Failure      500      {object}  map[string]interface{}  "internal server error"
// @Router       /products/{id} [put]
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

// Delete godoc
// @Summary      Delete a product
// @Description  Delete a product by its ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  map[string]interface{}  "success response"
// @Failure      400  {object}  map[string]interface{}  "invalid product ID"
// @Failure      404  {object}  map[string]interface{}  "product not found"
// @Failure      500  {object}  map[string]interface{}  "internal server error"
// @Router       /products/{id} [delete]
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
