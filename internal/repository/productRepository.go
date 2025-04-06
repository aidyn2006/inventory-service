package repository

import (
	"context"
	"errors"
	"inventory-service/internal/domain"

	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *domain.Product) error
	GetProductById(ctx context.Context, id uint) (*domain.Product, error)
	UpdateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error)
	DeleteProduct(ctx context.Context, id uint) error
	ListProducts(ctx context.Context, filter map[string]interface{}) ([]domain.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) CreateProduct(ctx context.Context, product *domain.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) GetProductById(ctx context.Context, id uint) (*domain.Product, error) {
	var product domain.Product
	err := r.db.First(&product, id).Error
	return &product, err
}
func (r *productRepository) UpdateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	result := r.db.WithContext(ctx).Model(&domain.Product{}).Where("id = ?", product.ID).Updates(product)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	// Fetch updated product
	var updatedProduct domain.Product
	if err := r.db.WithContext(ctx).First(&updatedProduct, product.ID).Error; err != nil {
		return nil, err
	}
	return &updatedProduct, nil
}

func (r *productRepository) DeleteProduct(ctx context.Context, id uint) error {
	// Check if product exists
	var product domain.Product
	if err := r.db.WithContext(ctx).First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	// Delete product
	return r.db.WithContext(ctx).Delete(&product).Error
}

func (r *productRepository) ListProducts(ctx context.Context, filter map[string]interface{}) ([]domain.Product, error) {
	var products []domain.Product
	query := r.db.WithContext(ctx)

	for key, value := range filter {
		query = query.Where(key+" = ?", value)
	}

	err := query.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
