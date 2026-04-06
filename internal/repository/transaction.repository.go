package repository

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/model"
)

type TransactionRepository struct {
	db *pgxpool.Pool
}

func (r *TransactionRepository) UpdateTransactionStatus(ctx context.Context, transactionID int, status string) error {
	panic("unimplemented")
}

func NewTransactionRepository(db *pgxpool.Pool) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func generateTransactionCode() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return fmt.Sprintf("%05d-%05d",
		r.Intn(100000),
		r.Intn(100000),
	)
}

func (r *TransactionRepository) GetTransactionItems(ctx context.Context, transactionID int) ([]model.TransactionItem, error) {

	rows, err := r.db.Query(ctx, `
		SELECT 
			product_id,
			qty,
			size,
			variant,
			price
		FROM transaction_products
		WHERE transaction_id = $1
	`, transactionID)

	if err != nil {
		return nil, err
	}

	items, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.TransactionItem])
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r *TransactionRepository) GetTransactionsByUserID(ctx context.Context, userID int) ([]dto.TransactionResult, error) {
	rows, err := r.db.Query(ctx, `
		SELECT 
			id,
			transaction_code,
			full_name,
			email,
			address,
			delivery_method,
			subtotal_price,
			total_price,
			delivery_fee,
			tax,
			payment_method,
			status,
			created_at
		FROM transactions
		WHERE user_id = $1
		ORDER BY created_at DESC
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []dto.TransactionResult

	for rows.Next() {
		var t dto.TransactionResult

		err := rows.Scan(
			&t.ID,
			&t.TransactionCode,
			&t.FullName,
			&t.Email,
			&t.Address,
			&t.DeliveryMethod,
			&t.SubtotalPrice,
			&t.TotalPrice,
			&t.DeliveryFee,
			&t.Tax,
			&t.PaymentMethod,
			&t.Status,
			&t.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		transactions = append(transactions, t)
	}

	return transactions, nil
}

func (r *TransactionRepository) CreateTransaction(ctx context.Context, userID int, req dto.CreateTransactionRequest) (int, string, error) {

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return 0, "", err
	}

	// rollback kalau error
	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	var transactionID int
	var transactionCode string

	// retry generate code (max 3x)
	for i := range 3 {
		transactionCode = generateTransactionCode()

		err = tx.QueryRow(ctx, `
			INSERT INTO transactions 
			(user_id, transaction_code, full_name, email, address, delivery_method, subtotal_price, total_price, delivery_fee, tax, payment_method, status)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
			RETURNING id
		`,
			userID,
			transactionCode,
			req.FullName,
			req.Email,
			req.Address,
			req.DeliveryMethod,
			req.SubtotalPrice,
			req.TotalPrice,
			req.DeliveryFee,
			req.Tax,
			req.PaymentMethod,
			"pending",
		).Scan(&transactionID)

		if err == nil {
			break
		}

		if i == 2 {
			return 0, "", fmt.Errorf("failed to generate unique transaction code: %w", err)
		}
	}

	for _, item := range req.Items {

		// 1. cek & update stock (ANTI OVERSELL)
		cmd, err := tx.Exec(ctx, `
			UPDATE products
			SET stock = stock - $1
			WHERE id = $2 AND stock >= $1
		`,
			item.Qty,
			item.ProductID,
		)
		if err != nil {
			return 0, "", err
		}

		if cmd.RowsAffected() == 0 {
			return 0, "", fmt.Errorf("stock not enough for product_id %d", item.ProductID)
		}

		// 2. insert ke transaction_products
		_, err = tx.Exec(ctx, `
			INSERT INTO transaction_products
			(product_id, transaction_id, qty, size, variant, price)
			VALUES ($1,$2,$3,$4,$5,$6)
		`,
			item.ProductID,
			transactionID,
			item.Qty,
			item.Size,
			item.Variant,
			item.Price,
		)
		if err != nil {
			return 0, "", err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return 0, "", err
	}

	return transactionID, transactionCode, nil
}

func (r *TransactionRepository) GetTransactionDetail(ctx context.Context, transactionID, userID int) (model.Transaction, error) {
	query := `
	SELECT 
		id,
		user_id,
		transaction_code,
		full_name,
		email,
		address,
		delivery_method,
		subtotal_price,
		total_price,
		delivery_fee,
		tax,
		payment_method,
		status,
		created_at
	FROM transactions
	WHERE id = $1 AND user_id = $2
	`

	rows, err := r.db.Query(ctx, query, transactionID, userID)
	if err != nil {
		return model.Transaction{}, err
	}

	transaction, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.Transaction])
	if err != nil {
		return model.Transaction{}, err
	}

	return transaction, nil
}
