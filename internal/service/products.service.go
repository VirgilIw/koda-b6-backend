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

func (p *ProductService) UpdateProduct(ctx context.Context, req dto.UpdateProductRequest) error {
	if err := p.repo.UpdateProduct(ctx, req); err != nil {
		return err
	}
	return nil
}

func (p *ProductService) CreateProduct(ctx context.Context, req dto.CreateProductRequest) (dto.CreateProductResponse, error) {

	var result dto.CreateProductResponse

	product, err := p.repo.CreateProduct(ctx, req)

	if err != nil {
		return dto.CreateProductResponse{}, err
	}

	result = dto.CreateProductResponse{
		Id:                product.Id,
		Name:              product.Name,
		Description:       product.Description,
		Price:             product.Price,
		IsBuy1Get1:        product.IsBuy1Get1,
		IsFlashSale:       product.IsFlashSale,
		IsBirthdayPackage: product.IsBirthdayPackage,
	}

	return result, nil
}

func (p *ProductService) DeleteProduct(ctx context.Context, id int) error {
	if err := p.repo.DeleteProduct(ctx, id); err != nil {
		return err
	}
	return nil
}

func (p *ProductService) GetDetailProductById(ctx context.Context, id int, selectedSize string, selectedVariant string) (dto.ProductDetail, error) {
	detail, err := p.repo.GetDetailProductById(ctx, id)
	if err != nil {
		return dto.ProductDetail{}, err
	}

	productDetail := dto.ProductDetail{
		ID:                detail.ID,
		Name:              detail.Name,
		Price:             detail.Price,
		Variants:          detail.Variants,
		VariantPrices:     detail.VariantPrices,
		TotalTestimonials: detail.TotalTestimonials,
		Sizes:             detail.Sizes,
		SizePrices:        detail.SizePrices,
	}

	return productDetail, nil
}
