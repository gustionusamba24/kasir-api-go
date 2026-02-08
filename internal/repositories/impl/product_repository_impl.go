package impl

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/gustionusamba24/kasir-api-go/internal/domain/entities"
	"github.com/gustionusamba24/kasir-api-go/internal/repositories"
)

type productRepositoryImpl struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) repositories.ProductRepository {
	return &productRepositoryImpl{db: db}
}

func (r *productRepositoryImpl) FindAll(ctx context.Context) ([]entities.Product, error) {
	query := `SELECT id, name, price, stock, active, category_id, created_at, updated_at FROM products ORDER BY id`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query products: %w", err)
	}
	defer rows.Close()

	var products []entities.Product
	for rows.Next() {
		var product entities.Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Stock,
			&product.Active,
			&product.CategoryID,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating products: %w", err)
	}

	return products, nil
}

func (r *productRepositoryImpl) FindByID(ctx context.Context, id int) (*entities.Product, error) {
	query := `SELECT id, name, price, stock, active, category_id, created_at, updated_at FROM products WHERE id = $1`

	var product entities.Product
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Stock,
		&product.Active,
		&product.CategoryID,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to find product: %w", err)
	}

	return &product, nil
}

func (r *productRepositoryImpl) FindByCategoryID(ctx context.Context, categoryID int) ([]entities.Product, error) {
	query := `SELECT id, name, price, stock, active, category_id, created_at, updated_at FROM products WHERE category_id = $1 ORDER BY id`

	rows, err := r.db.QueryContext(ctx, query, categoryID)
	if err != nil {
		return nil, fmt.Errorf("failed to query products by category: %w", err)
	}
	defer rows.Close()

	var products []entities.Product
	for rows.Next() {
		var product entities.Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Stock,
			&product.Active,
			&product.CategoryID,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating products: %w", err)
	}

	return products, nil
}

func (r *productRepositoryImpl) FindByFilters(ctx context.Context, name string, active *bool) ([]entities.Product, error) {
	query := `SELECT id, name, price, stock, active, category_id, created_at, updated_at FROM products WHERE 1=1`
	args := []interface{}{}

	// Add name filter with ILIKE for case-insensitive partial matching
	if name != "" {
		query += fmt.Sprintf(" AND name ILIKE $%d", len(args)+1)
		args = append(args, "%"+name+"%")
	}

	// Add active filter
	if active != nil {
		query += fmt.Sprintf(" AND active = $%d", len(args)+1)
		args = append(args, *active)
	}

	query += " ORDER BY id"

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query products by filters: %w", err)
	}
	defer rows.Close()

	var products []entities.Product
	for rows.Next() {
		var product entities.Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Stock,
			&product.Active,
			&product.CategoryID,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating products: %w", err)
	}

	return products, nil
}

func (r *productRepositoryImpl) Create(ctx context.Context, product *entities.Product) error {
	query := `
        INSERT INTO products (name, price, stock, active, category_id, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id
    `

	now := time.Now()
	err := r.db.QueryRowContext(
		ctx,
		query,
		product.Name,
		product.Price,
		product.Stock,
		product.Active,
		product.CategoryID,
		now,
		now,
	).Scan(&product.ID)

	if err != nil {
		return fmt.Errorf("failed to create product: %w", err)
	}

	product.CreatedAt = now
	product.UpdatedAt = now

	return nil
}

func (r *productRepositoryImpl) Update(ctx context.Context, product *entities.Product) error {
	query := `
        UPDATE products 
        SET name = $1, price = $2, stock = $3, active = $4, category_id = $5, updated_at = $6
        WHERE id = $7
    `

	now := time.Now()
	result, err := r.db.ExecContext(
		ctx,
		query,
		product.Name,
		product.Price,
		product.Stock,
		product.Active,
		product.CategoryID,
		now,
		product.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("product not found")
	}

	product.UpdatedAt = now

	return nil
}

func (r *productRepositoryImpl) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM products WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("product not found")
	}

	return nil
}
