package service

import (
	"context"
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

func (s *TransactionService) CreateTransaction(ctx context.Context, userID int, req dto.CreateTransactionRequest) (string, error) {

	//  basic validation
	if req.FullName == "" || req.Email == "" {
		return "", fmt.Errorf("invalid input")
	}

	if len(req.Items) == 0 {
		return "", fmt.Errorf("items cannot be empty")
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
