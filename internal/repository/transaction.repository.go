package repository

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
)

type TransactionRepository struct {
	db *pgxpool.Pool
}

func NewTransactionRepository(db *pgxpool.Pool) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

// generate code (support leading zero)
func generateTransactionCode() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return fmt.Sprintf("%05d-%05d",
		r.Intn(100000),
		r.Intn(100000),
	)
}

func (r *TransactionRepository) CreateTransaction(ctx context.Context, userID int, req dto.CreateTransactionRequest) (string, error) {

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return "", err
	}

	// rollback
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	var transactionID int
	var transactionCode string

	// retry kalau code duplicate (max 3x)
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
			"completed",
		).Scan(&transactionID)

		if err == nil {
			break
		}

		// retry kalau gagal (misal duplicate)
		if i == 2 {
			return "", err
		}
	}

	// insert transaction_products
	for _, item := range req.Items {
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
			return "", err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return "", err
	}

	return transactionCode, nil
}

func (r *TransactionRepository) UpdateTransactionStatus(ctx context.Context, transactionID int, status string) error {
	_, err := r.db.Exec(ctx, `
	UPDATE transactions
	SET status = $1, updated_at = NOW()
	WHERE id = $2
`, status, transactionID)

	if err != nil {
		return err
	}

	return nil
}
