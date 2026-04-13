package repository

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
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

	return fmt.Sprintf("#%05d-%05d",
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
		t.id,
		t.transaction_code,
		t.full_name,
		t.email,
		t.address,
		t.delivery_method,
		t.subtotal_price,
		t.total_price,
		t.delivery_fee,
		t.tax,
		t.payment_method,
		t.status,
		t.user_id,
		t.created_at,
		(
			SELECT i.image_path
			FROM transaction_products tp
			LEFT JOIN product_images pi ON pi.product_id = tp.product_id
			LEFT JOIN images i ON i.id = pi.image_id
			WHERE tp.transaction_id = t.id
			LIMIT 1
		) AS product_image
	FROM transactions t
	WHERE t.user_id = $1
	ORDER BY t.created_at DESC
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
			&t.UserId,
			&t.CreatedAt,
			&t.ProductImage,
		)
		fmt.Println(err)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, t)
	}

	return transactions, nil
}

func (r *TransactionRepository) InsertTransaction(ctx context.Context, tx pgx.Tx, userID int, req dto.CreateTransactionRequest, subtotal int, total int, deliveryFee int, tax int,
) (int, string, error) {

	var transactionID int
	var transactionCode string

	for range 3 {

		transactionCode = generateTransactionCode()

		err := tx.QueryRow(ctx, `
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
			subtotal,
			total,
			deliveryFee,
			tax,
			req.PaymentMethod,
			"pending",
		).Scan(&transactionID)

		if err == nil {
			return transactionID, transactionCode, nil
		}

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			continue
		}

		return 0, "", fmt.Errorf("insert transaction failed: %w", err)
	}

	return 0, "", fmt.Errorf("failed to generate unique transaction code")
}

func (r *ProductRepository) UpdateStock(ctx context.Context, tx pgx.Tx, productID int, qty int) error {

	cmd, err := tx.Exec(ctx, `
		UPDATE products
		SET stock = stock - $1
		WHERE id = $2 AND stock >= $1
	`, qty, productID)

	if err != nil {
		return err
	}
	fmt.Println(err)

	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("stock not enough for product_id %d", productID)
	}

	return nil
}
func (r *TransactionRepository) InsertTransactionProduct(ctx context.Context, tx pgx.Tx, transactionID int, item dto.TransactionItemRequest) error {

	_, err := tx.Exec(ctx, `
		INSERT INTO transaction_products
		(product_id, transaction_id, qty, size, variant, price, product_name)
		VALUES ($1,$2,$3,$4,$5,$6,$7)
	`,
		item.ProductID,
		transactionID,
		item.Qty,
		item.Size,
		item.Variant,
		item.Price,
		item.ProductName,
	)

	if err != nil {
		return fmt.Errorf("failed to insert transaction product: %w", err)
	}

	return nil
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
