# Swagger API Documentation

This project includes Swagger/OpenAPI documentation for all API endpoints.

## Accessing Swagger UI

Once the server is running, you can access the interactive Swagger UI at:

```
http://localhost:8080/swagger/index.html
```

## Available Endpoints

### Categories

- **GET** `/categories` - Get all categories
- **GET** `/categories/{id}` - Get a category by ID
- **POST** `/categories` - Create a new category
- **PUT** `/categories/{id}` - Update a category
- **DELETE** `/categories/{id}` - Delete a category

### Products

- **GET** `/products` - Get all products (supports `?category_id=` query parameter)
- **GET** `/products/{id}` - Get a product by ID
- **POST** `/products` - Create a new product
- **PUT** `/products/{id}` - Update a product
- **DELETE** `/products/{id}` - Delete a product

## Regenerating Swagger Documentation

If you make changes to the API endpoints or add new annotations, regenerate the documentation:

```bash
swag init -g cmd/main.go -o docs
```

## Swagger Annotations

The API documentation is generated from annotations in the code:

- General API info: `cmd/main.go`
- Category endpoints: `internal/controllers/category_controller.go`
- Product endpoints: `internal/controllers/product_controller.go`

## Request/Response Examples

All request and response examples are available in the Swagger UI, including:

- Request body schemas
- Response schemas
- Status codes
- Error responses

Visit the Swagger UI for interactive API testing and detailed documentation.
