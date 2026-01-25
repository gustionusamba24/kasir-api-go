# Kasir API - Point of Sale REST API

A simple and efficient REST API for Point of Sale (POS) systems built with Go and PostgreSQL. This API provides CRUD operations for managing categories and products with comprehensive Swagger documentation.

## 📋 About the Project

Kasir API is a backend service designed for point of sale applications. It provides a robust API for managing product categories and products with features including:

- **RESTful Architecture**: Clean and intuitive REST API endpoints
- **PostgreSQL Database**: Reliable data persistence with PostgreSQL
- **Swagger Documentation**: Interactive API documentation with Swagger UI
- **Clean Architecture**: Organized codebase following separation of concerns
- **Repository Pattern**: Abstract database operations for better maintainability
- **Service Layer**: Business logic separated from controllers
- **DTO Pattern**: Data Transfer Objects for request/response handling

### Built With

- **Go 1.25.6** - Programming language
- **PostgreSQL** - Database (using Neon serverless PostgreSQL)
- **net/http** - HTTP server (Go standard library)
- **lib/pq** - PostgreSQL driver
- **Swagger/OpenAPI** - API documentation
- **godotenv** - Environment variable management

## 🔧 CRUD Specification

### Categories API

#### Get All Categories

- **Endpoint**: `GET /categories`
- **Description**: Retrieve a list of all categories
- **Response**: 200 OK with array of categories

#### Get Category by ID

- **Endpoint**: `GET /categories/{id}`
- **Description**: Retrieve a single category by its ID
- **Parameters**: `id` (path parameter) - Category ID
- **Response**:
  - 200 OK with category data
  - 404 Not Found if category doesn't exist

#### Create Category

- **Endpoint**: `POST /categories`
- **Description**: Create a new category
- **Request Body**:
  ```json
  {
    "name": "string (required, min 3, max 100 characters)",
    "description": "string (optional, max 500 characters)"
  }
  ```
- **Response**: 201 Created with created category data

#### Update Category

- **Endpoint**: `PUT /categories/{id}`
- **Description**: Update an existing category
- **Parameters**: `id` (path parameter) - Category ID
- **Request Body**:
  ```json
  {
    "name": "string (required, min 3, max 100 characters)",
    "description": "string (optional, max 500 characters)"
  }
  ```
- **Response**:
  - 200 OK with updated category data
  - 404 Not Found if category doesn't exist

#### Delete Category

- **Endpoint**: `DELETE /categories/{id}`
- **Description**: Delete a category by its ID
- **Parameters**: `id` (path parameter) - Category ID
- **Response**:
  - 200 OK with success message
  - 404 Not Found if category doesn't exist

### Products API

#### Get All Products

- **Endpoint**: `GET /products`
- **Description**: Retrieve a list of all products
- **Query Parameters**:
  - `category_id` (optional) - Filter products by category ID
- **Response**: 200 OK with array of products

#### Get Product by ID

- **Endpoint**: `GET /products/{id}`
- **Description**: Retrieve a single product by its ID
- **Parameters**: `id` (path parameter) - Product ID
- **Response**:
  - 200 OK with product data
  - 404 Not Found if product doesn't exist

#### Create Product

- **Endpoint**: `POST /products`
- **Description**: Create a new product
- **Request Body**:
  ```json
  {
    "name": "string (required, min 3, max 100 characters)",
    "price": "number (required, must be greater than 0)",
    "stock": "integer (required, must be >= 0)",
    "category_id": "integer (optional, must be > 0 if provided)"
  }
  ```
- **Response**:
  - 201 Created with created product data
  - 404 Not Found if category_id doesn't exist

#### Update Product

- **Endpoint**: `PUT /products/{id}`
- **Description**: Update an existing product
- **Parameters**: `id` (path parameter) - Product ID
- **Request Body**:
  ```json
  {
    "name": "string (required, min 3, max 100 characters)",
    "price": "number (required, must be greater than 0)",
    "stock": "integer (required, must be >= 0)",
    "category_id": "integer (optional, must be > 0 if provided)"
  }
  ```
- **Response**:
  - 200 OK with updated product data
  - 404 Not Found if product or category doesn't exist

#### Delete Product

- **Endpoint**: `DELETE /products/{id}`
- **Description**: Delete a product by its ID
- **Parameters**: `id` (path parameter) - Product ID
- **Response**:
  - 200 OK with success message
  - 404 Not Found if product doesn't exist

### Response Format

All responses follow a consistent JSON structure:

**Success Response:**

```json
{
  "success": true,
  "data": { ... },
  "message": "Optional success message"
}
```

**Error Response:**

```json
{
  "success": false,
  "error": "Error message description"
}
```

## 📁 Folder Structure

```
kasir-api-go/
├── cmd/
│   └── main.go                    # Application entry point with route definitions
│
├── internal/
│   ├── config/
│   │   └── database.go            # Database connection configuration
│   │
│   ├── controllers/
│   │   ├── category_controller.go # Category HTTP handlers
│   │   ├── product_controller.go  # Product HTTP handlers
│   │   └── helpers.go             # Controller helper functions
│   │
│   ├── domain/
│   │   ├── dtos/                  # Data Transfer Objects
│   │   │   ├── category_create_request_dto.go
│   │   │   ├── category_update_request_dto.go
│   │   │   ├── category_dto.go
│   │   │   ├── product_create_request_dto.go
│   │   │   ├── product_update_request_dto.go
│   │   │   └── product_dto.go
│   │   │
│   │   └── entities/              # Database entities
│   │       ├── category.go
│   │       └── product.go
│   │
│   ├── mappers/                   # Entity to DTO converters
│   │   ├── category_mapper.go
│   │   └── product_mapper.go
│   │
│   ├── repositories/              # Data access layer
│   │   ├── category_repository.go      # Category repository interface
│   │   ├── product_repository.go       # Product repository interface
│   │   └── impl/                       # Repository implementations
│   │       ├── category_repository_impl.go
│   │       └── product_repository_impl.go
│   │
│   └── services/                  # Business logic layer
│       ├── category_service.go         # Category service interface
│       ├── product_service.go          # Product service interface
│       └── impl/                       # Service implementations
│           ├── category_service_impl.go
│           └── product_service_impl.go
│
├── docs/                          # Swagger documentation (auto-generated)
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
│
├── .env                           # Environment variables (not in git)
├── .gitignore                     # Git ignore rules
├── go.mod                         # Go module dependencies
├── go.sum                         # Go module checksums
├── README.md                      # This file
└── SWAGGER.md                     # Swagger documentation guide
```

### Architecture Explanation

The project follows a **Clean Architecture** approach with clear separation of concerns:

- **cmd/**: Entry point of the application. Contains `main.go` which initializes dependencies and defines routes.

- **internal/config/**: Configuration management, including database connection setup.

- **internal/controllers/**: HTTP handlers (presentation layer). Handles HTTP requests, validates input, and returns responses.

- **internal/domain/**: Core business domain models
  - **entities/**: Database models that represent tables
  - **dtos/**: Data Transfer Objects for API requests and responses

- **internal/mappers/**: Convert between entities and DTOs to keep layers independent.

- **internal/repositories/**: Data access layer (persistence). Handles all database operations.
  - **Interface**: Defines contracts for data operations
  - **impl/**: Concrete implementations of repository interfaces

- **internal/services/**: Business logic layer. Contains application business rules and orchestrates data flow.
  - **Interface**: Defines contracts for business operations
  - **impl/**: Concrete implementations of service interfaces

- **docs/**: Auto-generated Swagger/OpenAPI documentation files.

This structure promotes:

- **Testability**: Easy to mock interfaces for unit testing
- **Maintainability**: Clear separation makes code easier to understand and modify
- **Scalability**: New features can be added without affecting existing code
- **Flexibility**: Implementations can be swapped without changing interfaces

## 🚀 How to Run This Project

### Prerequisites

Before running this project, ensure you have the following installed:

- **Go** (version 1.20 or higher)
- **PostgreSQL** (or access to a PostgreSQL database)
- **Git** (for cloning the repository)

### Installation Steps

1. **Clone the repository**

   ```bash
   git clone https://github.com/gustionusamba24/kasir-api-go.git
   cd kasir-api-go
   ```

2. **Install Go dependencies**

   ```bash
   go mod download
   go mod tidy
   ```

3. **Set up environment variables**

   Create a `.env` file in the root directory:

   ```bash
   cp .env.example .env  # If you have an example file
   # OR create new .env file
   touch .env
   ```

   Add your database configuration to `.env`:

   ```env
   DB_URL=postgresql://username:password@host:port/database?sslmode=require
   PORT=8080
   ```

   Replace with your actual PostgreSQL credentials:
   - `username`: Your database username
   - `password`: Your database password
   - `host`: Database host (e.g., localhost or remote host)
   - `port`: Database port (default: 5432)
   - `database`: Your database name

4. **Set up the database**

   Create the database tables by running these SQL commands in your PostgreSQL database:

   ```sql
   -- Create categories table
   CREATE TABLE categories (
       id SERIAL PRIMARY KEY,
       name VARCHAR(100) NOT NULL,
       description TEXT,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
   );

   -- Create products table
   CREATE TABLE products (
       id SERIAL PRIMARY KEY,
       name VARCHAR(100) NOT NULL,
       price DECIMAL(10, 2) NOT NULL,
       stock INTEGER NOT NULL DEFAULT 0,
       category_id INTEGER REFERENCES categories(id) ON DELETE SET NULL,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
   );

   -- Create indexes for better performance
   CREATE INDEX idx_products_category_id ON products(category_id);
   ```

5. **Install Swagger CLI (Optional, for regenerating docs)**

   ```bash
   go install github.com/swaggo/swag/cmd/swag@latest
   ```

6. **Run the application**

   ```bash
   cd cmd
   go run main.go
   ```

   Or build and run the binary:

   ```bash
   go build -o kasir-api cmd/main.go
   ./kasir-api
   ```

7. **Verify the server is running**

   You should see output similar to:

   ```
   2026/01/25 15:00:00 Database connected successfully
   2026/01/25 15:00:00 Server starting on port 8080
   ```

8. **Access the API**
   - **API Base URL**: http://localhost:8080
   - **Swagger UI**: http://localhost:8080/swagger/index.html
   - **Test endpoint**: http://localhost:8080/categories

### Quick Test

Test if the API is working by creating a category:

```bash
curl -X POST http://localhost:8080/categories \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Electronics",
    "description": "Electronic devices and accessories"
  }'
```

Or open http://localhost:8080/swagger/index.html in your browser for interactive API testing.

### Development

**Running in development mode:**

```bash
cd cmd
go run main.go
```

**Regenerating Swagger documentation** (after updating API annotations):

```bash
swag init -g cmd/main.go -o docs
```

**Running tests** (if tests are added):

```bash
go test ./...
```

### Troubleshooting

**Database connection issues:**

- Verify your `.env` file has the correct database credentials
- Ensure PostgreSQL is running and accessible
- Check firewall settings if using a remote database

**Port already in use:**

- Change the `PORT` in your `.env` file to a different port
- Or kill the process using port 8080:
  ```bash
  lsof -ti:8080 | xargs kill -9
  ```

**Swagger UI not loading:**

- Make sure you've run `swag init` to generate documentation
- Check that the `docs/` folder exists with generated files

## 📚 API Documentation

Interactive API documentation is available via Swagger UI once the server is running:

**Swagger UI**: http://localhost:8080/swagger/index.html

For detailed Swagger documentation guide, see [SWAGGER.md](SWAGGER.md)

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📝 License

This project is licensed under the Apache License 2.0 - see the LICENSE file for details.

## 📧 Contact

For support or questions, please contact: support@kasir-api.com

---

**Happy Coding! 🚀**
