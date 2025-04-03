package main

import (
	"inventory-service/database"
	"inventory-service/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Подключение к БД
	database.ConnectDB()

	// Регистрация маршрутов
	routes.RegisterRoutes(r)

	// Запуск сервера
	log.Println("Inventory Service is running on port 8081...")
	r.Run(":8081")
}
