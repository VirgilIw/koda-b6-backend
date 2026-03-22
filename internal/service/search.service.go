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

	if len(data) == 0 && req.Name != "" {
		return nil, fmt.Errorf("product '%s' not found", req.Name)
	}

	var result []dto.ProductFilter

	for _, v := range data {

		finalPrice := v.Price

		if v.IsFlashSale {
			finalPrice = v.Price - 2000
		}

		result = append(result, dto.ProductFilter{
			ID:                v.ID,
			Name:              v.Name,
			Description:       v.Description,
			Price:             v.Price,
			FinalPrice:        finalPrice,
			Categories:        v.Categories,
			Images:            v.Images,
			Rating:            v.Rating,
			IsFlashSale:       v.IsFlashSale,
			IsBuy1Get1:        v.IsBuy1Get1,
			IsBirthdayPackage: v.IsBirthdayPackage,
		})
	}

	return result, nil
}
