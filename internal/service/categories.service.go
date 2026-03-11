package service

import (
	"context"
	"errors"

	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/repository"
)

type CategoriesService struct {
	repo *repository.CategoriesRepository
}

func NewCategoriesService(repo *repository.CategoriesRepository) *CategoriesService {
	return &CategoriesService{
		repo: repo,
	}
}

func (c *CategoriesService) GetCategories(ctx context.Context) ([]dto.Category, error) {
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

func (c *CategoriesService) GetCategoryById(ctx context.Context, id int) (dto.Category, error) {
	category, err := c.repo.GetCategoryById(ctx, id)

	if err != nil {
		return dto.Category{}, err
	}

	var result dto.Category

	result = dto.Category{
		Id:             category.Id,
		CategoriesName: category.CategoriesName,
	}

	return result, nil
}

func (c *CategoriesService) CreateCategory(ctx context.Context, catName dto.CategoryRequest) (dto.Category, error) {
	category, err := c.repo.CreateCategory(ctx, catName)
	if err != nil {
		return dto.Category{}, err
	}

	var result dto.Category

	result = dto.Category{
		Id:             category.Id,
		CategoriesName: category.CategoriesName,
		CreatedAt:      category.CreatedAt,
	}

	return result, nil
}

func (c *CategoriesService) UpdateCategory(ctx context.Context, req dto.CreateCategoryRequest) (dto.Category, error) {

	category, err := c.repo.UpdateCategory(ctx, req)
	if err != nil {
		return dto.Category{}, err
	}

	if req.Id == 0 {
		return dto.Category{}, err
	}

	result := dto.Category{
		Id:             category.Id,
		CategoriesName: category.CategoriesName,
	}

	return result, nil
}

func (c *CategoriesService) DeleteCategory(ctx context.Context, id int) error {
	// Validasi id
	if id <= 0 {
		return errors.New("invalid category id")
	}

	// Panggil repository
	err := c.repo.DeleteCategory(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
