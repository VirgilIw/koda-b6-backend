package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/repository"
)

type TransactionService struct {
	repo *repository.TransactionRepository
}

func NewTransactionService(repo *repository.TransactionRepository) *TransactionService {
	return &TransactionService{
		repo: repo,
	}
}

func (s *TransactionService) GetTransactionsByUserID(ctx context.Context, userID int) ([]dto.TransactionResult, error) {
	result, err := s.repo.GetTransactionsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, errors.New("no transactions found")
	}

	return result, nil
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

	//  basic validation
	if req.FullName == "" || req.Email == "" {
		return 0, "", fmt.Errorf("invalid input")
	}

	if len(req.Items) == 0 {
		return 0, "", fmt.Errorf("items cannot be empty")
	}

	return s.repo.CreateTransaction(ctx, userID, req)
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
