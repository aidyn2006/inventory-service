package domain

type Product struct {
	ID          uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string  `gorm:"type:varchar(100);not null" json:"name"`
	Description string  `gorm:"type:text" json:"description"`
	Price       float64 `gorm:"not null" json:"price"`
	Stock       int     `gorm:"not null" json:"stock"`
	CategoryID  uint    `gorm:"not null" json:"category_id"`
}

type ProductResponse struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	CategoryID  uint    `json:"category_id"`
}

type ProductUseCase interface {
	CreateProduct(product *Product) error
	GetProductById(id uint) (*Product, error)
	UpdateProduct(product *Product) error
	DeleteProduct(id uint) error
	ListProducts(filter map[string]interface{}) ([]Product, error)
}

type ProductRepository interface {
	CreateProduct(product *Product) error
	GetProductById(id uint) (*Product, error)
	UpdateProduct(product *Product) error
	DeleteProduct(id uint) error
	ListProducts(filter map[string]interface{}) ([]Product, error)
}
