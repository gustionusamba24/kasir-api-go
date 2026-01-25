package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gustionusamba24/kasir-api-go/internal/config"
	"github.com/gustionusamba24/kasir-api-go/internal/controllers"
	"github.com/gustionusamba24/kasir-api-go/internal/repositories/impl"
	serviceImpl "github.com/gustionusamba24/kasir-api-go/internal/services/impl"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load("../.env"); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	// Connect to the database
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize repositories
	categoryRepo := impl.NewCategoryRepository(db)
	productRepo := impl.NewProductRepository(db)

	// Initialize services
	categoryService := serviceImpl.NewCategoryService(categoryRepo)
	productService := serviceImpl.NewProductService(productRepo, categoryRepo)

	// Initialize controllers
	categoryController := controllers.NewCategoryController(categoryService)
	productController := controllers.NewProductController(productService)

	// Setup routes
	mux := http.NewServeMux()

	// Category routes
	mux.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			categoryController.GetAll(w, r)
		case http.MethodPost:
			categoryController.Create(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/categories/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			categoryController.GetByID(w, r)
		case http.MethodPut:
			categoryController.Update(w, r)
		case http.MethodDelete:
			categoryController.Delete(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Product routes
	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			productController.GetAll(w, r)
		case http.MethodPost:
			productController.Create(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/products/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			productController.GetByID(w, r)
		case http.MethodPut:
			productController.Update(w, r)
		case http.MethodDelete:
			productController.Delete(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	addr := fmt.Sprintf(":%s", port)
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
