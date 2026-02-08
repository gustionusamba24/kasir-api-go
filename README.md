# Kasir API - Point of Sale REST API

A simple and efficient REST API for Point of Sale (POS) systems built with Go and PostgreSQL. This API provides complete POS functionality including product management, inventory tracking, advanced search features, and transaction processing with comprehensive Swagger documentation.

## 📋 About the Project

Kasir API is a backend service designed for point of sale applications. It provides a robust API for managing product categories, products, and transactions with features including:

- **RESTful Architecture**: Clean and intuitive REST API endpoints
- **PostgreSQL Database**: Reliable data persistence with PostgreSQL
- **Advanced Product Search**: Search by name with partial matching and filter by active status
- **Transaction Management**: Complete checkout system with automatic stock updates
- **Reports & Analytics**: Daily and date range transaction reports with revenue tracking and best selling products
- **Inventory Tracking**: Real-time stock management with product availability checks
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
- **Description**: Retrieve a list of all products with optional filters
- **Query Parameters**:
  - `category_id` (optional) - Filter products by category ID
  - `name` (optional) - Search products by name (case-insensitive partial match)
  - `active` (optional) - Filter by active status (true/false)
- **Examples**:
  - `GET /products` - Get all products
  - `GET /products?category_id=1` - Get products by category
  - `GET /products?name=hand` - Search products containing "hand"
  - `GET /products?active=true` - Get only active products
  - `GET /products?name=laptop&active=true` - Combined filters
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
    "active": "boolean (optional, default: true)",
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
    "active": "boolean (optional)",
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

### Transactions API

#### Checkout - Create Transaction

- **Endpoint**: `POST /transactions/checkout`
- **Description**: Create a new transaction (checkout). Automatically validates products, checks stock availability, calculates totals, and updates inventory.
- **Request Body**:
  ```json
  {
    "items": [
      {
        "product_id": "integer (required, must be > 0)",
        "quantity": "integer (required, must be > 0)"
      }
    ]
  }
  ```
- **Example**:
  ```json
  {
    "items": [
      {
        "product_id": 1,
        "quantity": 2
      },
      {
        "product_id": 3,
        "quantity": 1
      }
    ]
  }
  ```
- **Response**:
  - 201 Created with transaction details including:
    - Transaction ID
    - Total amount
    - Transaction details with product names, quantities, and subtotals
    - Created timestamp
  - 400 Bad Request if insufficient stock or inactive products
  - 404 Not Found if product doesn't exist

#### Get All Transactions

- **Endpoint**: `GET /transactions`
- **Description**: Retrieve a list of all transactions with their details
- **Response**: 200 OK with array of transactions

#### Get Transaction by ID

- **Endpoint**: `GET /transactions/{id}`
- **Description**: Retrieve a single transaction by its ID with complete details
- **Parameters**: `id` (path parameter) - Transaction ID
- **Response**:
  - 200 OK with transaction data including all line items
  - 404 Not Found if transaction doesn't exist

### Reports API

#### Get Today's Report

- **Endpoint**: `GET /report/today`
- **Description**: Retrieve today's transaction report including total revenue, transaction count, and best selling product
- **Response**: 200 OK with today's report data
  ```json
  {
    "success": true,
    "data": {
      "total_revenue": 150000,
      "total_transactions": 5,
      "best_selling_product": {
        "name": "Product Name",
        "qty_sold": 10
      }
    }
  }
  ```

#### Get Date Range Report

- **Endpoint**: `GET /report?start_date={YYYY-MM-DD}&end_date={YYYY-MM-DD}`
- **Description**: Retrieve transaction report for a specific date range including total revenue, transaction count, and best selling product
- **Query Parameters**:
  - `start_date` (required) - Start date in YYYY-MM-DD format
  - `end_date` (required) - End date in YYYY-MM-DD format
- **Example**: `GET /report?start_date=2026-01-01&end_date=2026-02-01`
- **Response**: 200 OK with date range report data
  ```json
  {
    "success": true,
    "data": {
      "start_date": "2026-01-01",
      "end_date": "2026-02-01",
      "total_revenue": 500000,
      "total_transactions": 25,
      "best_selling_product": {
        "name": "Product Name",
        "qty_sold": 50
      }
    }
  }
  ```
- **Response**: 400 Bad Request if start_date or end_date is missing

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
│   │   ├── category_controller.go      # Category HTTP handlers
│   │   ├── product_controller.go       # Product HTTP handlers
│   │   ├── transaction_controller.go   # Transaction HTTP handlers
│   │   └── helpers.go                  # Controller helper functions
│   │
│   ├── domain/
│   │   ├── dtos/                  # Data Transfer Objects
│   │   │   ├── category_create_request_dto.go
│   │   │   ├── category_update_request_dto.go
│   │   │   ├── category_dto.go
│   │   │   ├── product_create_request_dto.go
│   │   │   ├── product_update_request_dto.go
│   │   │   ├── product_dto.go
│   │   │   ├── transaction_create_request_dto.go
│   │   │   └── transaction_dto.go
│   │   │
│   │   └── entities/              # Database entities
│   │       ├── category.go
│   │       ├── product.go
│   │       └── transaction.go
│   │
│   ├── mappers/                   # Entity to DTO converters
│   │   ├── category_mapper.go
│   │   ├── product_mapper.go
│   │   └── transaction_mapper.go
│   │
│   ├── repositories/              # Data access layer
│   │   ├── category_repository.go           # Category repository interface
│   │   ├── product_repository.go            # Product repository interface
│   │   ├── transaction_repository.go        # Transaction repository interface
│   │   └── impl/                            # Repository implementations
│   │       ├── category_repository_impl.go
│   │       ├── product_repository_impl.go
│   │       └── transaction_repository_impl.go
│   │
│   └── services/                  # Business logic layer
│       ├── category_service.go              # Category service interface
│       ├── product_service.go               # Product service interface
│       ├── transaction_service.go           # Transaction service interface
│       └── impl/                            # Service implementations
│           ├── category_service_impl.go
│           ├── product_service_impl.go
│           └── transaction_service_impl.go
│
├── docs/                          # Swagger documentation (auto-generated)
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
│
├── migrations/                    # Database migration scripts
│   └── add_active_column_to_products.sql
│
├── .env                           # Environment variables (not in git)
├── .gitignore                     # Git ignore rules
├── go.mod                         # Go module dependencies
├── go.sum                         # Go module checksums
├── README.md                      # This file
├── SWAGGER.md                     # Swagger documentation guide
├── PRODUCT_SEARCH_GUIDE.md        # Product search feature documentation
├── seed_data.sql                  # Sample data for testing
└── Kasir_API.postman_collection.json  # Postman collection for API testing
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
       active BOOLEAN NOT NULL DEFAULT true,
       category_id INTEGER REFERENCES categories(id) ON DELETE SET NULL,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
   );

   -- Create transactions table
   CREATE TABLE transactions (
       id SERIAL PRIMARY KEY,
       total_amount INT NOT NULL,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
   );

   -- Create transaction_details table
   CREATE TABLE transaction_details (
       id SERIAL PRIMARY KEY,
       transaction_id INT REFERENCES transactions(id) ON DELETE CASCADE,
       product_id INT REFERENCES products(id),
       quantity INT NOT NULL,
       subtotal INT NOT NULL
   );

   -- Create indexes for better performance
   CREATE INDEX idx_products_category_id ON products(category_id);
   CREATE INDEX idx_products_active ON products(active);
   CREATE INDEX idx_products_name_lower ON products(LOWER(name));
   CREATE INDEX idx_transaction_details_transaction_id ON transaction_details(transaction_id);
   ```

   Or use the migration script:

   ```bash
   psql -d your_database_name -f migrations/add_active_column_to_products.sql
   ```

5. **Load sample data (Optional)**

   To populate your database with test data for development/testing:

   ```bash
   psql -d your_database_name -f seed_data.sql
   ```

   Or using a PostgreSQL client, execute the SQL script in `seed_data.sql`. This will add:
   - 5 sample categories (Electronics, Food & Beverages, Clothing, Books, Home & Garden)
   - 25+ sample products across different categories

6. **Install Swagger CLI (Optional, for regenerating docs)**

   ```bash
   go install github.com/swaggo/swag/cmd/swag@latest
   ```

7. **Run the application**

   ```bash
   cd cmd
   go run main.go
   ```

   Or build and run the binary:

   ```bash
   go build -o kasir-api cmd/main.go
   ./kasir-api
   ```

8. **Verify the server is running**

   You should see output similar to:

   ```
   2026/01/25 15:00:00 Database connected successfully
   2026/01/25 15:00:00 Server starting on port 8080
   ```

9. **Access the API**
   - **API Base URL**: http://localhost:8080
   - **Swagger UI**: http://localhost:8080/swagger/index.html
   - **Test endpoint**: http://localhost:8080/categories

## 🧪 Testing the API

### Using Postman

A complete Postman collection is included in this repository for easy API testing.

#### Import Postman Collection

1. **Download and install** [Postman](https://www.postman.com/downloads/)

2. **Import the collection**:
   - Open Postman
   - Click "Import" button
   - Select `Kasir_API.postman_collection.json` from the project root
   - The collection will be added to your workspace

3. **Configure environment** (Optional):
   - The collection uses `{{base_url}}` variable set to `http://localhost:8080` by default
   - To change the base URL, edit the collection variable or create a Postman environment

4. **Start testing**:
   - Make sure your server is running (`go run cmd/main.go`)
   - Execute requests from the collection
   - All CRUD operations for Categories and Products are included

#### Postman Collection Contents

The collection includes:

**Categories Endpoints:**

- Get All Categories
- Get Category by ID
- Create Category
- Update Category
- Delete Category

**Products Endpoints:**

- Get All Products
- Get Products by Category (with query parameter)
- Search Products by Name
- Filter Products by Active Status
- Search with Multiple Filters (name + active)
- Get Product by ID
- Create Product (with category and active status)
- Create Product (without category)
- Update Product
- Delete Product

**Transactions Endpoints:**

- Checkout - Create Transaction (single item)
- Checkout - Create Transaction (multiple items)
- Get All Transactions
- Get Transaction by ID

### Using cURL

Test if the API is working by creating a category:

```bash
curl -X POST http://localhost:8080/categories \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Electronics",
    "description": "Electronic devices and accessories"
  }'
```

Or open http://localhost:8080/swagger/index.html in your browser for interactive API testing with Swagger UI.

### Sample Test Data

The `seed_data.sql` file contains sample data with:

- **5 Categories**: Electronics, Food & Beverages, Clothing, Books, Home & Garden
- **25+ Products**: Various products across different categories with realistic prices and stock levels

This data is perfect for testing all API endpoints without manually creating entries.

**Testing Transactions:**

After loading the seed data, you can test the checkout endpoint with sample product IDs:

```bash
curl -X POST http://localhost:8080/transactions/checkout \
  -H "Content-Type: application/json" \
  -d '{
    "items": [
      {"product_id": 1, "quantity": 2},
      {"product_id": 3, "quantity": 1}
    ]
  }'
```

The transaction will automatically:

- Validate product existence and active status
- Check stock availability
- Calculate subtotals and total amount
- Update product stock levels
- Return complete transaction details with product names

## 🛠️ Development

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

## ✨ Key Features

### 1. Product Management

- **CRUD Operations**: Complete create, read, update, and delete functionality
- **Category Assignment**: Link products to categories or leave uncategorized
- **Stock Tracking**: Real-time inventory management
- **Active Status**: Mark products as active/inactive for availability control

### 2. Advanced Product Search

- **Partial Name Search**: Search products by name with case-insensitive partial matching
  - Example: Searching "hand" will find "Hand Sanitizer", "Handbag", "Handset"
- **Active Status Filter**: Filter products by active/inactive status
- **Combined Filters**: Use multiple filters simultaneously (name + active)
- **Category Filter**: Get all products within a specific category

### 3. Transaction Processing

- **Checkout System**: Complete point-of-sale transaction processing
- **Automatic Calculations**: System automatically calculates subtotals and total amounts
- **Stock Validation**: Real-time stock availability checks before transaction
- **Product Validation**: Ensures products exist and are active
- **Automatic Stock Updates**: Inventory automatically decreases after successful checkout
- **Transaction History**: View all past transactions with complete details
- **Detailed Reports**: Each transaction includes product names, quantities, and individual subtotals

### 4. Reports & Analytics

- **Today's Report**: Get real-time insights for current day's transactions
  - Total revenue for today
  - Count of transactions completed
  - Best selling product with quantity sold
- **Date Range Reports**: Historical transaction analysis for any date range
  - Revenue tracking across custom periods
  - Transaction volume analysis
  - Best seller identification for specified dates
- **Automated Calculations**: All reports generated automatically from transaction data
- **Business Intelligence**: Make data-driven decisions with detailed sales analytics

### 5. Category Management

- **Hierarchical Organization**: Organize products into logical categories
- **Flexible Assignment**: Products can belong to a category or remain uncategorized
- **Easy Maintenance**: Simple CRUD operations for category management

### 6. Data Integrity

- **Database Transactions**: Uses SQL transactions for data consistency
- **Foreign Key Constraints**: Maintains referential integrity
- **Validation**: Input validation at multiple layers (DTO, service, repository)
- **Error Handling**: Comprehensive error messages for debugging

### 7. Developer Experience

- **Swagger Documentation**: Interactive API testing and documentation
- **Postman Collection**: Pre-configured API requests for quick testing
- **Sample Data**: Seed data included for immediate testing
- **Clean Architecture**: Well-organized code structure for easy maintenance and extension
