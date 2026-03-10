package service

import (
	"context"

	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/repository"
)

type CategorieService struct {
	repo *repository.CategorieRepository
}

func NewCategoriesService(repo *repository.CategorieRepository) *CategorieService {
	return &CategorieService{
		repo: repo,
	}
}

func (c *CategorieService) GetCategories(ctx context.Context) ([]dto.Category, error) {
	data, err := c.repo.GetCategories(ctx)
	if err != nil {
		return []dto.Category{}, err
	}

	var result []dto.Category

	for _, v := range data {
		result = append(result, dto.Category{
			Id:             v.Id,
			CategoriesName: v.CategoriesName,
			CreatedAt:      v.CreatedAt,
			UpdatedAt:      v.UpdatedAt,
			DeletedAt:      v.DeletedAt,
		})
	}

	return result, nil
}
