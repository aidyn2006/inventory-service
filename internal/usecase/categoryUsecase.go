package usecase

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"inventory-service/internal/domain"
	"inventory-service/internal/repository"
)

type CategoryUseCase struct {
	repo repository.CategoryRepository
}

func NewCategoryUseCase(repo repository.CategoryRepository) *CategoryUseCase {
	return &CategoryUseCase{repo: repo}
}

func (uc *CategoryUseCase) CreateCategory(ctx context.Context, category *domain.Category) error {
	return uc.repo.CreateCategory(ctx, category) // Added `ctx`
}

func (uc *CategoryUseCase) GetCategoryByID(ctx context.Context, id uint) (*domain.Category, error) {
	return uc.repo.GetCategoryById(ctx, id) // Added `ctx`
}

func (uc *CategoryUseCase) UpdateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	updatedCategory, err := uc.repo.UpdateCategory(ctx, category)
	if err != nil {
		return nil, err
	}
	return updatedCategory, nil
}

func (uc *CategoryUseCase) ListCategories(ctx context.Context, filter map[string]interface{}) ([]domain.Category, error) {
	categories, err := uc.repo.ListCategories(ctx, filter)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (uc *CategoryUseCase) DeleteCategory(ctx context.Context, id uint) error {
	err := uc.repo.DeleteCategory(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("category not found")
		}
		return err
	}
	return nil
}
