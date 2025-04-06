package migrations

import (
	"gorm.io/gorm"
	"inventory-service/internal/domain"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&domain.Product{}, &domain.Category{},
	)
}
