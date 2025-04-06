package handlers

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(productHandler *ProductHandler, categoryHandler *CategoryHandler) *gin.Engine {
	router := gin.Default()
	api := router.Group("/api/v1")
	{
		products := api.Group("/products")
		{
			products.GET("", productHandler.ListProducts)
			products.POST("", productHandler.CreateProduct)
			products.PUT("/:id", productHandler.UpdateProduct)
			products.DELETE("/:id", productHandler.DeleteProduct)
			products.GET("/:id", productHandler.GetProductById)
		}
		categories := api.Group("/categories")
		{
			categories.POST("", categoryHandler.CreateCategory)
			categories.GET("", categoryHandler.ListCategories)
			categories.GET("/:id", categoryHandler.GetCategoryById)
			categories.PATCH("/:id", categoryHandler.UpdateCategory)
			categories.DELETE("/:id", categoryHandler.DeleteCategory)
		}
	}
	return router
}
