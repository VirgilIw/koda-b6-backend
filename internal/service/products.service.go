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

func (p *ProductService) GetProducts(ctx context.Context, page int) ([]dto.Product, error) {
	data, err := p.repo.GetProducts(ctx, page)
	if err != nil {
		return nil, err
	}

	result := make([]dto.Product, 0, len(data))

	for _, v := range data {
		var rating float64

		if v.Rating != nil {
			rating = *v.Rating
		} else {
			rating = 0
		}

		result = append(result, dto.Product{
			ID:                v.ID,
			Name:              v.Name,
			Rating:            rating,
			Description:       v.Description,
			Price:             v.Price,
			IsBuy1Get1:        v.IsBuy1Get1,
			IsFlashSale:       v.IsFlashSale,
			IsBirthdayPackage: v.IsBirthdayPackage,
			Image:             v.Image,
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
		Id:                product.ID,
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
		Images:            detail.Images,
		Description:       detail.Description,
		Price:             detail.Price,
		Variants:          detail.Variants,
		VariantPrices:     detail.VariantPrices,
		TotalTestimonials: detail.TotalTestimonials,
		Sizes:             detail.Sizes,
		SizePrices:        detail.SizePrices,
	}

	return productDetail, nil
}
