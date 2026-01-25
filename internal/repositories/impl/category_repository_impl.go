package impl

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/gustionusamba24/kasir-api-go/internal/domain/entities"
	"github.com/gustionusamba24/kasir-api-go/internal/repositories"
)

type categoryRepositoryImpl struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) repositories.CategoryRepository {
	return &categoryRepositoryImpl{db: db}
}

func (r *categoryRepositoryImpl) FindAll(ctx context.Context) ([]entities.Category, error) {
	query := `SELECT id, name, description, created_at, updated_at FROM categories ORDER BY id`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query categories: %w", err)
	}
	defer rows.Close()

	var categories []entities.Category
	for rows.Next() {
		var category entities.Category
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan category: %w", err)
		}
		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating categories: %w", err)
	}

	return categories, nil
}

func (r *categoryRepositoryImpl) FindByID(ctx context.Context, id int) (*entities.Category, error) {
	query := `SELECT id, name, description, created_at, updated_at FROM categories WHERE id = $1`

	var category entities.Category
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&category.ID,
		&category.Name,
		&category.Description,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("failed to find category: %w", err)
	}

	return &category, nil
}

func (r *categoryRepositoryImpl) Create(ctx context.Context, category *entities.Category) error {
	query := `
        INSERT INTO categories (name, description, created_at, updated_at)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `

	now := time.Now()
	err := r.db.QueryRowContext(
		ctx,
		query,
		category.Name,
		category.Description,
		now,
		now,
	).Scan(&category.ID)

	if err != nil {
		return fmt.Errorf("failed to create category: %w", err)
	}

	category.CreatedAt = now
	category.UpdatedAt = now

	return nil
}

func (r *categoryRepositoryImpl) Update(ctx context.Context, category *entities.Category) error {
	query := `
        UPDATE categories 
        SET name = $1, description = $2, updated_at = $3
        WHERE id = $4
    `

	now := time.Now()
	result, err := r.db.ExecContext(
		ctx,
		query,
		category.Name,
		category.Description,
		now,
		category.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update category: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("category not found")
	}

	category.UpdatedAt = now

	return nil
}

func (r *categoryRepositoryImpl) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM categories WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete category: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("category not found")
	}

	return nil
}
