# Product Search Implementation Guide

## Overview

This implementation adds search functionality to the Product API with support for:

- **Name search**: Case-insensitive partial matching (e.g., searching "hand" will match "hand sanitizer", "handbag", etc.)
- **Active status filtering**: Filter products by their active/inactive status
- Combined filtering with both parameters

## Database Migration

Before using the new search functionality, run the database migration:

```bash
psql -U your_username -d kasir_db -f migrations/add_active_column_to_products.sql
```

This migration will:

- Add an `active` column (BOOLEAN, default: true) to the products table
- Create indexes for better query performance on `active` column and lowercase `name`
- Set all existing products to active by default

## API Endpoints

### Search Products by Name

**URL**: `GET /products?name=<search_term>`

**Example**:

```bash
curl "http://localhost:8080/products?name=hand"
```

This will return all products containing "hand" in their name (case-insensitive).

### Filter Products by Active Status

**URL**: `GET /products?active=<true|false>`

**Examples**:

```bash
# Get only active products
curl "http://localhost:8080/products?active=true"

# Get only inactive products
curl "http://localhost:8080/products?active=false"
```

### Combined Search

**URL**: `GET /products?name=<search_term>&active=<true|false>`

**Example**:

```bash
# Search for products containing "hand" that are active
curl "http://localhost:8080/products?name=hand&active=true"
```

### Get All Products (no filters)

**URL**: `GET /products`

**Example**:

```bash
curl "http://localhost:8080/products"
```

### Filter by Category (existing functionality)

**URL**: `GET /products?category_id=<category_id>`

**Example**:

```bash
curl "http://localhost:8080/products?category_id=1"
```

## Implementation Details

### Changes Made

1. **Entity Updates** ([internal/domain/entities/product.go](internal/domain/entities/product.go))
   - Added `Active bool` field to Product entity

2. **DTO Updates**
   - [product_dto.go](internal/domain/dtos/product_dto.go): Added Active field
   - [product_create_request_dto.go](internal/domain/dtos/product_create_request_dto.go): Added optional Active field (defaults to true)
   - [product_update_request_dto.go](internal/domain/dtos/product_update_request_dto.go): Added optional Active field

3. **Repository Layer** ([internal/repositories/impl/product_repository_impl.go](internal/repositories/impl/product_repository_impl.go))
   - Added `FindByFilters` method for dynamic query building
   - Updated all existing methods to include `active` column in SELECT, INSERT, and UPDATE queries
   - Uses PostgreSQL ILIKE for case-insensitive partial matching

4. **Service Layer** ([internal/services/impl/product_service_impl.go](internal/services/impl/product_service_impl.go))
   - Added `Search` method to handle name and active filtering

5. **Controller Layer** ([internal/controllers/product_controller.go](internal/controllers/product_controller.go))
   - Updated `GetAll` method to parse and handle `name` and `active` query parameters
   - Updated Swagger documentation

6. **Mapper Updates** ([internal/mappers/product_mapper.go](internal/mappers/product_mapper.go))
   - Updated all mapping methods to include Active field
   - Set default value (true) when Active is not provided in create/update requests

## Testing with Postman

### 1. Create Test Products

**Create Active Product**:

```json
POST /products
{
  "name": "Hand Sanitizer",
  "price": 25000.00,
  "stock": 100,
  "active": true,
  "category_id": 1
}
```

**Create Inactive Product**:

```json
POST /products
{
  "name": "Old Handbag",
  "price": 150000.00,
  "stock": 5,
  "active": false,
  "category_id": 3
}
```

**Create Another Active Product**:

```json
POST /products
{
  "name": "Wireless Handset",
  "price": 450000.00,
  "stock": 20,
  "active": true,
  "category_id": 1
}
```

### 2. Test Search Queries

**Search for "hand"** (should return all 3 products):

```
GET /products?name=hand
```

**Search for "hand" with active=true** (should return 2 products):

```
GET /products?name=hand&active=true
```

**Search for "hand" with active=false** (should return 1 product):

```
GET /products?name=hand&active=false
```

**Get only active products**:

```
GET /products?active=true
```

**Get only inactive products**:

```
GET /products?active=false
```

## Response Format

All successful responses follow this format:

```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "Hand Sanitizer",
      "price": 25000.0,
      "stock": 100,
      "active": true,
      "category_id": 1,
      "created_at": "2026-02-08T10:00:00Z",
      "updated_at": "2026-02-08T10:00:00Z"
    }
  ]
}
```

## Query Parameter Priority

The controller checks query parameters in the following order:

1. If `name` or `active` is present → use Search method
2. If `category_id` is present → use GetByCategoryID method
3. Otherwise → use GetAll method

## Performance Considerations

- Database indexes are created on `active` column and `LOWER(name)` for efficient querying
- The `ILIKE` operator provides case-insensitive search but may be slower on very large tables
- Consider using full-text search for more advanced search requirements

## Error Handling

- Invalid query parameters return appropriate HTTP status codes
- Missing or malformed parameters are handled gracefully
- Database errors are logged and return 500 Internal Server Error

## Notes

- The `active` field defaults to `true` if not provided during product creation
- Empty string for `name` parameter will be ignored (returns all products)
- The search is partial match: "hand" matches "Hand Sanitizer", "handbag", "Wireless Handset"
- Search is case-insensitive: "HAND", "hand", "Hand" all produce the same results
