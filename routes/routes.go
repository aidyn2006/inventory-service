package routes

import (
	"inventory-service/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/products", handlers.CreateProduct)
	r.GET("/products/:id", handlers.GetProduct)
	r.PATCH("/products/:id", handlers.UpdateProduct)
	r.DELETE("/products/:id", handlers.DeleteProduct)
	r.GET("/products", handlers.ListProducts)
}
