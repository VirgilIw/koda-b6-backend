package service

import (
	"context"
	"fmt"

	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/repository"
)

type SearchService struct {
	repo         *repository.SearchRepository
	repoCategory *repository.CategoriesRepository
}

func NewSearchService(repo *repository.SearchRepository, repoCategory *repository.CategoriesRepository) *SearchService {
	return &SearchService{
		repo:         repo,
		repoCategory: repoCategory,
	}
}

func (s *SearchService) SearchProducts(ctx context.Context, req dto.SearchProductRequest) ([]dto.ProductFilter, error) {

	if req.Category != "" {
		exists, err := s.repoCategory.ExistsByName(ctx, req.Category)
		if err != nil {
			return nil, err
		}

		if !exists {
			return nil, fmt.Errorf("category '%s' not found", req.Category)
		}
	}

	data, err := s.repo.SearchProducts(ctx, req)
	if err != nil {
		return nil, err
	}

	var result []dto.ProductFilter

	for _, v := range data {
		result = append(result, dto.ProductFilter{
			Name:              v.Name,
			Description:       v.Description,
			Price:             v.Price,
			Category:          v.Category,
			IsFlashSale:       v.IsFlashSale,
			IsBuy1Get1:        v.IsBuy1Get1,
			IsBirthdayPackage: v.IsBirthdayPackage,
		})
	}

	return result, nil
}
