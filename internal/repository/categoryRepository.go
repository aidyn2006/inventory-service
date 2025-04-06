package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"inventory-service/internal/domain"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, category *domain.Category) error
	GetCategoryById(ctx context.Context, id uint) (*domain.Category, error)
	UpdateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error)
	ListCategories(ctx context.Context, filter map[string]interface{}) ([]domain.Category, error)
	DeleteCategory(ctx context.Context, id uint) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) CreateCategory(ctx context.Context, category *domain.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) GetCategoryById(ctx context.Context, id uint) (*domain.Category, error) {
	var category domain.Category
	err := r.db.First(&category, id).Error
	return &category, err
}

func (r *categoryRepository) UpdateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	result := r.db.WithContext(ctx).Model(&domain.Category{}).Where("id = ?", category.ID).Updates(category)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	// Fetch updated category
	var updatedCategory domain.Category
	if err := r.db.WithContext(ctx).First(&updatedCategory, category.ID).Error; err != nil {
		return nil, err
	}
	return &updatedCategory, nil
}

func (r *categoryRepository) ListCategories(ctx context.Context, filter map[string]interface{}) ([]domain.Category, error) {
	var categories []domain.Category
	query := r.db.WithContext(ctx)

	for key, value := range filter {
		query = query.Where(key+" = ?", value)
	}

	err := query.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) DeleteCategory(ctx context.Context, id uint) error {
	// Check if  exists
	var category domain.Category
	if err := r.db.WithContext(ctx).First(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	// Delete category
	return r.db.WithContext(ctx).Delete(&category).Error
}
