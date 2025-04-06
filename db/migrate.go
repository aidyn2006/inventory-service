package db

import (
	"fmt"
	"gorm.io/gorm"
	"inventory-service/db/migrations"
)

func RunMigrations(db *gorm.DB) {
	fmt.Println("Running database migrations...")
	err := migrations.Migrate(db)
	if err != nil {
		fmt.Println("Migration failed:", err)
	} else {
		fmt.Println("Migrations applied successfully!")
	}
}
