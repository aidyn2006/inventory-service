package usecase

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"inventory-service/internal/domain"
	"inventory-service/internal/repository"
)

type ProductUseCase struct {
	repo repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) *ProductUseCase {
	return &ProductUseCase{repo: repo}
}

func (uc *ProductUseCase) CreateProduct(ctx context.Context, product *domain.Product) error {
	if product.Stock < 0 {
		return errors.New("stock cannot be negative")
	}
	return uc.repo.CreateProduct(ctx, product) // Added `ctx`
}

func (uc *ProductUseCase) GetProductByID(ctx context.Context, id uint) (*domain.Product, error) {
	return uc.repo.GetProductById(ctx, id) // Added `ctx`
}

func (uc *ProductUseCase) UpdateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	if product.Stock < 0 {
		return nil, errors.New("stock cannot be negative")
	}

	updatedProduct, err := uc.repo.UpdateProduct(ctx, product)
	if err != nil {
		return nil, err
	}
	return updatedProduct, nil
}

func (uc *ProductUseCase) ListProducts(ctx context.Context, filter map[string]interface{}) ([]domain.Product, error) {
	products, err := uc.repo.ListProducts(ctx, filter)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (uc *ProductUseCase) DeleteProduct(ctx context.Context, id uint) error {
	err := uc.repo.DeleteProduct(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("product not found")
		}
		return err
	}
	return nil
}
