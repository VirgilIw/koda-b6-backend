package service

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/repository"
)

type TransactionService struct {
	repo        *repository.TransactionRepository
	productRepo *repository.ProductRepository
	sizeRepo    *repository.SizesRepository
	variantRepo *repository.VariantRepository
	db          *pgxpool.Pool
}

func NewTransactionService(repo *repository.TransactionRepository, productRepo *repository.ProductRepository, sizeRepo *repository.SizesRepository, variantRepo *repository.VariantRepository, db *pgxpool.Pool) *TransactionService {
	return &TransactionService{
		repo:        repo,
		productRepo: productRepo,
		sizeRepo:    sizeRepo,
		variantRepo: variantRepo,
		db:          db,
	}
}

func (s *TransactionService) GetTransactionsByUserID(ctx context.Context, userID int) ([]dto.TransactionResult, error) {
	return s.repo.GetTransactionsByUserID(ctx, userID)
}

func (s *TransactionService) GetTransactionDetail(ctx context.Context, userID, transactionID int) (*dto.TransactionDetailResponse, error) {

	t, err := s.repo.GetTransactionDetail(ctx, transactionID, userID)
	if err != nil {
		return nil, err
	}

	items, err := s.repo.GetTransactionItems(ctx, transactionID)
	if err != nil {
		return nil, err
	}

	var itemDTO []dto.TransactionItemResponse
	for _, i := range items {
		itemDTO = append(itemDTO, dto.TransactionItemResponse{
			ProductID: i.ProductID,
			Qty:       i.Qty,
			Size:      i.Size,
			Variant:   i.Variant,
			Price:     i.Price,
		})
	}

	// 4. mapping header → dto
	result := &dto.TransactionDetailResponse{
		ID:              t.ID,
		TransactionCode: t.TransactionCode,
		FullName:        t.FullName,
		Email:           t.Email,
		Address:         t.Address,
		DeliveryMethod:  t.DeliveryMethod,
		SubtotalPrice:   t.SubtotalPrice,
		TotalPrice:      t.TotalPrice,
		DeliveryFee:     t.DeliveryFee,
		Tax:             t.Tax,
		PaymentMethod:   t.PaymentMethod,
		Status:          t.Status,
		CreatedAt:       &t.CreatedAt,
		Items:           itemDTO,
	}

	return result, nil
}

func (s *TransactionService) CreateTransaction(ctx context.Context, userID int, req dto.CreateTransactionRequest) (int, string, error) {

	if req.FullName == "" || req.Email == "" || req.Address == "" {
		return 0, "", fmt.Errorf("invalid input")
	}

	if len(req.Items) == 0 {
		return 0, "", fmt.Errorf("items cannot be empty")
	}

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return 0, "", err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback(ctx)
			panic(p)
		} else if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	type itemSnapshot struct {
		ProductID    int
		Qty          int
		Size         string
		Variant      string
		Price        int
		ProductName  string
		ProductImage string
	}

	var (
		items    []itemSnapshot
		subtotal int
	)

	for _, item := range req.Items {

		if item.Qty <= 0 {
			return 0, "", fmt.Errorf("invalid qty for product_id %d", item.ProductID)
		}

		product, image, err := s.productRepo.GetProductSnapshot(ctx, tx, item.ProductID)
		if err != nil {
			return 0, "", err
		}

		if err := s.productRepo.UpdateStock(ctx, tx, item.ProductID, item.Qty); err != nil {
			return 0, "", err
		}

		basePrice := product.Price

		sizePrice, err := s.sizeRepo.GetSizePrice(ctx, tx, item.Size)
		if err != nil {
			return 0, "", err
		}

		variantPrice, err := s.variantRepo.GetVariantPrice(ctx, tx, item.Variant)
		if err != nil {
			return 0, "", err
		}

		finalPrice := basePrice + sizePrice + variantPrice

		subtotal += finalPrice * item.Qty

		items = append(items, itemSnapshot{
			ProductID:    item.ProductID,
			Qty:          item.Qty,
			Size:         item.Size,
			Variant:      item.Variant,
			Price:        finalPrice,
			ProductName:  product.Name,
			ProductImage: image,
		})
	}

	deliveryFee := 0
	if req.DeliveryMethod == "door delivery" {
		deliveryFee = 10000
	}

	tax := 4000
	total := subtotal + tax + deliveryFee

	transactionID, transactionCode, err := s.repo.InsertTransaction(
		ctx,
		tx,
		userID,
		req,
		subtotal,
		total,
		deliveryFee,
		tax,
	)
	if err != nil {
		return 0, "", err
	}

	for _, item := range items {
		err = s.repo.InsertTransactionProduct(
			ctx,
			tx,
			transactionID,
			dto.TransactionItemRequest{
				ProductID:    item.ProductID,
				Qty:          item.Qty,
				Size:         item.Size,
				Variant:      item.Variant,
				Price:        item.Price,
				ProductName:  item.ProductName,
				ProductImage: item.ProductImage,
			},
		)
		if err != nil {
			return 0, "", fmt.Errorf("failed insert transaction product: %w", err)
		}
	}

	if err = tx.Commit(ctx); err != nil {
		return 0, "", err
	}

	return transactionID, transactionCode, nil
}

func isValidStatus(status string) bool {
	switch status {
	case "on_progress", "sending", "completed":
		return true
	default:
		return false
	}
}

func (s *TransactionService) UpdateTransactionStatus(ctx context.Context, transactionID int, status string) error {

	if !isValidStatus(status) {
		return fmt.Errorf("invalid status")
	}

	return s.repo.UpdateTransactionStatus(ctx, transactionID, status)
}
