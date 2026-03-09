package service

import (
	"context"

	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (p *ProductService) GetProducts(ctx context.Context) ([]dto.Product, error) {
	data, err := p.repo.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	var result []dto.Product
	for _, v := range data {
		result = append(result, dto.Product{
			Id:                v.Id,
			Name:              v.Name,
			Description:       v.Description,
			Price:             v.Price,
			IsBuy1Get1:        v.IsBuy1Get1,
			IsFlashSale:       v.IsFlashSale,
			IsBirthdayPackage: v.IsBirthdayPackage,
		})
	}

	return result, nil
}
