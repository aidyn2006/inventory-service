// @title Inventory Service API
// @version 1.0
// @description This is an API for managing inventory of products and categories
// @host localhost:8080
// @BasePath /
// @schemes http
package main

import (
	"inventory-service/config"
	"inventory-service/db"
	"inventory-service/internal/delivery/http/handlers"
	"inventory-service/internal/repository"
	"inventory-service/internal/usecase"
	"log"
)

func main() {
	cfg := config.Load()

	dbConn, err := db.NewPostgresDB(cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	db.RunMigrations(dbConn)

	productRepo := repository.NewProductRepository(dbConn)
	categoryRepo := repository.NewCategoryRepository(dbConn)

	productUseCase := usecase.NewProductUseCase(productRepo)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepo)

	productHandler := handlers.NewProductHandler(productUseCase)
	categoryHandler := handlers.NewCategoryHandler(categoryUseCase)

	router := handlers.NewRouter(productHandler, categoryHandler)

	// Start server
	log.Printf("Server running on port %s", cfg.ServerPort)
	router.Run(":" + cfg.ServerPort)

}

// @title Inventory Service API
// @version 1.0
// @description This is a REST API for managing products and categories.
// @host localhost:8080
// @BasePath /
