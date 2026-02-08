package impl

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/gustionusamba24/kasir-api-go/internal/domain/entities"
	"github.com/gustionusamba24/kasir-api-go/internal/repositories"
)

type transactionRepositoryImpl struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) repositories.TransactionRepository {
	return &transactionRepositoryImpl{db: db}
}

func (r *transactionRepositoryImpl) Create(ctx context.Context, transaction *entities.Transaction) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Insert transaction
	query := `INSERT INTO transactions (total_amount, created_at) VALUES ($1, $2) RETURNING id, created_at`
	now := time.Now()
	err = tx.QueryRowContext(ctx, query, transaction.TotalAmount, now).Scan(&transaction.ID, &transaction.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to create transaction: %w", err)
	}

	// Insert transaction details
	detailQuery := `INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal) VALUES ($1, $2, $3, $4) RETURNING id`
	for i := range transaction.Details {
		detail := &transaction.Details[i]
		detail.TransactionID = transaction.ID
		err = tx.QueryRowContext(ctx, detailQuery, detail.TransactionID, detail.ProductID, detail.Quantity, detail.Subtotal).Scan(&detail.ID)
		if err != nil {
			return fmt.Errorf("failed to create transaction detail: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *transactionRepositoryImpl) FindByID(ctx context.Context, id int) (*entities.Transaction, error) {
	// Get transaction
	query := `SELECT id, total_amount, created_at FROM transactions WHERE id = $1`
	var transaction entities.Transaction
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&transaction.ID,
		&transaction.TotalAmount,
		&transaction.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to find transaction: %w", err)
	}

	// Get transaction details with product names
	detailQuery := `
		SELECT td.id, td.transaction_id, td.product_id, p.name, td.quantity, td.subtotal
		FROM transaction_details td
		LEFT JOIN products p ON td.product_id = p.id
		WHERE td.transaction_id = $1
		ORDER BY td.id
	`
	rows, err := r.db.QueryContext(ctx, detailQuery, id)
	if err != nil {
		return nil, fmt.Errorf("failed to query transaction details: %w", err)
	}
	defer rows.Close()

	var details []entities.TransactionDetail
	for rows.Next() {
		var detail entities.TransactionDetail
		err := rows.Scan(
			&detail.ID,
			&detail.TransactionID,
			&detail.ProductID,
			&detail.ProductName,
			&detail.Quantity,
			&detail.Subtotal,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan transaction detail: %w", err)
		}
		details = append(details, detail)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating transaction details: %w", err)
	}

	transaction.Details = details
	return &transaction, nil
}

func (r *transactionRepositoryImpl) FindAll(ctx context.Context) ([]entities.Transaction, error) {
	// Get all transactions
	query := `SELECT id, total_amount, created_at FROM transactions ORDER BY created_at DESC`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query transactions: %w", err)
	}
	defer rows.Close()

	var transactions []entities.Transaction
	for rows.Next() {
		var transaction entities.Transaction
		err := rows.Scan(
			&transaction.ID,
			&transaction.TotalAmount,
			&transaction.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan transaction: %w", err)
		}
		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating transactions: %w", err)
	}

	// Get details for all transactions
	for i := range transactions {
		detailQuery := `
			SELECT td.id, td.transaction_id, td.product_id, p.name, td.quantity, td.subtotal
			FROM transaction_details td
			LEFT JOIN products p ON td.product_id = p.id
			WHERE td.transaction_id = $1
			ORDER BY td.id
		`
		detailRows, err := r.db.QueryContext(ctx, detailQuery, transactions[i].ID)
		if err != nil {
			return nil, fmt.Errorf("failed to query transaction details: %w", err)
		}

		var details []entities.TransactionDetail
		for detailRows.Next() {
			var detail entities.TransactionDetail
			err := detailRows.Scan(
				&detail.ID,
				&detail.TransactionID,
				&detail.ProductID,
				&detail.ProductName,
				&detail.Quantity,
				&detail.Subtotal,
			)
			if err != nil {
				detailRows.Close()
				return nil, fmt.Errorf("failed to scan transaction detail: %w", err)
			}
			details = append(details, detail)
		}
		detailRows.Close()

		if err = detailRows.Err(); err != nil {
			return nil, fmt.Errorf("error iterating transaction details: %w", err)
		}

		transactions[i].Details = details
	}

	return transactions, nil
}

func (r *transactionRepositoryImpl) CreateDetail(ctx context.Context, detail *entities.TransactionDetail) error {
	query := `INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal) VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, detail.TransactionID, detail.ProductID, detail.Quantity, detail.Subtotal).Scan(&detail.ID)
	if err != nil {
		return fmt.Errorf("failed to create transaction detail: %w", err)
	}
	return nil
}

func (r *transactionRepositoryImpl) GetTodayRevenue(ctx context.Context) (int, error) {
	query := `
		SELECT COALESCE(SUM(total_amount), 0) 
		FROM transactions 
		WHERE DATE(created_at) = CURRENT_DATE
	`
	var totalRevenue int
	err := r.db.QueryRowContext(ctx, query).Scan(&totalRevenue)
	if err != nil {
		return 0, fmt.Errorf("failed to get today's revenue: %w", err)
	}
	return totalRevenue, nil
}

func (r *transactionRepositoryImpl) GetTodayTransactionCount(ctx context.Context) (int, error) {
	query := `
		SELECT COUNT(*) 
		FROM transactions 
		WHERE DATE(created_at) = CURRENT_DATE
	`
	var count int
	err := r.db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to get today's transaction count: %w", err)
	}
	return count, nil
}

func (r *transactionRepositoryImpl) GetTodayBestSellingProduct(ctx context.Context) (string, int, error) {
	query := `
		SELECT p.name, SUM(td.quantity) as total_qty
		FROM transaction_details td
		JOIN transactions t ON td.transaction_id = t.id
		JOIN products p ON td.product_id = p.id
		WHERE DATE(t.created_at) = CURRENT_DATE
		GROUP BY p.id, p.name
		ORDER BY total_qty DESC
		LIMIT 1
	`
	var productName string
	var qtySold int
	err := r.db.QueryRowContext(ctx, query).Scan(&productName, &qtySold)
	if err == sql.ErrNoRows {
		return "", 0, nil
	}
	if err != nil {
		return "", 0, fmt.Errorf("failed to get today's best selling product: %w", err)
	}
	return productName, qtySold, nil
}
