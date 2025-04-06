package domain

type Category struct {
	ID   uint
	Name string
}

type CategoryUseCase interface {
	CreateCategory(category *Category) error
	GetCategoryById(id uint) (*Category, error)
	UpdateCategory(category *Category) error
	ListCategories(filter map[string]interface{}) ([]*Category, error)
	DeleteCategory(id uint) error
}

type CategoryResponse struct {
	Name string
}
